// model/order.go
package model

import "github.com/jinzhu/gorm"

type Order struct {
    gorm.Model
    CustomerID uint   `json:"customer_id"`
    Status     string `json:"status"`
    Total      uint   `json:"total"`
}
