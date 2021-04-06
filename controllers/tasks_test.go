package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/services"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTasksController_CreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		now := time.Now()
		mockTaskResp := models.Task{
			ID:          1,
			Summary:     "test",
			UserID:      1,
			PerformedAt: now,
			UpdatedAt:   now,
		}

		mockUserInput := models.User{
			ID:        1,
			Email:     "test@test.com",
			Name:      "test",
			Password:  "test",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockTaskService := new(services.TaskServiceMock)
		mockTaskHandler := new(services.TaskHandlerMock)
		mockTaskService.On("CreateTask", "test", mockUserInput).Return(mockTaskResp, nil)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", mockUserInput)
		})

		taskController := TasksController{TaskService: mockTaskService, TaskHandler: mockTaskHandler}
		taskController.CreateTask(ctx)

		mockTaskService.AssertExpectations(t)
		mockTaskHandler.AssertExpectations(t)
	})
}
