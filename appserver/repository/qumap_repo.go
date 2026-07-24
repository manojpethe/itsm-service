package repository

import (
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func CreateQUMap(newQUMap schema.QueueUserMap) schema.QueueUserMap {
	Connect()
	DB.Create(&newQUMap)
	return newQUMap
}

func GetQueueUsers(queueid int) []schema.QueueUsers {
	Connect()
	users := []schema.QueueUsers{}

	query := `
        SELECT id, name 
        FROM users 
        WHERE id IN (
            SELECT user_id 
            FROM queue_user_maps 
            WHERE queue_id = ?
        )`

	DB.Raw(query, queueid).Scan(&users)
	return users
}

func DeleteQUMap(QUMap schema.QueueUserMap) schema.QueueUserMap {
	Connect()
	// fmt.Println(QUMap.QueueID, QUMap.UserID)
	DB.Where("user_id = ? AND queue_id = ?", QUMap.UserID, QUMap.QueueID).Unscoped().Delete(&QUMap)
	return QUMap
}
