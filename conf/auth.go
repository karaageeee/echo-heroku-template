package conf

import (
	"os"
)

const (
	// AuthTokenKey is used for API request header
	AuthTokenKey string = "API-TOKEN"
)

// AuthUser defines API user
type AuthUser struct {
	Name  string
	Token string
}

// AuthUsers defines slice of user
type AuthUsers []AuthUser

var authUsersInstance AuthUsers

// Init singleton instance of AuthUsers
func init() {
	env := os.Getenv("ENV")
	if env == "production" {
		authUsersInstance = AuthUsers{
			AuthUser{Name: "WebAppUser", Token: "TBD"},
			AuthUser{Name: "SampleAppUser", Token: "TBD"},
		}
	} else {
		authUsersInstance = AuthUsers{
			AuthUser{Name: "WebAppUser", Token: "web-app-user-token"},
			AuthUser{Name: "SampleAppUser", Token: "sample-user-token"},
		}
	}
}

// GetAuthUsers returns singleton instance
func GetAuthUsers() AuthUsers {
	return authUsersInstance
}
