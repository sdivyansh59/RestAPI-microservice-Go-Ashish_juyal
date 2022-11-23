package service

import "github.com/sdivyansh59/banking/domain"

// here we will connect 1port with 2nd port

// port
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer () ([]domain.Customer,error){
	return s.repo.FindAll()
}

func NewCustomerService (respository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{respository}
}

