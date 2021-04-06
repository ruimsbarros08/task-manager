package services

import (
	"github.com/jinzhu/gorm"
	"github.com/ruimsbarros08/task-manager/models"
	"time"
)

type TaskService struct {
	Db *gorm.DB
}

func (s *TaskService) CreateTask(summary string, user models.User) (models.Task, error) {
	//secret := os.Getenv("ACCESS_SECRET")
	//encryptedSummary := Encrypt(summary, secret)
	//TODO encrypt
	task := models.Task{Summary: summary, PerformedAt: time.Now(), UserID: user.ID, UpdatedAt: time.Now()}
	err := s.Db.Create(&task).Error

	return task, err
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := s.Db.Model("Task").Find(&tasks).Error

	return tasks, err
}

func (s *TaskService) GetTechnicianTasks(technicianId uint) ([]models.Task, error) {
	var tasks []models.Task
	err := s.Db.Model("Task").Where("user_id = ?", technicianId).Find(&tasks).Error

	return tasks, err
}
