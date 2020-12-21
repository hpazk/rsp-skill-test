package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TrailingSlash is...
func TrailingSlash() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}
