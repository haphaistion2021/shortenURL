package main

import (
	"log"
	"short-url/config"
	api "short-url/route"

	"github.com/gin-gonic/gin"
)

// @title Shorten URL API Gateway
// @version 1.0
// @description API gateway for shortenURL

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	config := &config.Config
	engine := setupServer(config)

	err := engine.Run(*config.Server.Host + ":" + *config.Server.Port)
	if err != nil {
		log.Printf("router run err: %+v\n", err)
	}
}

func setupServer(config *config.Configure) *gin.Engine {
	engine := gin.Default()
	api.SetRoute(engine, config)

	return engine
}
