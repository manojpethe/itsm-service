package repository

import "github.com/manojpethe/itsm-service/appserver/schema"

func AuthenticateUser(username string, password string) schema.User {

	Connect()

	var foundUser schema.User

	DB.Where("email = ? and password = ?", username, password).First(&foundUser)
	// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	return foundUser

}
