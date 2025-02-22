package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/database"
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/utils"
	"net/http"
	"path/filepath"
)

var PASSWORD string

func InitRouter(password string) *gin.Engine {
	PASSWORD = password

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

// POST /api/health
func Health(c *gin.Context) {
	var body HealthRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if body.Type != "HEALTH_CHECK" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type must be HEALTH_CHECK"})
	}
	host, err := database.SelectHost(body.Hostname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if body.Secret != host.Secret {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong secret"})
	}

	c.JSON(200, gin.H{
		"type":          "success",
		"authenticated": true,
	})
}

// POST /api/register
func Register(c *gin.Context) {
	var body RegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if body.Hostname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hostname is required"})
	}
	if body.Password != PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
	}

	secret := utils.GenerateSecret()
	err := database.CreateHost(body.Hostname, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"type":   "success",
		"secret": secret,
	})
}

// POST /api/screenshot/:hostname
func Screenshot(c *gin.Context) {
	host, err := database.SelectHost(c.Param("hostname"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	secret := c.Request.Header["X-Secret"][0]
	if host.Secret != secret {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong secret"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "data/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(200, gin.H{})
}
