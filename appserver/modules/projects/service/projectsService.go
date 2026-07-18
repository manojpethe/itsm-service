package projectService

import (
	"github.com/manojpethe/itsm-service/appserver/repository"
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func GetProjectService() []schema.Project {
	return repository.GetProjects()
}

func CreateProjectService(newProject schema.Project) schema.Project {
	createdProject := repository.CreateProject(newProject)
	return createdProject
}
