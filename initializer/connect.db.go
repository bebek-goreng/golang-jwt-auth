package initializer

import (
	"fmt"
	"os"

	"github.com/bebek-goreng/golang-jwt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error

	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to conenct db" + err.Error())
	}

	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed migrate data to db" + err.Error())
	}

	fmt.Println("Database connection success")
}
