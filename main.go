package main

import (
	"github.com/hpazk/rsp-skill-test/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.Router(app)

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
