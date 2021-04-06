package services

import (
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/stretchr/testify/mock"
)
type TaskHandlerMock struct {
	mock.Mock
}

func (m *TaskHandlerMock) NotifyManagers(task models.Task) {
	m.Called(task)
}

func (m *TaskHandlerMock) HandleNotification(notification string) {
	m.Called(notification)
}
