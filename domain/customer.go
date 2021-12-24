package domain

import "github.com/leslesnoa/go-microservices-demo/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	// status == ! status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

/* cutomerRepositoryStubへ定義を移動 */
// type CustomerRepositoryStub struct {
// 	customers []Customer
// }

// func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
// 	return s.customers, nil
// }

// func NewCustomerRepositoryStub() CustomerRepositoryStub {
// 	customers := []Customer{
// 		{Id: "1001", Name: "Ashish", City: "New Delhi", Zipcode: "110075", DateofBirth: "2000-01-01", Status: "1"},
// 		{Id: "1002", Name: "Rob", City: "New Delhi", Zipcode: "110075", DateofBirth: "2000-01-01", Status: "1"},
// 	}
// 	return CustomerRepositoryStub{customers}
// }
