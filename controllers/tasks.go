package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/services"
	"net/http"
)

type TaskInput struct {
	Summary     string    `json:"summary" binding:"required"`
}

func CreateTask(c *gin.Context) {
	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bearToken := c.Request.Header.Get("Authorization")
	token := services.ExtractTokenFromHeader(bearToken)

	userId, err := services.ExtractTokenMetadata(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	task, err := services.CreateTask(input.Summary, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	c.Bind(&task)

}

func ListTasks(c *gin.Context) {

}

func DeleteTask(c *gin.Context) {

}
