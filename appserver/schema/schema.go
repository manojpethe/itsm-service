package schema

import (
	"gorm.io/gorm"
)

const Databasefile = "itsm.db"

type User struct {
	gorm.Model        // adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Superuser  bool   `json:"superuser"`
	Active     *bool  `json:"active" gorm:"default:true"`
}

type Project struct {
	gorm.Model
	Title string
	Info  string
}

type Queue struct {
	gorm.Model
	Name      string
	ProjectId string
}

type QueueUserMap struct {
	gorm.Model
	UserId  string
	QueueId string
}
