package dto

import (
	"banking-app/errs"
	"strings"
)

type AccountRequest struct {
	CustomerId  string  `json:"customerId"`
	OpeningDate string  `json:"openingDate"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountResponse struct {
	AccountId   string  `json:"id"`
	CustomerId  string  `json:"customerId"`
	OpeningDate string  `json:"openingDate"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

func (a AccountRequest) StatusToString() string {
	if a.Status == "active" {
		return "1"
	} else if a.Status == "inactive" {
		return "0"
	}
	return "0"
}

func (a AccountRequest) Validate() *errs.AppError {
	if a.Amount < 5000 {
		return errs.NewValidationError("The amount cannot be less than 5000")
	}

	if strings.ToLower(a.AccountType) != "saving" && strings.ToLower(a.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}

	return nil
}
