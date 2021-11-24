package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	errorStatus = "fail"
)

// HandleError for handle API error
type HandleError struct {
	// error message
	Message string `json:"message"`
	// error code:  * `3000` - Internal error
	Code float64 `json:"code"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

// ResponseError for api fail handle
func ResponseError(c *gin.Context, request interface{}, reason string, httpStatusCode int) {
	resp := ErrorResponse{
		Status: errorStatus,
		Reason: reason,
	}
	logrus.Error(c, request, resp)
	c.JSON(httpStatusCode, resp)
}
