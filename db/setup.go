package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"
	// Setup might call only once
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/karaageeee/echo-heroku-template/conf"
	"github.com/karaageeee/echo-heroku-template/model"
)

// Setup db connection
func Setup() *gorm.DB {
	log.Info("DB setup start")
	db, err := gorm.Open("postgres", connectionInfo())
	if err != nil {
		log.Fatal(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
		return nil
	}
	log.Info("DB setup done")

	log.Info("DB migration start")
	migrateDB(db)
	log.Info("DB migration done")

	createAdminUser(db)

	return db
}

// ConnectionInfo returns string for gorm.Open
func connectionInfo() string {
	env := os.Getenv("ENV")
	sslmode := "disable"
	if env == "production" || env == "staging" {
		sslmode = "require"
	}
	connInfo := conf.GetDBConnectionInfo()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", connInfo.HOST, connInfo.PORT, connInfo.User, connInfo.Pass, connInfo.Name, sslmode)
}

func migrateDB(db *gorm.DB) {

	// Create/Update Tables
	db.AutoMigrate(&model.User{})

	// Add unique index if necessary
	// db.Model(&model.User{}).AddUniqueIndex("idx_user_id", "id")

}

func createAdminUser(db *gorm.DB) {

	admin := model.User{Name: "Admin"}
	if err := db.Find(&admin, uint(1)).Error; gorm.IsRecordNotFoundError(err) {
		db.Create(&admin)
		log.Info("Admin User has been created")
		return
	}
	log.Info("Admin User exists")
}
