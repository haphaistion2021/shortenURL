package controller

import (
	"net/http"
	"short-url/config"
	"short-url/domain"
	"short-url/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Config    *config.Configure
	Accessors domain.Accessors
	RedisPool model.RedisPool
}

func (control *Controller) ResponseOK(c *gin.Context, req interface{}, response interface{}) {
	logrus.Info(c, response)
	c.JSON(http.StatusOK, response)
}
