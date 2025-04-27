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

	repo := repository.NewUserRepository(db)

	server.RegisterRoutes(repo)
	server.StartServer()
}
