package controllers

import (
	"net/http"

	"github.com/hpazk/rsp-skill-test/app/models"
	"github.com/hpazk/rsp-skill-test/app/models/operation"
	"github.com/labstack/echo/v4"
)

// FetchAllBookings is...
func FetchAllBookings(ctx echo.Context) error {
	result := operation.SelectAllBookings()
	response := models.Response{
		Code:    200,
		Data:    result,
		Message: "Success",
	}
	if result != nil {
		return ctx.JSON(200, response)
	}
	return ctx.JSON(204, "No Content")
}

// AddBooking is...
func AddBooking(ctx echo.Context) error {
	booking := new(models.Booking)
	if err := ctx.Bind(booking); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	result := operation.InsertBooking(booking)
	response := models.Response{
		Code:    200,
		Data:    result,
		Message: "success",
	}

	if result {
		return ctx.JSON(http.StatusOK, response)
	}

	return ctx.JSON(400, "Bad Request")
}
