package models

import (
	"github.com/jinzhu/gorm"
)

// Book struct is...
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}
