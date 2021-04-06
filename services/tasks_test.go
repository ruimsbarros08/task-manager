package services

import (
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskService_CreateTask(t *testing.T) {

	t.Run("Create", func(t *testing.T) {
		mockTaskRepository := new(repositories.TaskRepositoryMock)

		taskService := TaskService{
			TasksRepository: mockTaskRepository,
		}

		now := time.Now()
		task := models.Task{
			ID:          1,
			UserID:      1,
			Summary:     "test",
			PerformedAt: now,
			UpdatedAt: now,
		}

		user := models.User{
			ID: 1,
			Name: "test user",
			Email: "test@test.com",
			Password: "test",
			CreatedAt: now,
			UpdatedAt: now,
		}
		mockTaskRepository.On("Save", task).Return(task, nil)

		newTask, err := taskService.CreateTask("test", user)

		assert.Nil(t, err)
		assert.Equal(t, "test", newTask.Summary)

		mockTaskRepository.AssertExpectations(t)
	})

}