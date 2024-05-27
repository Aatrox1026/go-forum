package middleware

import (
	"fmt"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/app/service"
	. "kevinku/go-forum/lib/logger"
	"net/http"
	"strconv"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PermissionCheck(role int64) (handler gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		var result bool
		var err error

		var claims = ginjwt.ExtractClaims(ctx)

		if result, err = roleCheck(claims, role); err != nil {
			Logger.Error("role check failed", zap.Any("error", err))
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{"error": fmt.Sprintf("role check failed: %v", err)})
			return
		}
		if !result {
			ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]any{"error": "permission denied"})
			return
		}
	}
}

func roleCheck(claims ginjwt.MapClaims, role int64) (result bool, err error) {
	var id int64
	if tmp, ok := claims["id"].(string); !ok {
		Logger.Error("invalid token claims", zap.Any("claims", claims))
		return false, fmt.Errorf("invalid token: missing user id")
	} else if id, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		Logger.Error("invalid user id", zap.String("id", tmp), zap.Any("error", err))
	}

	var user *model.User = new(model.User)
	if _, user, err = service.GetUserByID(id); err != nil {
		return false, err
	}

	if user.Role > role {
		return false, nil
	}
	return true, nil
}
