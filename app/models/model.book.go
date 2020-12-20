package models

import (
	"time"

	"github.com/hpazk/rsp-skill-test/app/database"
	"gorm.io/gorm"
)

// Book struct is...
type Book struct {
	// gorm.Model
	ID        uint
	Name      string
	Age       string
	Address   string
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

// BookInsert is...
func BookInsert(book *Book) bool {
	res := database.DB.Create(&book)

	if res.Error == nil {
		return true
	}

	return false
}

// // BookList is...
// func BookList() []Book {
// 	books := []Book{}
// 	res := database.DBConn.Find(&books)
// 	if res.Error == nil {
// 		return books
// 	}
// 	return nil
// }

// BookFetchAll is...
func BookFetchAll() []Book {
	var books []Book
	// Get all records
	res := database.DBConn().Find(&books)
	if res.Error == nil {
		return books
	}
	return nil
}
