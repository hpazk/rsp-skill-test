package models

import (
	"github.com/hpazk/rsp-skill-test/app/database"
	"github.com/jinzhu/gorm"
)

// Book struct is...
type Book struct {
	gorm.Model
	Name    string
	Age     uint8
	Address string
}

// // BookInsert is...
// func BookInsert(book *Book) bool {
// 	res := database.DBConn.Create(&book)

// 	if res.Error == nil {
// 		return true
// 	}

// 	return false
// }

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
