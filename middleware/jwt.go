package middleware

import (
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/config"
	. "kevinku/go-forum/lib/logger"
	"time"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	identityKey = "id"
)

func JWTMiddleware() (middleware *ginjwt.GinJWTMiddleware) {
	var err error
	var authMiddleware *ginjwt.GinJWTMiddleware

	if authMiddleware, err = ginjwt.New(&ginjwt.GinJWTMiddleware{
		Realm:       "go-forum",
		Key:         []byte(config.Cfg.JWT.SecretKey),
		Timeout:     2 * time.Hour,
		MaxRefresh:  2 * time.Hour,
		IdentityKey: identityKey,
		// PayloadFunc: func(data any) ginjwt.MapClaims {
		// 	if v, ok := data.(); ok {
		// 		var claims = ginjwt.MapClaims{
		// 		}
		// 	}
		// },
		// IdentityHandler: func(ctx *gin.Context) any {
		// 	var claims = ginjwt.ExtractClaims(ctx)
		// 	return &model.User{
		// 		ID: claims[identityKey].(string),
		// 	}
		// },
		// Authenticator: func(c *gin.Context) (any, error) {

		// },
		Authorizator: func(data any, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(ctx *gin.Context, code int, message string) {
			ctx.JSON(
				code,
				gin.H{"status": code, "msg": message},
			)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}); err != nil {
		Logger.Fatal("JWT Error", zap.Any("error", err))
	}

	return authMiddleware
}
