package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ruimsbarros08/task-manager/models"
	"os"
)

type Database struct {
}

func (d *Database) ConnectDatabase() *gorm.DB {
	database, err := gorm.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		panic("Failed to connect to database!")
	}
	return database
}

func (d *Database) Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
}

func (d *Database) Seed(db *gorm.DB) {
	managerRole := models.Role{Name: "Manager"}
	technicianRole := models.Role{Name: "Technician"}

	db.Create(&managerRole)
	db.Create(&technicianRole)
}
