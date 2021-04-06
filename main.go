package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/controllers"
	"github.com/ruimsbarros08/task-manager/mappings"
	"github.com/ruimsbarros08/task-manager/repositories"
	"github.com/ruimsbarros08/task-manager/services"
)

func main() {
	router := inject()

	router.Run()
}

func inject() *gin.Engine {
	databaseService := services.Database{}
	db := databaseService.ConnectDatabase()
	redisService := services.RedisService{}
	redisService.ConnectRedis("producer")
	q := redisService.OpenQueue("tasks")

	userRepository := repositories.UserRepository{Db: db}
	tasksRepository := repositories.TasksRepository{Db: db}
	encryptionService := services.EncryptionService{}
	userService := services.UserService{UsersRepository: userRepository, Db: db, EncryptionService: encryptionService}
	taskService := services.TaskService{TasksRepository: &tasksRepository}
	authService := services.AuthenticationService{UserService: userService, EncryptionService: encryptionService}
	taskHandler := services.TaskHandler{Q: q}
	userController := controllers.UsersController{UserService: userService, AuthService: authService}
	taskController := controllers.TasksController{TaskService: &taskService, TaskHandler: &taskHandler}

	middleware := mappings.Middleware{AuthService: authService, UserService: userService}
	m := mappings.Mappings{TaskController: taskController, UserController: userController, Middleware: middleware}
	return m.CreateUrlMappings()
}
