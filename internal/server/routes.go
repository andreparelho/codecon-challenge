package server

import (
	"net/http"

	"github.com/andreparelho/codecon-challenge/internal/handler"
	"github.com/andreparelho/codecon-challenge/internal/repository"
)

func RegisterRoutes(u repository.UserRepository) {
	http.HandleFunc("/user", handler.SendUsersFile(u))
	http.HandleFunc("/superusers", handler.GetSuperUsers())
	http.HandleFunc("/top-countries", handler.GetSuperUsersByTopCountries())
	http.HandleFunc("/active-users-per-day", handler.GetActiveUsers())
	http.HandleFunc("/team-insights", handler.GetMembers())
	http.HandleFunc("/evaluation", handler.Evaluation())
}
