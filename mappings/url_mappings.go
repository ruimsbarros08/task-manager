package mappings

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/controllers"
)

type Mappings struct {
	UserController controllers.UsersController
	TaskController controllers.TasksController
}

func (m *Mappings)CreateUrlMappings() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.POST("/login", m.UserController.Login)
		users.POST("/register", m.UserController.CreateUser)
	}

	tasks := router.Group("/tasks")
	{
		tasks.POST("/", m.TaskController.CreateTask)
		tasks.GET("/", m.TaskController.ListTasks)
	}

	return router
}
