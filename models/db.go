package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func Migrate() {
	DB.AutoMigrate(&Role{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Task{})
}

func Seed() {
	managerRole := Role{Name: "Manager"}
	technicianRole := Role{Name: "Technician"}

	DB.Create(&managerRole)
	DB.Create(&technicianRole)
}
