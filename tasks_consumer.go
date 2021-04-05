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

	services.ConnectRedis("consumer")
	services.OpenTasksQueue()

	if err := services.Tasks.StartConsuming(prefetchLimit, pollDuration); err != nil {
		panic(err)
	}

	log.Print("Consuming...")
	consumer, err := services.Tasks.AddConsumerFunc("consumer", func(delivery rmq.Delivery) {
		services.HandleNotification(delivery.Payload())
		delivery.Ack()
	})

	if err!= nil {
		panic(err)
	}

	log.Printf("Consumer %s finished\n", consumer)
}
