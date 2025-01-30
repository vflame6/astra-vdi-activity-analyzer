package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/ping", Ping)
	router.POST("/api/health", Health)
	router.POST("/api/register", Register)
	router.POST("/api/screenshot/:hostname", Screenshot)

	return router
}

// POST /api/ping
func Ping(c *gin.Context) {
	var body PingRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if body.Type != "PING" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type must be PING"})
	}
	c.JSON(200, gin.H{
		"type": "PONG",
	})
}

func Health(c *gin.Context) {
	var body HealthRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//

	c.JSON(200, gin.H{
		"type":          "success",
		"authenticated": true,
	})
}

// POST /api/register
func Register(c *gin.Context) {

}

// POST /api/screenshot/:hostname
func Screenshot(c *gin.Context) {
	_ = c.Param("hostname")
}
