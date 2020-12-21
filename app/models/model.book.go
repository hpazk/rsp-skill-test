package models

import (
	"time"

	"gorm.io/gorm"
)

// Book struct is...
type Book struct {
	// gorm.Model
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Age       string         `json:"age"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"  gorm:"time"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"time"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"time"`
	// DeletedAt time.Time `json:"deleted_at"`
}

// type Book struct {
// ID        uint      `json:"id" gorm:"primaryKey"`
// Name      string    `json:"name"`
// Age       string    `json:"age"`
// Address   string    `json:"address"`
// CreatedAt time.Time `json:"created_at"`
// UpdatedAt time.Time `json:"updated_at"`
// DeletedAt time.Time `json:"deleted_at"`
// }
