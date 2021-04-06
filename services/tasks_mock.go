package services

import (
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/stretchr/testify/mock"
)

type TaskServiceMock struct {
	mock.Mock
}

func (m *TaskServiceMock) CreateTask(summary string, user models.User) (models.Task, error) {
	ret := m.Called(summary, user)
	var task models.Task

	if ret.Get(0) != nil {
		task = ret.Get(0).(models.Task)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return task, r1
}


func (m *TaskServiceMock) GetAllTasks() ([]models.Task, error) {
	ret := m.Called()
	var tasks []models.Task

	if ret.Get(0) != nil {
		tasks = ret.Get(0).([]models.Task)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return tasks, r1
}

func (m *TaskServiceMock) GetTechnicianTasks(technicianId uint) ([]models.Task, error) {
	ret := m.Called(technicianId)
	var tasks []models.Task

	if ret.Get(0) != nil {
		tasks = ret.Get(0).([]models.Task)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return tasks, r1
}

