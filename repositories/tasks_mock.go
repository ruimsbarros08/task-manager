package repositories

import (
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) Save(task models.Task) (models.Task, error) {
	ret := m.Called(task)

	if ret.Get(0) != nil {
		task = ret.Get(0).(models.Task)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return task, r1
}


func (m *TaskRepositoryMock) FindAll() ([]models.Task, error) {
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

func (m *TaskRepositoryMock) FindByTechnicianId(technicianId uint) ([]models.Task, error) {
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
