package main

import (
	_ "kevinku/go-forum/docs"
	serverpkg "kevinku/go-forum/server"
)

func main() {
	var server = serverpkg.NewServer()
	server.Run(":8081")
}
