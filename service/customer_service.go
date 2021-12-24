package service

import (
	"github.com/leslesnoa/go-microservices-demo/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepositoryStub
}

func (s DefaultCustomerService) GetAllCustomer() {

}
