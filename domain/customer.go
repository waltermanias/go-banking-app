package domain

import (
	"banking-app/dto"
	"banking-app/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type Customers []Customer

type CustomerRepository interface {
	FindAll(string) ([]dto.CustomerResponse, *errs.AppError)
	ById(string) (*dto.CustomerResponse, *errs.AppError)
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{Id: c.Id, Name: c.Name, City: c.City, Zipcode: c.Zipcode, DateOfBirth: c.DateOfBirth, Status: c.StatusAsText()}
}

func (c Customers) ToDto() []dto.CustomerResponse {
	var customers []dto.CustomerResponse
	for _, customer := range c {
		customerDto := customer.ToDto()
		customers = append(customers, customerDto)
	}
	return customers
}
