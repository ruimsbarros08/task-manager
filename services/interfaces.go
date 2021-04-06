package services

import (
	"github.com/ruimsbarros08/task-manager/models"
)

type TaskServiceInterface interface {
	CreateTask(summary string, user models.User) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTechnicianTasks(technicianId uint) ([]models.Task, error)
}

type TaskHandlerInterface interface {
	NotifyManagers(task models.Task)
	HandleNotification(notification string)
}
