package main

import (
	"github.com/hpazk/rsp-skill-test/app/database"
	router "github.com/hpazk/rsp-skill-test/app/router"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	database.InitDatabase()
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
