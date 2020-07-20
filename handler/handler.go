package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// JSONHTTPErrorHandler check errors in context
func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"
	// handle echo.NewHTTPError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = fmt.Sprintf("%v", he.Message)
	}
	if !c.Response().Committed {
		c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}
