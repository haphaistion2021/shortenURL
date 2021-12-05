package controller

import (
	"short-url/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
}

// @Tags Test
// @Summary test service API
// @version 1.0
// @Accept  json
// @Produce application/json
// @Router /api/ping [get]
// @Success 200 {object} service.PingResponse
func (u *UserController) APIPing(c *gin.Context) {
	resp := service.PingResponse{
		Message: "pong",
	}

	u.ResponseOK(c, "", resp)
}
