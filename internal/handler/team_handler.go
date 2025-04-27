package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andreparelho/codecon-challenge/internal/repository"
	"github.com/sirupsen/logrus"
)

func GetMembers(repo repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		start := time.Now()

		if r.Method != http.MethodGet {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("this method not supported")

			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		teams, err := repo.GetMembers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("error to get teams")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userHandlerResponse := UserHandlerResponse{
			Status: http.StatusOK,
			Body: map[string]interface{}{
				"teams": teams,
			},
		}

		var response []byte
		if response, err = json.Marshal(userHandlerResponse); err != nil {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("error to unmarshal teams")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logrus.WithFields(logrus.Fields{
			"timestamp": time.Since(start),
			"teams":     len(teams),
		}).Info("success to get teams insights")

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
