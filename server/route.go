package server

import (
	"context"
	"kevinku/go-forum/app/controller"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/database"
	"kevinku/go-forum/lib/redis"
	"kevinku/go-forum/middleware"
	"log"
	"time"

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
			v1.GET("/test", func(ctx *gin.Context) {
				rdb := redis.Client{Client: *database.RDB}

				log.Println(rdb.Set(context.Background(), "test", model.User{ID: 123, Name: "user", Email: "a@b.com", Password: "pwd", Role: 4}, 30*time.Second))

				var tmp = new(model.User)
				log.Println(rdb.Get(context.Background(), "test", &tmp))
				log.Printf("%+v\n", tmp)
			})

			var user = v1.Group("/user")
			{
				user.GET("")
			}
		}
	}
}
