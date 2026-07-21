package queues

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// projectService "github.com/manojpethe/itsm-service/appserver/modules/projects/service"

	queueService "github.com/manojpethe/itsm-service/appserver/modules/queues/service"
	"github.com/manojpethe/itsm-service/appserver/schema"
)

func GetQueues(c *gin.Context) {
	// projectIDStr := r.URL.Query().Get("projectid")
	//     if projectIDStr == "" {
	//         http.Error(w, "projectid query parameter is required", http.StatusBadRequest)
	//         return
	//     }

	var response = queueService.GetQueueService()
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
