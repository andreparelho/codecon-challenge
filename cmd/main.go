package main

import "github.com/andreparelho/codecon-challenge/internal/server"

func main() {
	server.RegisterRoutes()
	server.StartServer()
}
