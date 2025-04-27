package server

import (
	"net/http"

	"github.com/andreparelho/codecon-challenge/internal/handler"
	"github.com/andreparelho/codecon-challenge/internal/repository"
)

func RegisterRoutes(repo repository.UserRepository) {
	http.HandleFunc("/user", handler.SendUsersFile(repo))
	http.HandleFunc("/superusers", handler.GetSuperUsers(repo))
	http.HandleFunc("/top-countries", handler.GetTopCountries(repo))
	http.HandleFunc("/active-users-per-day", handler.GetActiveUsers(repo))
	http.HandleFunc("/team-insights", handler.GetMembers(repo))
	http.HandleFunc("/evaluation", handler.Evaluation())
}
