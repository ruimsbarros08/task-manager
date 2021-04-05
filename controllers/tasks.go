package controllers

import (
	"github.com/gin-gonic/gin"
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

	role, err := services.UserHasRole(user, "Technician")
	if role {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not technician"})
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
	//TODO

}

func ListTasks(c *gin.Context) {
	//TODO pagination
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

	role, err := services.UserHasRole(user, "Manager")
	if role {
		tasks, err := services.GetAllTasks()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": tasks})
		return
	}

	role, err = services.UserHasRole(user, "Technician")
	if role {
		tasks, err := services.GetTechnicianTasks(user.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": tasks})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "User has no permission"})
}

func DeleteTask(c *gin.Context) {
	//TODO
}
