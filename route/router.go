package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Setup is supponsed to be run to init web app
func Setup() *echo.Echo {

	e := echo.New()

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	}

	return e
}
