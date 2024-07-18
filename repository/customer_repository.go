// repository/customer_repository.go
package repository

import (
    "kstyle-hub/model"

    "github.com/jinzhu/gorm"
)

type CustomerRepository struct {
    DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
    return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) FindAll(limit, offset int) ([]model.Customer, error) {
    var customers []model.Customer
    if err := r.DB.Limit(limit).Offset(offset).Find(&customers).Error; err != nil {
        return nil, err
    }
    return customers, nil
}

func (r *CustomerRepository) FindByID(id uint) (*model.Customer, error) {
    var customer model.Customer
    if err := r.DB.First(&customer, id).Error; err != nil {
        return nil, err
    }
    return &customer, nil
}

func (r *CustomerRepository) Create(customer *model.Customer) error {
    return r.DB.Create(customer).Error
}

func (r *CustomerRepository) Update(customer *model.Customer) error {
    return r.DB.Save(customer).Error
}

func (r *CustomerRepository) Delete(id uint) error {
    return r.DB.Delete(&model.Customer{}, id).Error
}

func (r *CustomerRepository) Search(query string, limit, offset int) ([]model.Customer, error) {
    var customers []model.Customer
    if err := r.DB.Where("name LIKE ?", "%"+query+"%").Limit(limit).Offset(offset).Find(&customers).Error; err != nil {
        return nil, err
    }
    return customers, nil
}
