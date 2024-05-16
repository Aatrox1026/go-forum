package middleware

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func PermissionCheck(role int64) (handler gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		var claims = ginjwt.ExtractClaims(ctx)

		ctx.JSON(200, claims)
	}
}
