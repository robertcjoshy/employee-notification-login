package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func OpenDatabaseConnection() {
	var err error
	// //host := os.Getenv("POSTGRES_HOST")
	// username := os.Getenv("POSTGRES_USER")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// databaseName := os.Getenv("POSTGRES_DATABASE")
	// port := os.Getenv("POSTGRES_PORT")
	username := "postgres"
	password := "1234"
	databaseName := "startup"
	ssl := "disable"
	zone := "Asia/Shanghai"

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode = %s TimeZone=%s", username, password, databaseName, ssl, zone)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}
}

func AutoMigrateModels() {
	Database.AutoMigrate(&User{})
}
