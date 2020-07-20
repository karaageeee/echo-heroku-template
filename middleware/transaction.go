package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/karaageeee/echo-heroku-template/conf"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// TransactionHandler manage db session every request
func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx := db.Begin()

			c.Set(conf.TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				log.Debug("Transction Rollback: ", err)
				return err
			}
			log.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}
