package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/services"
	"net/http"
)

type TasksController struct {
	TaskService services.TaskService
	AuthService services.AuthenticationService
	UserService services.UserService
	TaskHandler services.TaskHandler
}

type TaskInput struct {
	Summary string `json:"summary" binding:"required"`
}

func (ctr *TasksController) CreateTask(c *gin.Context) {
	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bearToken := c.Request.Header.Get("Authorization")
	token := ctr.AuthService.ExtractTokenFromHeader(bearToken)

	userId, err := ctr.AuthService.ExtractTokenMetadata(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := ctr.UserService.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	role, err := ctr.UserService.UserHasRole(user, "Technician")
	if !role {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not technician"})
		return
	}

	task, err := ctr.TaskService.CreateTask(input.Summary, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
	ctr.TaskHandler.NotifyManagers(task)
}

func (ctr *TasksController) UpdateTask(c *gin.Context) {
	//TODO

}

func (ctr *TasksController) ListTasks(c *gin.Context) {
	//TODO pagination
	bearToken := c.Request.Header.Get("Authorization")
	token := ctr.AuthService.ExtractTokenFromHeader(bearToken)

	userId, err := ctr.AuthService.ExtractTokenMetadata(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := ctr.UserService.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	role, err := ctr.UserService.UserHasRole(user, "Manager")
	if role {
		tasks, err := ctr.TaskService.GetAllTasks()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": tasks})
		return
	}

	role, err = ctr.UserService.UserHasRole(user, "Technician")
	if role {
		tasks, err := ctr.TaskService.GetTechnicianTasks(user.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": tasks})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "User has no permission"})
}

func (ctr *TasksController) DeleteTask(c *gin.Context) {
	//TODO
}
