package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/manojpethe/itsm-service/appserver/modules/users/service"
)

func AuthUser(c *gin.Context) {
	var credential userService.Credential
	if err := c.BindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response := userService.AuthUserService(credential)
	fmt.Println(response)
	c.IndentedJSON(http.StatusOK, response)
}

func GetUsers(c *gin.Context) {
	var response = userService.GetUsersService()
	c.IndentedJSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	var response = `{message:"get single user endpoint"}`
	c.IndentedJSON(http.StatusOK, response)
}

func CreateUser(c *gin.Context) {
	var response = "{message:\"post user endpoint\"}"
	c.IndentedJSON(http.StatusOK, response)
}
