// model/user.go
package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
