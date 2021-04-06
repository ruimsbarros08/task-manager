package main

import (
	"github.com/adjust/rmq/v3"
	"github.com/ruimsbarros08/task-manager/services"
	"log"
	"time"
)

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
)

func main() {

	redisService := services.RedisService{}
	redisService.ConnectRedis("consumer")
	q := redisService.OpenQueue("tasks")
	taskHandler := services.TaskHandler{Q: q}

	if err := q.StartConsuming(prefetchLimit, pollDuration); err != nil {
		panic(err)
	}

	log.Print("Consuming...")
	consumer, err := q.AddConsumerFunc("consumer", func(delivery rmq.Delivery) {
		taskHandler.HandleNotification(delivery.Payload())
		delivery.Ack()
	})

	if err!= nil {
		panic(err)
	}

	log.Printf("Consumer %s finished\n", consumer)
}
