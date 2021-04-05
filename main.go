package main

import (
	"github.com/ruimsbarros08/task-manager/mappings"
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/services"
)

func main() {
	models.ConnectDatabase()
	services.ConnectRedis("producer")
	services.OpenTasksQueue()
	mappings.CreateUrlMappings()
	mappings.Router.Run()
}
