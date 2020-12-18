package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Context
var DB *gorm.DB

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
