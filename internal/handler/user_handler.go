package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andreparelho/codecon-challenge/internal/repository"
	"github.com/sirupsen/logrus"
)

func SendUsersFile(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {

			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("this method not supported")

			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseMultipartForm(128 << 20); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to set file size")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to get file")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to read file")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var users []*repository.User
		if err := json.Unmarshal([]byte(content), &users); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to unmarshal file to users")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := u.SaveUsers(users); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to save users")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetSuperUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func GetSuperUsersByTopCountries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func GetActiveUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
