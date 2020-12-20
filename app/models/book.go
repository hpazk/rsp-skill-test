package models

import (
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

// // BookFetchAll is...
// func BookFetchAll() []Book {
// 	books := []Book{}
// 	res := database.DBConn.Debug().Preload("Book").Find(&books)
// 	if res.Error != nil {
// 		return books
// 	}
// 	return nil
// }
