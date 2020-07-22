package middleware

import (
	appConf "github.com/karaageeee/echo-heroku-template/conf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
)

// AuthValidate is to check http request.
func AuthValidate() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		AuthScheme: appConf.AuthTokenKey,
		Validator: func(key string, c echo.Context) (bool, error) {
			for _, authUser := range getAuthUsers() {
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

// AuthUser defines API user
type AuthUser struct {
	Name  string
	Token string
}

// AuthUsers defines slice of user
type AuthUsers []AuthUser

var authUsersInstance AuthUsers

// getAuthUsers returns singleton instance
func getAuthUsers() AuthUsers {

	if authUsersInstance == nil {
		authUsersInstance = AuthUsers{
			AuthUser{Name: "WebAppUser", Token: os.Getenv("WEB_APP_USER_API_TOKEN")},
			// AuthUser{Name: "NewUser", Token: "xxxx"},
		}
	}
	return authUsersInstance
}
