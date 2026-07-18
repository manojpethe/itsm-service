package repository

import (
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func CreateProject(newProject schema.Project) schema.Project {
	Connect()
	DB.Create(&newProject)
	return newProject
}

// func UpdateProject(userID int, updateUser schema.Project) schema.Project {
// 	var user schema.Project
// 	DB.Model(&user).Where("id = ?", userID).Updates(updateUser)
// 	return updateUser
// }

func GetProjects() []schema.Project {
	Connect()
	var projects []schema.Project
	DB.Find(&projects)
	return projects
}
