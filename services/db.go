package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ruimsbarros08/task-manager/models"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) ConnectDatabase() *gorm.DB {
	database, err := gorm.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		panic("Failed to connect to database!")
	}
	return database
}

func (d *Database) Migrate() {
	d.DB.AutoMigrate(&models.Role{})
	d.DB.AutoMigrate(&models.User{})
	d.DB.AutoMigrate(&models.Task{})
}

func (d *Database) Seed() {
	managerRole := models.Role{Name: "Manager"}
	technicianRole := models.Role{Name: "Technician"}

	d.DB.Create(&managerRole)
	d.DB.Create(&technicianRole)
}
