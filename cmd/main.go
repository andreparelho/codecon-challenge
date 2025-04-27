package main

import (
	"github.com/andreparelho/codecon-challenge/internal/repository"
	"github.com/andreparelho/codecon-challenge/internal/server"
)

func main() {
	db, err := repository.CreateDatabase()
	if err != nil {
		panic(err)
	}

	u := repository.NewUserRepository(db)

	server.RegisterRoutes(u)
	server.StartServer()
}
