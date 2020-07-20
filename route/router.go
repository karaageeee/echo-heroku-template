package route

import (
	"github.com/karaageeee/echo-heroku-template/controller"
	"github.com/labstack/echo/v4"
)

// Setup is supponsed to be run to init web app
func Setup() *echo.Echo {

	e := echo.New()

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", controller.Hello())
	}

	return e
}
