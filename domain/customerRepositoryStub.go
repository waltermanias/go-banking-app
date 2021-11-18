package domain

import (
	"banking-app/errs"
)

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	for _, c := range s.customers {
		if c.Id == id {
			return &c, nil
		}
	}

	return nil, errs.NewNotFoundError("Customer was not found")
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "John Doe", City: "Seattle", Zipcode: "9999", DateOfBirth: "1987-06-12", Status: "1"},
		{Id: "1002", Name: "Danna Doe", City: "Seattle", Zipcode: "9999", DateOfBirth: "1992-02-10", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
