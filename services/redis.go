package services

import (
	"fmt"
	"github.com/adjust/rmq/v3"
	"github.com/ruimsbarros08/task-manager/models"
	"os"
)

var Connection rmq.Connection
var Tasks rmq.Queue

func ConnectRedis() {
	connection, err := rmq.OpenConnection("producer", "tcp", os.Getenv("REDIS_URL"), 2, nil)
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
