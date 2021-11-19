package domain

import (
	"banking-app/dto"
	"banking-app/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type Accounts []Account

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToDto() dto.AccountResponse {
	return dto.AccountResponse{AccountId: a.AccountId, CustomerId: a.CustomerId, OpeningDate: a.OpeningDate, AccountType: a.AccountType, Amount: a.Amount, Status: a.Status}
}

func (a Accounts) ToDto() []dto.AccountResponse {
	var dtos = make([]dto.AccountResponse, 0)
	for _, account := range a {
		dtos = append(dtos, account.ToDto())
	}
	return dtos
}
