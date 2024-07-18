// model/customer.go
package model

import "github.com/jinzhu/gorm"

type Customer struct {
    gorm.Model
    Name    string `json:"name"`
    Email   string `json:"email"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
}
