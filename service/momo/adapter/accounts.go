package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
	"math/rand"
	"time"
)

func GetNewMomoAccount(gps *domain.GPSLocation) (*domain.MomoAccount, error) {
	db := dbPool.NewConn().Begin()
	var momoAccount domain.MomoAccount
	dbResult := db.Where("status = ?", domain.MomoAccountStatusUnRegister).Order("tid").First(&momoAccount)
	if dbResult.RecordNotFound() {
		db.Rollback()
		return nil, errors.NotFound
	}
	if dbResult.Error != nil {
		db.Rollback()
		return nil, dbResult.Error
	}
	avatar, err := GetFreeAvatar()
	if err != nil {
		db.Rollback()
		return nil, err
	}

	momoAccount.PhotosID = avatar.PhotosID
	momoAccount.Avatar = avatar.URL
	momoAccount.Status = domain.MomoAccountStatusRegistering
	momoAccount.Province = gps.Province
	momoAccount.City = gps.City
	momoAccount.GPSID = gps.GPSID
	if err := db.Save(&momoAccount).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return &momoAccount, nil
}

type AccountQueryParam struct {
	Status      []domain.MomoAccountStatus
	Type        domain.MomoAccountType
	Account     string
	MomoAccount string
	Gender      domain.GenderType
	Province    string
	City        string
	Operator    string
	Begin       *time.Time
	End         *time.Time
	Limit       int
	Offset      int
}

func GetMomoAccounts(param *AccountQueryParam) ([]domain.MomoAccount, error) {
	var accounts []domain.MomoAccount
	db := dbPool.NewConn()
	if param.Status != nil {
		db = db.Where("status in (?)", param.Status)
	}
	if param.Type != 0 {
		db = db.Where("type = ?", param.Type)
	}
	if param.Account != "" {
		db = db.Where("account = ?", param.Account)
	}
	if param.MomoAccount != "" {
		db = db.Where("momo_account = ?", param.MomoAccount)
	}
	if param.Gender != 0 {
		db = db.Where("gender = ?", param.Gender)
	}
	if param.Province != "" {
		db = db.Where("province = ?", param.Province)
	}
	if param.City != "" {
		db = db.Where("city = ?", param.City)
	}
	if param.Operator != "" {
		db = db.Where("operator = ?", param.Operator)
	}
	if param.Begin != nil {
		db = db.Where("create_time > ?", param.Begin)
	}
	if param.End != nil {
		db = db.Where("create_time < ?", param.End)
	}
	if err := db.Offset(param.Offset).Limit(param.Limit).Order("create_time desc").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func AddAccounts(accounts *[]domain.MomoAccount) error {
	db := dbPool.NewConn().Begin()
	for i := range *accounts {
		account := (*accounts)[i]
		if err := db.Create(&account).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

func PatchMomoAccounts(accounts *[]domain.MomoAccount) error {
	db := dbPool.NewConn().Begin()
	for i := range *accounts {
		if (*accounts)[i].Account == "" {
			continue
		}
		account := domain.MomoAccount{
			Account: (*accounts)[i].Account,
		}
		if (*accounts)[i].Province != "" {
			account.Province = (*accounts)[i].Province
		}
		if (*accounts)[i].City != "" {
			account.City = (*accounts)[i].City
		}
		if (*accounts)[i].Status != 0 {
			account.Status = (*accounts)[i].Status
		}
		if err := db.Model(&account).Where("account = ?", account.Account).UpdateColumns(account).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

func GetRandNickName() (*domain.NickName, error) {
	db := dbPool.NewConn()
	var nickNames []domain.NickName
	dbResult := db.Find(&nickNames)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Intn(len(nickNames))
	return &nickNames[index], nil
}

func CompleteMomoAccount(account string, momoAccount *domain.MomoAccount) error {
	db := dbPool.NewConn()
	now := time.Now()
	updateMap := map[string]interface{}{
		"momo_account":  momoAccount.MomoAccount,
		"status":        momoAccount.Status,
		"register_time": &now,
		"register_host": momoAccount.RegisterHost,
	}
	if err := db.Model(&momoAccount).
		Where("account = ? and status = ?", account, domain.MomoAccountStatusRegistering).
		Updates(updateMap).Error; err != nil {
		return err
	}
	return nil
}

type FreeAccountsQueryParam struct {
	City     string
	Province string
	Limit    int
}

func GetFreeAccounts(param *FreeAccountsQueryParam) (*[]domain.MomoAccount, error) {
	var accounts []domain.MomoAccount
	db := dbPool.NewConn().Begin()
	db = db.Where("status = ?", domain.MomoAccountStatusFree)
	dbResult := db.Where("province = ? and city = ?", param.Province, param.City).Limit(param.Limit)

	return nil, nil
}
