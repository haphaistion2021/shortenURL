package api

import (
	"short-url/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// CORS setting
func setCORS(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Accept", "Authorization", "Content-Type", "city"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func setupFrameworkSetting(engine *gin.Engine) {
	// CORS
	setCORS(engine)

	// Request ID
	engine.Use(requestid.New())
}

// SetRoute for route definition
func SetRoute(engine *gin.Engine, config *config.Configure) {
	setupFrameworkSetting(engine)

	// init route
	engine.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
