package handler

import (
	"net/http"

	"github.com/hpazk/rsp-skill-test/app/database"
	book "github.com/hpazk/rsp-skill-test/app/models"
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

// Test

// GetBooks is...
func GetBooks(ctx echo.Context) error {
	db := database.DBConn
	var books []book.Book
	db.Find((&books))
	return ctx.JSON(http.StatusOK, books)
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
