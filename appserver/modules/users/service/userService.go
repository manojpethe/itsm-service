package userService

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model // adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string
	Password   string
	Phone      string
	Email      string
	Superuser  bool
	Active     bool
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsersService() []User {
	users := []User{
		{Name: "Alice", Phone: "92929292"},
		{Name: "Bob", Phone: "92929292"},
	}
	return users
}

func AuthUserService(credential Credential) User {
	fmt.Println("Debug...", credential.Username, credential.Password)
	user := User{Name: "Auth user", Active: true}
	return user
}
