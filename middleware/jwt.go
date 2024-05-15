package middleware

import (
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/app/service"
	"kevinku/go-forum/config"
	. "kevinku/go-forum/lib/logger"
	"strconv"
	"time"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const identityKey string = "id"

// User Login
// @Summary user login
// @tags auth
// @Accept  json
// @Param request body model.Login true "login data"
// @Produce json
// @Success 200 {object} string
// @Failure 401 {object} string
// @Router /auth/login [post]
func JWTMiddleware() (middleware *ginjwt.GinJWTMiddleware) {
	var err error
	var authMiddleware *ginjwt.GinJWTMiddleware

	if authMiddleware, err = ginjwt.New(&ginjwt.GinJWTMiddleware{
		Realm:       "go-forum",
		Key:         []byte(config.Cfg.JWT.SecretKey),
		Timeout:     2 * time.Hour,
		MaxRefresh:  2 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data any) ginjwt.MapClaims {
			if user, ok := data.(*model.User); ok {
				return ginjwt.MapClaims{
					identityKey: strconv.FormatInt(user.ID, 10),
				}
			}
			return ginjwt.MapClaims{}
		},
		IdentityHandler: func(ctx *gin.Context) any {
			var claims = ginjwt.ExtractClaims(ctx)
			id, _ := strconv.ParseInt(claims[identityKey].(string), 10, 64)
			return &model.User{
				ID: id,
			}
		},
		Authenticator: func(ctx *gin.Context) (any, error) {
			var err error
			var login = &model.Login{}
			if err = ctx.ShouldBindJSON(login); err != nil {
				return nil, ginjwt.ErrMissingLoginValues
			}

			var user *model.User
			if user, err = service.Login(login); err != nil {
				return nil, ginjwt.ErrFailedAuthentication
			}
			return user, nil
		},
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
