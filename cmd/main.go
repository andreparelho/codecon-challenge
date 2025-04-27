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

	w := db.Txn(true)
	r := db.Txn(false)

	u := repository.NewUserRepository(w, r)

	server.RegisterRoutes(u)
	server.StartServer()
}
