// service/customer_service.go
package service

import (
    "kstyle-hub/model"
    "kstyle-hub/repository"
)

type CustomerService struct {
    Repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
    return &CustomerService{Repo: repo}
}

func (s *CustomerService) GetAllCustomers(limit, offset int) ([]model.Customer, error) {
    return s.Repo.FindAll(limit, offset)
}

func (s *CustomerService) GetCustomerByID(id uint) (*model.Customer, error) {
    return s.Repo.FindByID(id)
}

func (s *CustomerService) CreateCustomer(customer *model.Customer) error {
    return s.Repo.Create(customer)
}

func (s *CustomerService) UpdateCustomer(customer *model.Customer) error {
    return s.Repo.Update(customer)
}

func (s *CustomerService) DeleteCustomer(id uint) error {
    return s.Repo.Delete(id)
}

func (s *CustomerService) SearchCustomers(query string, limit, offset int) ([]model.Customer, error) {
    return s.Repo.Search(query, limit, offset)
}
