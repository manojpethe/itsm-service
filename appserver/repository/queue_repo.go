package repository

import (
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func CreateQueue(newQueue schema.Queue) schema.Queue {
	Connect()
	DB.Create(&newQueue)
	return newQueue
}

func DeleteQueue(queueid int) schema.Queue {
	Connect()
	Queue := schema.Queue{}
	DB.Where("ID = ?", queueid).Unscoped().Delete(&Queue)
	return Queue
}

// func UpdateQueue(userID int, updateUser schema.Queue) schema.Queue {
// 	var user schema.Queue
// 	DB.Model(&user).Where("id = ?", userID).Updates(updateUser)
// 	return updateUser
// }

func GetQueues(projectid string) []schema.Queue {
	Connect()
	var Queues []schema.Queue
	if projectid == "" {
		DB.Find(&Queues)
	} else {
		DB.Where("project_id = ?", projectid).Find(&Queues)
	}
	return Queues
}
