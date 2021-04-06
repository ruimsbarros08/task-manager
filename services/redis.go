package services

import (
	"github.com/adjust/rmq/v3"
	"os"
)

type RedisService  struct {
	Connection rmq.Connection
}

func (s *RedisService) ConnectRedis(tag string) {
	connection, err := rmq.OpenConnection(tag, "tcp", os.Getenv("REDIS_URL"), 1, nil)
	if err != nil {
		panic(err)
	}

	s.Connection = connection
}

func (s *RedisService) OpenQueue(name string) rmq.Queue {
	q, err := s.Connection.OpenQueue(name)
	if err != nil {
		panic(err)
	}

	return q
}
