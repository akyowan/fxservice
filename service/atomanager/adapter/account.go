package adapter

import (
	"fxservice/service/atomanager/domain"
)

type AddAccountResult struct {
	Exists    []string         `json:"exists"`
	NoDevices []string         `json:"no_devices"`
	Errors    []domain.Account `json:"errors"`
	Success   int              `json:"success"`
}

func AddAccount(brief string, weight int, accounts []domain.Account) (*AddAccountResult, error) {
	db := dbPool.NewConn().Begin()
	result := AddAccountResult{
		Success:   0,
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
		db.Rollback()
		return nil, dbResult.Error
	}

	var devices []domain.Device
	if err := db.Where("bind_count = 0 AND status = 1").Limit(len(adds)).Find(&devices).Updates(domain.Device{BindCount: 1}).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	noDevices := adds[len(devices):]
	for i := range noDevices {
		result.NoDevices = append(result.NoDevices, noDevices[i].Account)
	}
	result.Success = len(devices)

	startId, err := getStartId(brief)
	if err != nil {
		return nil, err
	}

	for i := range devices {
		account = adds[i]
		device := devices[i]
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
		if err := db.Create(&account).Error; err != nil {
			db.Rollback()
			return nil, err
		}
		startId += 1
	}

	if err := updateBrief(brief, weight, result.Success); err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()
	return &result, nil
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