package queueService

import (
	"github.com/manojpethe/itsm-service/appserver/repository"
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func GetQueueService(projectid string) []schema.Queue {
	return repository.GetQueues(projectid)
}

func CreateQueueService(newQueue schema.Queue) schema.Queue {
	createdQueue := repository.CreateQueue(newQueue)
	return createdQueue
}

func AddUserToQueue(newQUMap schema.QueueUserMap) schema.QueueUserMap {
	createdMap := repository.CreateQUMap(newQUMap)
	return createdMap
}

func GetQueueDetails(queueid int) []schema.QueueUsers {
	queueUsers := repository.GetQueueUsers(queueid)
	return queueUsers
}

func DeleteQueueService(queueid int) schema.Queue {
	Queue := repository.DeleteQueue(queueid)
	return Queue
}

func DelUserFromQueue(QUMap schema.QueueUserMap) schema.QueueUserMap {
	deletedMap := repository.DeleteQUMap(QUMap)
	return deletedMap
}
