package database

import (
	"github.com/DenilsonNil/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	connectionString := "host=localhost user=root dbname=gin_api_rest password=root sslmode=disable port=5433"

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic("failed to connect database")
	} else {
		println("Database connected")
	}

	DB.AutoMigrate(&models.Aluno{})
}
