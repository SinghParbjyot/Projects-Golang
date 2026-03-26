package service

import (
	"github.com/singhparbjyot/banking/domain"
	"github.com/singhparbjyot/banking/dto"
	"github.com/singhparbjyot/banking/errors"
)

// Interface with the entry , which defines what  we can do in outr application
type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, error)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

// The logic of our application
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// Just calls the repository and this one gives back the customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {

	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

// Intialize the serivce is like a constructor
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
