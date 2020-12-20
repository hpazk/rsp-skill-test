package router

import (
	"net/http"

	"github.com/hpazk/rsp-skill-test/app/controllers"
	"github.com/hpazk/rsp-skill-test/app/database"
	"github.com/hpazk/rsp-skill-test/app/handler"
	"github.com/hpazk/rsp-skill-test/app/models"
	"github.com/labstack/echo/v4"
)

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("x-version", "Test/v1.0")
		return next((ctx))
	}
}

// // Init Routes
// func Init() *echo.Echo {
// 	e := echo.New()
// 	api := e.Group("/api/v1", serverHeader)
// 	api.POST("/register", handler.CreateUser)
// 	api.GET("/books", handler.GetBooks)

// 	return e
// }

// Router is...
func Router(e *echo.Echo) {

	api := e.Group("/api/v1", serverHeader)
	api.POST("/register", handler.CreateUser)
	// api.GET("/user/login", handler.Auth)
	// api.GET("/rooms", handler.GetRooms)
	// api.GET("/bookings", handler.GetBookings)

	// Test
	books := api.Group("/books")
	{
		// books.GET("", func(ctx echo.Context) error {
		// 	var books []models.Book
		// 	// Get all records
		// 	result := database.DBConn().Find(&books)

		// 	// res := models.Response{Code: 200, Data: books, Message: "Success"}
		// 	return ctx.JSON(http.StatusOK, result)
		// })
		books.GET("", controllers.BookList)
		// books.POST("", func(ctx echo.Context) error {
		// 	book := models.Book{Name: "Xor", Age: 18, Address: "Smi"}

		// 	result := database.DBConn().Create(&book) // pass pointer of data to Create

		// 	return ctx.JSON(http.StatusOK, result)
		// })
		books.POST("", controllers.BookStore)
		books.DELETE("", func(ctx echo.Context) error {
			book := models.Book{}
			result := database.DB.Where(models.Book{ID: 47}).Delete(&book)
			return ctx.JSON(http.StatusOK, result)
		})

		books.GET("/softdelete", func(ctx echo.Context) error {
			book := models.Book{}
			result := database.DB.Unscoped().Where(models.Book{ID: 47}).Find(&book)

			return ctx.JSON(http.StatusOK, result)
		})
	}
}
