package server

import (
	"kevinku/go-forum/app/controller"
	"kevinku/go-forum/docs"
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

	Route(engine)

	return &Server{
		engine: engine,
	}
}

func Route(ginServer *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginServer.GET("doc/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	var api = ginServer.Group("/api")
	{
		var v1 = api.Group("/v1")
		{
			var auth = ginServer.Group("/auth")
			{
				auth.POST("/sign-up", controller.SignUp)
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
