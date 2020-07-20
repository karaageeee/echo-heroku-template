package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/karaageeee/echo-heroku-template/conf"
	"github.com/karaageeee/echo-heroku-template/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUser is a sample to return user
func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user := model.User{Model: gorm.Model{ID: uint(id)}}
		tx := c.Get(conf.TxKey).(*gorm.DB)
		if err := user.Get(tx).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return echo.NewHTTPError(http.StatusNotFound, "User not found.")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "System Error.")
		}
		return c.JSON(http.StatusOK, user)
	}
}
