package api

import (
	"log"
	"short-url/config"
	"short-url/controller"
	"short-url/domain"
	"short-url/middleware"
	"short-url/model"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func setupFrameworkSetting(engine *gin.Engine) {
	// CORS setting
	engine.Use(middleware.Cors())
	// Request ID
	engine.Use(requestid.New())
}

func setupInstance(engine *gin.Engine, config *config.Configure) *controller.Controller {
	// init Redis
	redisPool, err := model.InitRedis(config)
	if err != nil {
		log.Println("redis pool cannot be created")
	}

	// init Database
	db := model.InitDB(config)
	log.Println(db) // FIXME: tmp print, need remove

	// init base controller
	baseController := controller.Controller{
		Config: config,
		Accessors: domain.Accessors{
			Paste: &domain.Paste{},
		},
		RedisPool: *redisPool,
	}

	return &baseController
}

// SetRoute for route definition
func SetRoute(engine *gin.Engine, config *config.Configure) {
	setupFrameworkSetting(engine)

	// init controller
	baseCon := *setupInstance(engine, config)
	userCon := controller.UserController{Controller: &baseCon}

	// utility
	engine.GET("/api/ping", userCon.APIPing)
}
