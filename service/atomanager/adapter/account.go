package adapter

import (
	"fxlibraries/loggers"
	"fxservice/service/atomanager/domain"
)

type AddAccountResult struct {
	Exists    []string         `json:"exists"`
	NoDevices []string         `json:"no_devices"`
	Errors    []domain.Account `json:"errors"`
	Success   int              `json:"success"`
}

func AddAccount(brief string, dGroup string, weight int, accounts []domain.Account) (*AddAccountResult, error) {
	db := dbPool.NewConn()
	result := AddAccountResult{
		Success:   0,
		Exists:    []string{},
		NoDevices: []string{},
		Errors:    []domain.Account{},
	}
	var adds []domain.Account
	var account domain.Account
	for i := range accounts {
		if accounts[i].Account == "" {
			continue
		}
		if accounts[i].Passwd == "" {
			result.Errors = append(result.Errors, accounts[i])
			continue
		}
		dbResult := db.Select("account").Where("account = ?", accounts[i].Account).First(&account)
		if dbResult.Error == nil {
			result.Exists = append(result.Exists, accounts[i].Account)
			continue
		}
		if dbResult.RecordNotFound() {
			adds = append(adds, accounts[i])
			continue
		}
		return nil, dbResult.Error
	}
	if len(adds) <= 0 {
		db.Rollback()
		return &result, nil
	}

	trans := db.Begin()
	var devices []domain.Device
	cur := trans
	if dGroup != "" {
		cur = cur.Where(&domain.Device{Group: dGroup})
	}
	if err := cur.Where("bind_count = 0").Where("status = 1").Limit(len(adds)).Find(&devices).Error; err != nil {
		cur.Rollback()
		return nil, err
	}
	var enableDevices []domain.Device
	for i := range devices {
		device := devices[i]
		if err := trans.Table(device.TableName()).Where("id = ?", device.Id).Updates(domain.Device{BindCount: 1}).Error; err != nil {
			trans.Rollback()
			return nil, err
		}
		dbResult := trans.Table(account.TableName()).Where("sn = ?", device.Sn).First(&account)
		if dbResult.RecordNotFound() {
			enableDevices = append(enableDevices, device)
			continue
		}
		if dbResult.Error != nil {
			trans.Rollback()
			return nil, dbResult.Error
		}
		loggers.Warn.Printf("AddAccount sn used [%v]", account)
	}

	noDevices := adds[len(enableDevices):]
	for i := range noDevices {
		result.NoDevices = append(result.NoDevices, noDevices[i].Account)
	}
	result.Success = len(enableDevices)
	if result.Success == 0 {
		trans.Rollback()
		return &result, nil
	}

	startId, err := getStartId(brief)
	if err != nil {
		trans.Rollback()
		return nil, err
	}

	for i := range enableDevices {
		account = adds[i]
		device := enableDevices[i]
		account.Id = startId
		account.Sn = device.Sn
		account.Imei = device.Imei
		account.Seq = device.Seq
		account.Model = device.Model
		account.BuildNum = device.BuildNum
		account.Mac = device.Mac
		account.HardWare = device.HardWare
		account.Wifi = device.Wifi
		account.Version = device.Version
		account.Brief = brief
		account.Status = 1
		if err := trans.Create(&account).Error; err != nil {
			trans.Rollback()
			return nil, err
		}
		startId += 1
	}

	if err := updateBrief(brief, weight, result.Success); err != nil {
		trans.Rollback()
		return nil, err
	}
	if err := deleteBriefCache(); err != nil {
		trans.Rollback()
		return nil, err
	}

	trans.Commit()
	return &result, nil
}

func RebindAccount() error {
	db := dbPool.NewConn().Begin()
	var (
		accounts []domain.Account
		devices  []domain.Device
	)
	limit := 100
	for {
		if err := db.Where("errno = ?", 22).Limit(limit).Find(&accounts).Error; err != nil {
			loggers.Info.Printf("RebindAccount no account need rebind")
			db.Rollback()
			return err
		}
		if len(accounts) == 0 {
			loggers.Info.Printf("RebindAccount no account need rebind")
			db.Rollback()
			return nil
		}

		if err := db.Where("bind_count = 0").Where("status = 1").Limit(len(accounts)).Find(&devices).Error; err != nil {
			db.Rollback()
			return err
		}

		if len(devices) < len(accounts) {
			loggers.Warn.Printf("RebindAccount no enought device for rebind account need:%d get:%d", len(accounts), len(devices))
		}

		for i, device := range devices {
			account := accounts[i]
			account.Sn = device.Sn
			account.Errno = 0
			account.Status = 1
			device.BindCount = 1
			device.Status = 0
			if err := db.Model(account).Save(account).Error; err != nil {
				db.Rollback()
				return err
			}
			if err := db.Model(device).Update(device).Error; err != nil {
				db.Rollback()
				return err
			}
			loggers.Info.Printf("RebindAccount account:%s sn:%s", account.Account, device.Sn)
		}
		if len(accounts) < limit {
			break
		}
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

func getStartId(brief string) (int, error) {
	db := dbPool.NewConn()
	var account domain.Account
	dbResult := db.Where("brief = ?", brief).Order("id desc").First(&account)
	if dbResult.Error == nil {
		return account.Id + 1, nil
	}
	if !dbResult.RecordNotFound() {
		return 0, dbResult.Error
	}

	dbResult = db.Order("id desc").First(&account)
	if dbResult.RecordNotFound() {
		return 1000100001, nil
	}
	if dbResult.Error != nil {
		return 0, dbResult.Error
	}
	id := (((account.Id / 100000) + 1) * 100000) + 1
	return id, nil
}

func updateBrief(brief string, weight, total int) error {
	db := dbPool.NewConn()
	var group domain.AccountGroup
	dbResult := db.Where("brief = ?", brief).First(&group)
	if (dbResult.Error != nil) && (!dbResult.RecordNotFound()) {
		return dbResult.Error
	}
	if dbResult.RecordNotFound() {
		group = domain.AccountGroup{
			Brief:  brief,
			Weight: weight,
			Total:  total,
		}
		if err := db.Create(&group).Error; err != nil {
			return err
		}
		return nil
	} else {
		group.Weight = weight
		group.Total += total
		if err := db.Save(&group).Error; err != nil {
			return err
		}
		return nil
	}
}

func deleteBriefCache() error {
	key := "ACCOUNT_ALL_BRIEFS"
	if err := redisPool.Del(key).Err(); err != nil {
		return err
	}
	key = "ACCOUNT_MAX_RANGE"
	if err := redisPool.Del(key).Err(); err != nil {
		return err
	}
	return nil
}
