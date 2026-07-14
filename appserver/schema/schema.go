package schema

import (
	"gorm.io/gorm"
)

const Databasefile = "itsm.db"

type User struct {
	gorm.Model // adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string
	Password   string
	Phone      string
	Email      string
	Superuser  bool
	Active     bool
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
