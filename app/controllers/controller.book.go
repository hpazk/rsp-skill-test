package controllers

import (
	"github.com/hpazk/rsp-skill-test/app/models"
	"github.com/labstack/echo/v4"
)

// BookStore Controller is...
func BookStore(ctx echo.Context) error {
	book := models.Book{}
	book.Name = ctx.FormValue("name")
	book.Age = ctx.FormValue("age")
	book.Address = ctx.FormValue("address")

	res := models.BookInsert(&book)
	if res {
		return ctx.JSON(201, book)
	}
	return ctx.JSON(400, "Bad Request")
}

// BookList is...
func BookList(ctx echo.Context) error {
	books := models.BookFetchAll()
	response := models.Response{Code: 200, Message: "Success", Data: books}
	if books != nil {
		return ctx.JSON(200, response)
	}
	return ctx.JSON(204, "No Content")
}
