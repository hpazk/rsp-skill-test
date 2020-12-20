package controllers

// // BookStore Controller is...
// func BookStore(ctx echo.Context) error {
// 	book := models.Book{}
// 	book.Author = ctx.FormValue("author")
// 	book.Title = ctx.FormValue("title")
// 	book.Rating = ctx.FormValue("rating")

// 	res := models.BookInsert(&book)
// 	if res {
// 		return ctx.JSON(201, book)
// 	}
// 	return ctx.JSON(400, "Bad Request")
// }

// // BookList is...
// func BookList(ctx echo.Context) error {
// 	books := models.BookList()
// 	if books != nil {
// 		return ctx.JSON(200, books)
// 	}
// 	return ctx.JSON(204, "No Content")
// }
