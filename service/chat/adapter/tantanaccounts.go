package adapter

import (
	"fxservice/domain"
)

func AddTantanAccount(account *domain.TantanAccount) error {
	db := dbPool.NewConn()
	if err := db.Create(account).Error; err != nil {
		return err
	}
	return nil
}
