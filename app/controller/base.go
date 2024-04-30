package controller

import (
	. "kevinku/go-forum/lib/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleResponse(ctx *gin.Context, statusCode int, data any) {
	switch statusType := statusCode / 100; statusType {
	case 1, 2, 3:
		ctx.JSON(
			statusCode,
			data,
		)
	case 4, 5:
		ctx.AbortWithStatusJSON(
			statusCode,
			data,
		)
	default:
		Logger.Error(
			"invalid status code",
			zap.Int("code", statusCode))
	}
}
