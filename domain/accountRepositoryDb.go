package domain

import (
	"banking-app/errs"
	"banking-app/logger"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (a AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	t, err2 := time.Parse(time.RFC3339, account.OpeningDate)
	if err2 != nil {
		return nil, errs.NewUnexpectedError()
	}
	result, err := a.client.Exec(sqlInsert, account.CustomerId, t, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while creating new account")
		fmt.Println(err)
		return nil, errs.NewUnexpectedError()
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			logger.Error("Error getting id")
			return nil, errs.NewUnexpectedError()
		}
		account.AccountId = strconv.FormatInt(id, 10)
		return &account, nil
	}
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client}
}
