package services

import (
	"fmt"
	"github.com/adjust/rmq/v3"
	"github.com/ruimsbarros08/task-manager/models"
	"log"
)

type TaskHandler struct {
	Q rmq.Queue
}

func (s *TaskHandler) NotifyManagers(task models.Task) {
	message := fmt.Sprintf("The technician %d performed the task %d on date %s", task.UserID, task.ID, task.PerformedAt)
	err := s.Q.Publish(message)

	if err != nil {
		panic(err)
	}
}

func (s *TaskHandler) HandleNotification(notification string) {
	log.Print(notification)
}

