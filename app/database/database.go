package database

import (
	"log"

	"github.com/hpazk/rsp-skill-test/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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

// DB is...
var DB *gorm.DB

// DBConn is...
func DBConn() *gorm.DB {
	var err error
	DB, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		log.Panic(err)
	}
	log.Println("Database Connected")

	DB.AutoMigrate(
		&models.Book{},
		&models.Room{},
		&models.Booking{},
	)

	return DB
}

// // DBMigrate is...
// func DBMigrate() {
// 	DBConn().AutoMigrate(&models.Book{})
// }
