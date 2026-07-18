package userService

import (
	"github.com/manojpethe/itsm-service/appserver/repository"

	"github.com/manojpethe/itsm-service/appserver/schema"
)

// type User struct {
// 	gorm.Model // adds ID, CreatedAt, UpdatedAt, DeletedAt fields
// 	Name       string
// 	Password   string
// 	Phone      string
// 	Email      string
// 	Superuser  bool
// 	Active     bool
// }

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsersService() []schema.User {
	var users = repository.GetUsers()
	return users
}

func AuthUserService(credential Credential) schema.User {
	var foundUser = repository.AuthenticateUser(credential.Username, credential.Password)
	return foundUser
}

func CreateUser(newUser schema.User) schema.User {
	// var createdUser schema.User
	createdUser := repository.CreateUser(newUser)
	return createdUser
}

func UpdateUser(userID int, updateUser schema.User) schema.User {
	updatedUser := repository.UpdateUser(userID, updateUser)
	return updatedUser
}
