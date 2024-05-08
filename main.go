package main

import (
	"kevinku/go-forum/app/service"
	"kevinku/go-forum/config"
	"kevinku/go-forum/database"
	_ "kevinku/go-forum/docs"
	"kevinku/go-forum/lib"
	serverpkg "kevinku/go-forum/server"
)

func init() {
	config.Init()
	lib.Init()
	database.Init()

	service.Init()
}

func main() {

	var server = serverpkg.NewServer()
	server.Run(":8081")
}
