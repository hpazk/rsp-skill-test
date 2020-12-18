package database

import (
	"fmt"

	book "github.com/hpazk/rsp-skill-test/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Context
var DB *gorm.DB

// InitDatabase is...
func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed connect to database ")
	}
	fmt.Println("Database successfully connected")

	// Auto Migration
	DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

// Database Initialization
// func Database() *gorm.DB {
// 	driver := os.Getenv("DATABASE_DRIVER")
// 	database := os.Getenv("DATABASE_URL")

// 	var err error
// 	DB, err = gorm.Open(driver, database)

// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Println("Database Connected")

// 	return DB
// }

var (
	// DBConn is...
	DBConn *gorm.DB
)
