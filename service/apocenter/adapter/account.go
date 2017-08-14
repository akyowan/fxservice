package adapter

import (
	"fxservice/service/apocenter/domain"
)

type AccountQueryParam struct {
	Account      string
	SN           string
	Group        string
	Status       domain.ApoTaskStatus
	MaxRoundUsed int
	MaxTotalUsed int
	EnableApp    int64
}

// GetAndLockFreeAccount
// Get an free account from account cache
// lock it
func GetAndLockFreeAccount(queryParam *AccountQueryParam) (error, *domain.Account) {
	return nil, nil
}

// UnlockFreeAccount
// Unlock an locked account in account cache
func UnlockFreeAccount(account int64) error {
	return nil
}

// RoundRestAccount
// Round rest accout in account cache
func RoundRestAccount(account int64) error {
	return nil
}
