package queues

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// projectService "github.com/manojpethe/itsm-service/appserver/modules/projects/service"

	queueService "github.com/manojpethe/itsm-service/appserver/modules/queues/service"
	"github.com/manojpethe/itsm-service/appserver/schema"

	"strconv"
)

func GetQueues(c *gin.Context) {
	// projectIDStr := r.URL.Query().Get("projectid")
	//     if projectIDStr == "" {
	//         http.Error(w, "projectid query parameter is required", http.StatusBadRequest)
	//         return
	//     }

	var projectid = c.Query("projectid")

	var response = queueService.GetQueueService(projectid)
	c.IndentedJSON(http.StatusOK, response)
}

func CreateQueue(c *gin.Context) {
	var newQueue schema.Queue // or &models.User{}
	if err := c.ShouldBindJSON(&newQueue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var response = queueService.CreateQueueService(newQueue)
	c.IndentedJSON(http.StatusCreated, response)
}

func AddUserToQueue(c *gin.Context) {
	var newQUMap schema.QueueUserMap
	if err := c.ShouldBindJSON(&newQUMap); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var response = queueService.AddUserToQueue(newQUMap)
	c.IndentedJSON(http.StatusCreated, response)
}

func GetQueueDetails(c *gin.Context) {
	// 1. Get param from URL (e.g. /queues/:id)
	queueid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid queue id"})
		return // Exits handler without returning a value
	}

	// 2. Fetch details from service layer
	response := queueService.GetQueueDetails(queueid)
	// 3. Send successful response
	c.IndentedJSON(http.StatusOK, response)
}
