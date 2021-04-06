package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/ruimsbarros08/task-manager/models"
)

type TasksRepository struct {
	Db *gorm.DB
}

func (r *TasksRepository) Save(task models.Task) (models.Task, error) {
	err := r.Db.Create(&task).Error

	return task, err
}

func (r *TasksRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.Db.Model("Task").Find(&tasks).Error

	return tasks, err
}

func (r *TasksRepository) FindByTechnicianId(technicianId uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.Db.Model("Task").Where("user_id = ?", technicianId).Find(&tasks).Error

	return tasks, err
}