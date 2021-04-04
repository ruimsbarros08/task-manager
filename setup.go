package main

import "github.com/ruimsbarros08/task-manager/models"

func main() {
	models.ConnectDatabase()
	models.Migrate()
	models.Seed()
}
