package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/andreparelho/codecon-challenge/internal/repository"
	"github.com/sirupsen/logrus"
)

func SendUsersFile(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		start := time.Now()

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

		userHandlerResponse := UserHandlerResponse{
			Status: http.StatusOK,
			Body: map[string]interface{}{
				"timestamp":  time.Since(start).Milliseconds(),
				"message":    "Arquivo recebido com sucesso",
				"user_count": len(users),
			},
		}

		var response []byte
		if response, err = json.Marshal(userHandlerResponse); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to marshal users to response bytes")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func GetSuperUsers(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		users, err := u.GetSuperusers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to get superusers")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userHandlerResponse := UserHandlerResponse{
			Status: http.StatusOK,
			Body: map[string]interface{}{
				"timestamp": time.Since(start).Milliseconds(),
				"data":      users,
			},
		}

		var response []byte
		if response, err = json.Marshal(userHandlerResponse); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to marshal users to response bytes")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func GetTopCountries(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		countries, err := u.GetTopCountries()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to get top countries")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userHandlerResponse := UserHandlerResponse{
			Status: http.StatusOK,
			Body: map[string]interface{}{
				"timestamp": time.Since(start).Milliseconds(),
				"countries": countries,
			},
		}

		var response []byte
		if response, err = json.Marshal(userHandlerResponse); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to marshal users to response bytes")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
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
