package services

import (
	"github.com/ruimsbarros08/task-manager/models"
	"time"
)

func CreateTask(summary string, user models.User) (models.Task, error) {
	//secret := os.Getenv("ACCESS_SECRET")
	//encryptedSummary := Encrypt(summary, secret)
	task := models.Task{Summary: summary, PerformedAt: time.Now(), UserID: user.ID, UpdatedAt: time.Now()}
	err := models.DB.Create(&task).Error

	return task, err
}
