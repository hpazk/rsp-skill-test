package controllers

import (
	"net/http"
	"strconv"

	"github.com/hpazk/rsp-skill-test/app/models"
	"github.com/hpazk/rsp-skill-test/app/models/operation"
	"github.com/labstack/echo/v4"
)

// FetchAllRooms is...
func FetchAllRooms(ctx echo.Context) error {
	result := operation.SelectAllRooms()
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

// FetchRoom is...
func FetchRoom(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(400, "Bad Request")
	}

	result := operation.SelectRoom(id)
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

// AddRoom is...
func AddRoom(ctx echo.Context) error {
	room := new(models.Room)
	if err := ctx.Bind(room); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	result := operation.InsertRoom(room)
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
