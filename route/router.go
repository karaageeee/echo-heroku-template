package route

import (
	"github.com/karaageeee/echo-heroku-template/controller"
	"github.com/karaageeee/echo-heroku-template/db"
	appHandler "github.com/karaageeee/echo-heroku-template/handler"
	appMw "github.com/karaageeee/echo-heroku-template/middleware"
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

// Setup is supponsed to be run to init web app
func Setup() *echo.Echo {

	e := echo.New()

	// Setup DB
	db := db.Setup()

	// Set middleware
	e.Use(echoMw.Recover())
	e.Use(echoMw.LoggerWithConfig(echoMw.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","method":"${method}","uri":"${uri}","status":${status},"error":"${error}"}` + "\n",
	}))
	e.Use(appMw.AuthValidate())
	e.Use(appMw.TransactionHandler(db))

	// Set error handler
	e.HTTPErrorHandler = appHandler.JSONHTTPErrorHandler

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", controller.Hello())
		v1.GET("/users/:id", controller.GetUser())
	}

	return e
}
