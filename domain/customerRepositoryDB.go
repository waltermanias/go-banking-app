package domain

import (
	"banking-app/dto"
	"banking-app/errs"
	"banking-app/logger"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type CustormerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustormerRepositoryDb) ById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customerSql := "select customer_id, name , city, zipcode, date_of_birth, status from customers where customer_id= ?"

	var c Customer

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("customer not found")
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customers table" + err.Error())
			return nil, errs.NewUnexpectedError()
		}
	}

	response := c.ToDto()

	return &response, nil
}

func (d CustormerRepositoryDb) FindAll(status string) ([]dto.CustomerResponse, *errs.AppError) {

	findAllSql := "select customer_id, name , city, zipcode, date_of_birth, status from customers"

	if status != "" {
		if status == "active" {
			findAllSql = findAllSql + " where status='1'"
		} else {
			findAllSql = findAllSql + " where status='0'"
		}
	}

	customers := make([]Customer, 0)

	err := d.client.Select(&customers, findAllSql)

	if err != nil {
		logger.Error("Error while querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError()
	}

	response := Customers(customers).ToDto()

	return response, nil
}

func NewCustomerRepositoryDb() CustormerRepositoryDb {

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dataSource := fmt.Sprintf("%s@tcp(localhost:%s)/%s", dbUser, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustormerRepositoryDb{client}
}
