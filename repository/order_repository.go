// repository/order_repository.go
package repository

import (
    "kstyle-hub/model"

    "github.com/jinzhu/gorm"
)

type OrderRepository struct {
    DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{DB: db}
}

func (r *OrderRepository) FindAll(limit, offset int) ([]model.Order, error) {
    var orders []model.Order
    if err := r.DB.Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (r *OrderRepository) FindByID(id uint) (*model.Order, error) {
    var order model.Order
    if err := r.DB.First(&order, id).Error; err != nil {
        return nil, err
    }
    return &order, nil
}

func (r *OrderRepository) Create(order *model.Order) error {
    return r.DB.Create(order).Error
}

func (r *OrderRepository) Update(order *model.Order) error {
    return r.DB.Save(order).Error
}

func (r *OrderRepository) Delete(id uint) error {
    return r.DB.Delete(&model.Order{}, id).Error
}

func (r *OrderRepository) Search(query string, limit, offset int) ([]model.Order, error) {
    var orders []model.Order
    if err := r.DB.Where("status LIKE ?", "%"+query+"%").Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}
