package adapter

import (
	"fxservice/service/chatcenter/domain"
	"time"
)

func AddTantanAccount(account *domain.TantanAccount) error {
	db := dbPool.NewConn()
	if err := db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func CompleteTantanAccount(tid int64, tantanAccount *domain.TantanAccount) error {
	db := dbPool.NewConn()
	now := time.Now()
	updateMap := map[string]interface{}{
		"account":       tantanAccount.Account,
		"Password":      tantanAccount.Password,
		"status":        tantanAccount.Status,
		"register_time": &now,
		"register_host": tantanAccount.RegisterHost,
	}
	if err := db.Model(&tantanAccount).
		Where("tid = ? and status = ?", tid, domain.AccountStatusRegistering).
		Updates(updateMap).Error; err != nil {
		return err
	}
	return nil
}
