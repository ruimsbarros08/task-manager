package services

import (
	"github.com/ruimsbarros08/task-manager/models"
	"github.com/ruimsbarros08/task-manager/repositories"
	"time"
)

type TaskService struct {
	TasksRepository repositories.TaskRepositoryInterface
}

func (s *TaskService) CreateTask(summary string, user models.User) (models.Task, error) {
	//secret := os.Getenv("ACCESS_SECRET")
	//encryptedSummary := Encrypt(summary, secret)
	//TODO encrypt
	t := models.Task{Summary: summary, PerformedAt: time.Now(), UserID: user.ID, UpdatedAt: time.Now()}
	return s.TasksRepository.Save(t)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.TasksRepository.FindAll()
}

func (s *TaskService) GetTechnicianTasks(technicianId uint) ([]models.Task, error) {
	return s.TasksRepository.FindByTechnicianId(technicianId)
}
