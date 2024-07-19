package model

import "time"

// User represents the login model
// @Description Login model
type Login struct {
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// User represents the register model
// @Description Register model
type Register struct {
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// Customers represents the customers model
// @Description Customers model
type Customers struct {
	ID        uint      `gorm:"primary_key" json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	Phone     string    `json:"phone" example:"123-456-7890"`
	Address   string    `json:"address" example:"123 Main St"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// Orders represents the orders model
// @Description Orders model
type Orders struct {
	ID         uint      `gorm:"primary_key" json:"id" example:"1"`
	CustomerID uint      `json:"customer_id" example:"1"`
	Status     string    `json:"status" example:"pending"`
	Total      float64   `json:"total" example:"99.99"`
	CreatedAt  time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}
