package domain

import (
	"github.com/singhparbjyot/banking/dto"
	"github.com/singhparbjyot/banking/errors"
)

// We define a customer
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}
func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// Repository interface
type CustomerRepository interface {
	//Function which return all the customers or error
	FindAll(status string) ([]Customer, error)
	//Function which return a cutomer pointer and error,we use a pointer because without pointer we
	// can t return nil
	ById(id string) (*Customer, *errors.AppError)
}
