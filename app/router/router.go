package router

import (
	"net/http"

	"github.com/hpazk/rsp-skill-test/app/controllers"
	"github.com/hpazk/rsp-skill-test/app/database"
	"github.com/hpazk/rsp-skill-test/app/middleware"
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
// 	app := echo.New()
// 	api := app.Group("/api/v1", serverHeader)
// 	api.POST("/register", handler.CreateUser)
// 	api.GET("/books", handler.GetBooks)

// 	return app
// }

// Router is...
func Router(app *echo.Echo) {

	app.Pre(middleware.TrailingSlash())

	api := app.Group("/api/v1", serverHeader)

	rooms := api.Group("/rooms")
	{
		rooms.GET("", controllers.FetchAllRooms)
		rooms.GET("/:id", controllers.FetchRoom)
		rooms.POST("", controllers.AddRoom)
	}

	bookings := api.Group("/bookings")
	{
		bookings.GET("", controllers.FetchAllBookings)
		bookings.POST("", controllers.AddBooking)
	}

	// Test
	books := api.Group("/books")
	{
		books.GET("", controllers.BookList)
		books.POST("", controllers.BookStore)

		// Test
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
