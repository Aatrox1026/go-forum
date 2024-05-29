package controller

import (
	"fmt"
	l "kevinku/go-forum/lib/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Json = map[string]any

var f = fmt.Sprintf
var logger = l.Logger

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
			Json{"msg": data.(error).Error()},
		)
	default:
		logger.Error(
			"invalid status code",
			zap.Int("code", statusCode))
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			Json{"msg": f("invalid status code: %d", statusCode)},
		)
	}
}
