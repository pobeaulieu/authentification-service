package database

import (
	"backend/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load("environment.env")
	// garder DB_prod sur branche master
	dbURL := os.Getenv("DB_prod")
	connection, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Client{}, &models.PasswordPolicy{}, &models.LoginPolicy{}, &models.Log{})
}
