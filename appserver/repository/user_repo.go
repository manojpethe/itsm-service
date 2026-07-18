package repository

import (
	"fmt"

	"github.com/manojpethe/itsm-service/appserver/schema"
)

func AuthenticateUser(username string, password string) schema.User {
	Connect()
	var foundUser schema.User
	DB.Where("email = ? and password = ?", username, password).First(&foundUser)
	return foundUser
}

func CreateUser(newUser schema.User) schema.User {
	DB.Create(&newUser)
	return newUser
}

func UpdateUser(userID int, updateUser schema.User) schema.User {
	var user schema.User
	fmt.Println(&updateUser.Active)
	DB.Model(&user).Where("id = ?", userID).Updates(updateUser)
	return updateUser
}

func GetUsers() []schema.User {
	Connect()
	var users []schema.User
	DB.Find(&users)
	return users
}
