package server

import (
	"kevinku/go-forum/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() (server *Server) {
	var engine = gin.New()
	engine.Use(
		middleware.GinLogger(middleware.Logger),
		middleware.GinRecovery(middleware.Logger, true),
	)

	return &Server{
		engine: engine,
	}
}

func Route(ginServer *gin.Engine) {
	ginServer.GET("doc/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	var api = ginServer.Group("/api")
	{
		api.GET("/ping")
		// 	var v1 = api.Group("/v1")
		// 	{

		// 	}
	}
}

func (server *Server) Run(addr string) {
	server.engine.Run(addr)
}
