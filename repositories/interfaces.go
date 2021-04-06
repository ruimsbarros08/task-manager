package repositories

import "github.com/ruimsbarros08/task-manager/models"

type TaskRepositoryInterface interface {
	Save(task models.Task) (models.Task, error)
	FindAll() ([]models.Task, error)
	FindByTechnicianId(technicianId uint) ([]models.Task, error)
}
