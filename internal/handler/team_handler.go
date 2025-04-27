package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andreparelho/codecon-challenge/internal/repository"
)

func GetMembers(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		teams, err := u.GetMembers()
		if err != nil {

		}

		userHandlerResponse := UserHandlerResponse{
			Status: http.StatusOK,
			Body: map[string]interface{}{
				"teams": teams,
			},
		}

		var response []byte
		if response, err = json.Marshal(userHandlerResponse); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
