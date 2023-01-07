package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/stickpro/vpn-sass/pkg/logger"
)

type response struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
