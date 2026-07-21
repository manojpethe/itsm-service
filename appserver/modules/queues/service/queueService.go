package queueService

import (
	"github.com/manojpethe/itsm-service/appserver/repository"
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func GetQueueService() []schema.Queue {
	return repository.GetQueues()
}

func CreateQueueService(newQueue schema.Queue) schema.Queue {
	createdQueue := repository.CreateQueue(newQueue)
	return createdQueue
}
