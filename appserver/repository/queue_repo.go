package repository

import (
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func CreateQueue(newQueue schema.Queue) schema.Queue {
	Connect()
	DB.Create(&newQueue)
	return newQueue
}

// func UpdateQueue(userID int, updateUser schema.Queue) schema.Queue {
// 	var user schema.Queue
// 	DB.Model(&user).Where("id = ?", userID).Updates(updateUser)
// 	return updateUser
// }

func GetQueues() []schema.Queue {
	Connect()
	var Queues []schema.Queue
	DB.Find(&Queues)
	return Queues
}
