package server

import (
	"net/http"

	"github.com/andreparelho/codecon-challenge/internal/handler"
)

func RegisterRoutes() {
	http.HandleFunc("/user", handler.SaveUsers())
	http.HandleFunc("/superusers", handler.GetSuperUsers())
	http.HandleFunc("/top-countries", handler.GetSuperUsersByTopCountries())
	http.HandleFunc("/active-users-per-day", handler.GetActiveUsers())
	http.HandleFunc("/team-insights", handler.GetMembers())
	http.HandleFunc("/evaluation", handler.Evaluation())
}
