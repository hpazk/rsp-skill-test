package handler

import (
	"net/http"

	user "github.com/hpazk/rsp-skill-test/app/models"
	"github.com/labstack/echo/v4"
)

// CreateUser is...
func CreateUser(ctx echo.Context) error {
	user := new(user.User)
	if err := ctx.Bind(&user); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

// // Auth is...
// func Auth(ctx echo.Context) error {
// 	return ctx.String(http.StatusOK, "login")
// }

// // GetRooms is...
// func GetRooms(ctx echo.Context) error {
// 	return ctx.String(http.StatusOK, "rooms")
// }

// // GetBookings is...
// func GetBookings(ctx echo.Context) error {
// 	return ctx.String(http.StatusOK, "bookings")
// }
