package server

import (
	"kevinku/go-forum/app/controller"
	"kevinku/go-forum/config"
	l "kevinku/go-forum/lib/logger"
	"kevinku/go-forum/middleware"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() (server *Server) {
	accessLogFile, _ := os.OpenFile(config.Cfg.LogFile.AccessLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	gin.DefaultWriter = accessLogFile

	var engine = gin.New()
	engine.Use(
		middleware.GinLogger(l.Logger),
		middleware.GinRecovery(l.Logger, true),
	)

	Route(engine)

	return &Server{
		engine: engine,
	}
}

func Route(ginServer *gin.Engine) {
	ginServer.GET("doc/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	var api = ginServer.Group("/api")
	{
		var v1 = api.Group("/v1")
		{
			var auth = v1.Group("/auth")
			{
				auth.POST("/sign-up", controller.Register)
				auth.POST("/login")
			}

			var user = v1.Group("/user")
			{
				user.GET("")
			}
		}
	}
}

func (server *Server) Run(addr string) {
	server.engine.Run(addr)
}
