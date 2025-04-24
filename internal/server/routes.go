package server

import (
	"net/http"

	"github.com/andreparelho/codecon-challenge/internal/handler"
)

func RegisterRoutes() {
	mux := http.NewServeMux()

	mux.Handle("/user", handler.SaveUsers())
	mux.Handle("/superusers", handler.GetSuperUsers())
	mux.Handle("/top-countries", handler.GetSuperUsersByTopCountries())
	mux.Handle("/active-users-per-day", handler.GetActiveUsers())
	mux.Handle("/team-insights", handler.GetMembers())
	mux.Handle("/evaluation", handler.Evaluation())
}
