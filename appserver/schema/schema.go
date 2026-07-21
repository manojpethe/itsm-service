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
	Title string `json:"title"`
	Info  string `json:"info"`
}

type Queue struct {
	gorm.Model
	Name      string `json:"name"`
	ProjectID uint   `json:"projectid"`
}

type QueueUserMap struct {
	gorm.Model
	UserID  uint
	QueueID uint
}
