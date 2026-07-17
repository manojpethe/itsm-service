package userService

import (
	"fmt"

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
	users := []schema.User{
		{Name: "Alice", Phone: "92929292"},
		{Name: "Bob", Phone: "92929292"},
	}
	return users
}

func AuthUserService(credential Credential) schema.User {
	fmt.Println("Debug...", credential.Username, credential.Password)
	var foundUser = repository.AuthenticateUser(credential.Username, credential.Password)
	fmt.Println("Debug service", foundUser)
	return foundUser
}
