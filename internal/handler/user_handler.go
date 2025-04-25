package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/andreparelho/codecon-challenge/internal/repository"
)

func SendUsersFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		path := "/data/" + handler.Filename
		f, err := os.Create(path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content, err := os.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var users []repository.User
		if err := json.Unmarshal(content, &users); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
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
