package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userService "github.com/manojpethe/itsm-service/appserver/modules/users/service"
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func AuthUser(c *gin.Context) {
	var credential userService.Credential
	if err := c.BindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response := userService.AuthUserService(credential)
	if response.ID != 0 {
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, `{message:"Authentication Failure"}`)
	}
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
	var newUser schema.User // or &models.User{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Pass the pointer to the service layer
	var createdUser = userService.CreateUser(newUser)

	c.IndentedJSON(http.StatusCreated, createdUser)
}

func UpdateUser(c *gin.Context) {
	var userID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser schema.User // or &models.User{}
	if err := c.ShouldBindJSON(&existingUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Pass the pointer to the service layer
	var createdUser = userService.UpdateUser(userID, existingUser)

	c.IndentedJSON(http.StatusCreated, createdUser)
}
