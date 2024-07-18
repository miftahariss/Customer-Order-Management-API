// service/order_service.go
package service

import (
    "kstyle-hub/model"
    "kstyle-hub/repository"
)

type OrderService struct {
    Repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
    return &OrderService{Repo: repo}
}

func (s *OrderService) GetAllOrders(limit, offset int) ([]model.Order, error) {
    return s.Repo.FindAll(limit, offset)
}

func (s *OrderService) GetOrderByID(id uint) (*model.Order, error) {
    return s.Repo.FindByID(id)
}

func (s *OrderService) CreateOrder(order *model.Order) error {
    return s.Repo.Create(order)
}

func (s *OrderService) UpdateOrder(order *model.Order) error {
    return s.Repo.Update(order)
}

func (s *OrderService) DeleteOrder(id uint) error {
    return s.Repo.Delete(id)
}

func (s *OrderService) SearchOrders(query string, limit, offset int) ([]model.Order, error) {
    return s.Repo.Search(query, limit, offset)
}
