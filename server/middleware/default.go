package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func DefaultMiddleware(config interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("this is default middleware with custom config", config)
			next(c)
			return nil
		}
	}
}
