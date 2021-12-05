package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ResponseSuccess for api success handle
func ResponseSuccess(c *gin.Context, response interface{}) {
	logrus.Info(c, response)
	c.JSON(http.StatusOK, response)
}
