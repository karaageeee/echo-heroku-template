package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello is a sample controller
func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Process here
		text := "Hello, World!"

		// error sample
		// return echo.NewHTTPError(http.StatusOK, "Faild to map request")

		return c.String(http.StatusOK, text)
	}
}
