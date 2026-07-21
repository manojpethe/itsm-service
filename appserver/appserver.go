package appserver

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manojpethe/itsm-service/appserver/modules/projects"
	"github.com/manojpethe/itsm-service/appserver/modules/queues"
	"github.com/manojpethe/itsm-service/appserver/modules/users"
	// "github.com/manojpethe/itsm-service/appserver/repository"
	// "github.com/manojpethe/itsm-service/appserver/schema"
)

func StartServer() {

	// repository.Connect()

	// repository.DB.AutoMigrate(&schema.User{})
	// repository.DB.AutoMigrate(&schema.Project{})
	// repository.DB.AutoMigrate(&schema.Queue{})
	// repository.DB.AutoMigrate(&schema.QueueUserMap{})

	// Initialize a Gin router with default middleware (Logger and Recovery)
	appServer := gin.Default()

	// enabled for development
	appServer.Use(cors.Default())

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
	appServer.PATCH("/api/users/:id", users.UpdateUser)
	appServer.POST("/api/users", users.CreateUser)
	appServer.POST("/api/users/auth", users.AuthUser)

	// projects endoint config
	appServer.GET("/api/projects", projects.GetProjects)
	appServer.POST("/api/projects", projects.CreateProject)

	appServer.GET("/api/queues", queues.GetQueues)
	appServer.POST("/api/queues", queues.CreateQueue)

	// Start and run the server on localhost:8080
	appServer.Run()
}
