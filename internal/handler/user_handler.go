package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

		file, header, err := r.FormFile("file")
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

		logrus.WithFields(logrus.Fields{
			"timestamp":   time.Since(start),
			"users_total": len(users),
			"file_size":   header.Size,
			"file_name":   header.Filename,
		}).Info("success to save user on database")

		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

func GetSuperUsers(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		start := time.Now()

		if r.Method != http.MethodGet {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("this method not supported")

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

		logrus.WithFields(logrus.Fields{
			"timestamp":        time.Since(start),
			"superusers_total": len(users),
		}).Info("success to get superusers")

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func GetTopCountries(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		start := time.Now()

		if r.Method != http.MethodGet {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("this method not supported")

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		total, err := strconv.Atoi(r.URL.Query().Get("total"))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to get path param")

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		countries, err := u.GetTopCountries(total)
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

		logrus.WithFields(logrus.Fields{
			"timestamp":       time.Since(start),
			"total_countries": len(countries),
		}).Info("success to get top countries")

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func GetActiveUsers(u repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		start := time.Now()

		if r.Method != http.MethodGet {
			logrus.WithFields(logrus.Fields{
				"method": r.Method,
				"url":    r.URL,
			}).Error("this method not supported")

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		users, err := u.GetActiveUsers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("error to get active users")

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userHandlerResponse := UserHandlerResponse{
			Status: http.StatusOK,
			Body: map[string]interface{}{
				"timestamp": time.Since(start).Milliseconds(),
				"logins":    users,
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

		logrus.WithFields(logrus.Fields{
			"timestamp":    time.Since(start),
			"active_users": len(users),
		}).Info("success to get active users")

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
