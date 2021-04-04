package main

import (
	"github.com/ruimsbarros08/task-manager/mappings"
	"github.com/ruimsbarros08/task-manager/models"
)

func main() {
	models.ConnectDatabase()
	mappings.CreateUrlMappings()
	mappings.Router.Run()
}
