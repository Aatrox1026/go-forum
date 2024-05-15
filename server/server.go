package server

import (
	"kevinku/go-forum/config"
	l "kevinku/go-forum/lib/logger"
	"kevinku/go-forum/middleware"
	"os"

	"github.com/gin-gonic/gin"
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

func (server *Server) Run(addr string) {
	server.engine.Run(addr)
}
