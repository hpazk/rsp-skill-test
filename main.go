package main

import (
	"fmt"

	database "github.com.com/hpazk/rsp-skill-test/app/config"
	router "github.com/hpazk/rsp-skill-test/app/router"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitDatabase is...
func InitDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed connect to database ")
	}
	fmt.Println("Database successfully connected")
}

func main() {
	app := echo.New()
	InitDatabase()
	defer database.DBConn.Close()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	router.Router(app)

	// server := echo.New()
	// server.Any("/*", func(c echo.Context) (err error) {
	// 	req := c.Request()
	// 	res := c.Response()
	// 	if req.URL.Path[:4] == "/api" {
	// 		app.ServeHTTP(res, req)
	// 	}

	// 	return
	// })

	app.Logger.Fatal(app.Start(":9000"))
}
