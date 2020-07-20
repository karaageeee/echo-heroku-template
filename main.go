package main

import (
	"github.com/joho/godotenv"
	"github.com/karaageeee/echo-heroku-template/route"
	"github.com/sirupsen/logrus"
	"os"
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

	router := route.Setup()

	host := ""
	env := os.Getenv("ENV")
	if env == "" || env == "dev" {
		// macで毎回warningが出ないようにする
		host = "localhost"
	}
	logrus.Fatal(router.Start(host + ":" + port))
}
