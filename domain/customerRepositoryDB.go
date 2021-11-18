package domain

import (
	"banking-app/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustormerRepositoryDb struct {
	client *sql.DB
}

func (d CustormerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name , city, zipcode, date_of_birth, status from customers where customer_id= ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning customers table" + err.Error())
			return nil, errs.NewUnexpectedError()
		}
	}

	return &c, nil
}

func (d CustormerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name , city, zipcode, date_of_birth, status from customers where 0=1"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError()
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers table" + err.Error())
			return nil, errs.NewUnexpectedError()
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustormerRepositoryDb {
	client, err := sql.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustormerRepositoryDb{client}
}
