package appserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manojpethe/itsm-service/appserver/modules/projects"
	"github.com/manojpethe/itsm-service/appserver/modules/users"
)

func StartServer() {

	// var errDBconnect error
	// db, errDBconnect = gorm.Open(sqlite.Open(schema.Databasefile), &gorm.Config{})
	// if errDBconnect != nil {
	// 	panic("failed to connect database")
	// }

	// db.AutoMigrate(&schema.User{})
	// db.AutoMigrate(&schema.Project{})
	// db.AutoMigrate(&schema.Queue{})
	// db.AutoMigrate(&schema.QueueUserMap{})

	// Initialize a Gin router with default middleware (Logger and Recovery)
	appServer := gin.Default()

	appServer.LoadHTMLGlob("html/*")

	appServer.GET("/", func(c *gin.Context) {
		// Render the index.html template with dynamic data
		c.HTML(200, "index.html", gin.H{})
	})

	// Define a simple GET route
	appServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// API (Application Programming Interface)

	// users endoint config
	appServer.GET("/api/users", users.GetUsers)
	appServer.GET("/api/users/:id", users.GetUser)
	appServer.POST("/api/users", users.CreateUser)
	appServer.POST("/api/users/auth", users.AuthUser)

	// projects endoint config
	appServer.GET("/api/projects", projects.GetProjects)
	appServer.POST("/api/projects", projects.CreateProject)

	// Start and run the server on localhost:8080
	appServer.Run()
}
