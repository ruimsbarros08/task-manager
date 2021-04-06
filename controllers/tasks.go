package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/models"
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

	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is missing"})
		return
	}
	user := u.(models.User)

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
	tasks, err := ctr.TaskService.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (ctr *TasksController) ListTechnicianTasks(c *gin.Context) {
	//TODO pagination
	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is missing"})
		return
	}
	user := u.(models.User)
	tasks, err := ctr.TaskService.GetTechnicianTasks(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (ctr *TasksController) DeleteTask(c *gin.Context) {
	//TODO
}
