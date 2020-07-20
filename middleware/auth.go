package middleware

import (
	appConf "github.com/karaageeee/echo-heroku-template/conf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

// AuthValidate is to check http request.
func AuthValidate() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		AuthScheme: appConf.AuthTokenKey,
		Validator: func(key string, c echo.Context) (bool, error) {
			for _, authUser := range appConf.GetAuthUsers() {
				if authUser.Token == key {
					log.WithFields(log.Fields{"User": authUser.Name}).Info("Request accepted")
					return true, nil
				}
			}
			log.WithFields(log.Fields{"Key": key}).Info("Request rejected")
			return false, nil
		},
	})
}
