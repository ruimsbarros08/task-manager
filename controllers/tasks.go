package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/models"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	c.Bind(&task)

}

func ListTasks(c *gin.Context) {

}
