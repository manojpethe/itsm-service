package repository

import "github.com/manojpethe/itsm-service/appserver/schema"

func CreateQUMap(newQUMap schema.QueueUserMap) schema.QueueUserMap {
	Connect()
	DB.Create(&newQUMap)
	return newQUMap
}

func GetQueueUsers(queueid int) []schema.QueueUsers {
	Connect()
	var users []schema.QueueUsers

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
