package services

import (
	"fmt"
	"github.com/adjust/rmq/v3"
	"github.com/ruimsbarros08/task-manager/models"
	"log"
	"os"
)

var Connection rmq.Connection
var Tasks rmq.Queue

func ConnectRedis(tag string) {
	connection, err := rmq.OpenConnection(tag, "tcp", os.Getenv("REDIS_URL"), 1, nil)
	if err != nil {
		panic(err)
	}

	Connection = connection
}

func OpenTasksQueue() {
	tasks, err := Connection.OpenQueue("tasks")
	if err != nil {
		panic(err)
	}

	Tasks = tasks
}

func NotifyManagers(task models.Task) {
	message := fmt.Sprintf("The technician %d performed the task %d on date %s", task.UserID, task.ID, task.PerformedAt)
	err := Tasks.Publish(message)

	if err != nil {
		panic(err)
	}
}

func HandleNotification(notification string) {
	log.Print(notification)
}
