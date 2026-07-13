package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var response = `{message:"get project endpoint"}`
	c.IndentedJSON(http.StatusOK, response)
}

func CreateProject(c *gin.Context) {
	var response = "{message:\"post project endpoint\"}"
	c.IndentedJSON(http.StatusOK, response)
}
