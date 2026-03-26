package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

// Method which reutrn all the customers
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// Method which creates a new repoistory of customers ,basically the warehouse
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1000", Name: "Marcos", City: "Zaragoza", Zipcode: "100015", DateOfBirth: "12/6/2026", Status: "1"},
		{Id: "1001", Name: "Guillermo", City: "Huesca", Zipcode: "10011", DateOfBirth: "26/9/2026", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
