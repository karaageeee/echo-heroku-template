package main

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/joho/godotenv"
)

func init() {

	errDotEnv := godotenv.Load()
	if errDotEnv != nil {
		logrus.WithFields(logrus.Fields{"error": errDotEnv}).Info("Failed to load .env")
	}

	// set up logger
	if os.Getenv("ENV") == "production" {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("$PORT must be set")
	}

	router := router()

	host := ""
	env := os.Getenv("ENV")
	if env == "" || env == "dev" {
		// macで毎回warningが出ないようにする
		host = "localhost"
	}
	logrus.Fatal(router.Start(host + ":" + port))
}

func router() *echo.Echo  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}