package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateUser is...
func CreateUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "crate user")
}

// Auth is...
func Auth(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "login")
}

// GetRooms is...
func GetRooms(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "rooms")
}

// GetBookings is...
func GetBookings(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "bookings")
}
