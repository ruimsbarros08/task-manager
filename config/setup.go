package main

import "github.com/ruimsbarros08/task-manager/services"

func main() {
	databaseService := services.Database{}
	database := databaseService.ConnectDatabase()
	databaseService.Migrate(database)
	databaseService.Seed(database)
}
