package server

import (
	"kevinku/go-forum/app/controller"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func Route(ginServer *gin.Engine) {
	ginServer.GET("doc/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	var auth = ginServer.Group("/auth")
	{
		auth.POST("/sign-up", controller.Register)
		auth.POST("/login", middleware.JWTMiddleware().LoginHandler)
	}

	var api = ginServer.Group("/api", middleware.JWTMiddleware().MiddlewareFunc())
	{
		var v1 = api.Group("/v1")
		{
			v1.GET("/test", controller.Test())

			var user = v1.Group("/users")
			{
				user.GET("", middleware.PermissionCheck(model.ROLE_ADMINISTRATOR), controller.GetUsers)
				user.GET("/:id", controller.GetUserByID)

				user.PATCH("/ban/:id", middleware.PermissionCheck(model.ROLE_MANAGER), controller.BanUser)
			}
		}
	}
}
