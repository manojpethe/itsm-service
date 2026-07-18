package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	projectService "github.com/manojpethe/itsm-service/appserver/modules/projects/service"
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func GetProjects(c *gin.Context) {
	var response = projectService.GetProjectService()
	c.IndentedJSON(http.StatusOK, response)
}

func CreateProject(c *gin.Context) {
	var newProject schema.Project // or &models.User{}
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var response = projectService.CreateProjectService(newProject)
	c.IndentedJSON(http.StatusCreated, response)
}
