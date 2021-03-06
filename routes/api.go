package routes

import (
	"github.com/hpazk/rsp-skill-test/routes/handler"
	"github.com/labstack/echo/v4"
)

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("x-version", "Test/v1.0")
		return next((ctx))
	}
}

// Router is...
func Router(e *echo.Echo) {

	api := e.Group("/api/v1", serverHeader)
	api.GET("/register", handler.CreateUser)
	api.GET("/user/login", handler.Auth)
	api.GET("/rooms", handler.GetRooms)
	api.GET("/bookings", handler.GetBookings)

	// e.GET("/register", func(ctx echo.Context) (err error) {
	// 	return ctx.String(http.StatusOK, "Hello World!")
	// })
}
