package database

import (
	"github.com/sudheerpal/learn-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/learngo"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Lead{})
}
