// repository/user_repository.go
package repository

import (
	"fmt"
	"kstyle-hub/model"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (*model.Users, error) {
	var user model.Users
	if r.DB == nil {
		return nil, fmt.Errorf("koneksi database tidak tersedia")
	}
	result := r.DB.Where("username = ?", username).Debug().First(&user)
	fmt.Println(result)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.Users) error {
	return r.DB.Create(user).Error
}
