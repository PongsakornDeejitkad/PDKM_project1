package models

import (
	"fmt"
	"order-management/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		utils.ViperGetString("postgres.host"),
		utils.ViperGetString("postgres.user"),
		utils.ViperGetString("postgres.password"),
		utils.ViperGetString("postgres.dbname"),
		utils.ViperGetString("postgres.port"))

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	// DB.Debug()

	// migrateDB()
	// migrateData()

	return err
}

func migrateDB() {
	DB.AutoMigrate(&User{})
}
