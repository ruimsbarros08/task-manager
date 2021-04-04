package mappings

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/controllers"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	users := Router.Group("/users")
	{
		users.POST("/login", controllers.Login)
		users.POST("/register", controllers.CreateUser)
	}

	tasks := Router.Group("/tasks")
	{
		tasks.POST("/", controllers.CreateTask)
		tasks.GET("/", controllers.ListTasks)
	}
}
