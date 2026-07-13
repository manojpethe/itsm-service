package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var response = `{message:"get user endpoint"}`
	c.IndentedJSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	var response = "{message:\"post user endpoint\"}"
	c.IndentedJSON(http.StatusOK, response)
}
