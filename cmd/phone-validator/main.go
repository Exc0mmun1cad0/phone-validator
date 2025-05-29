// 
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	phonenumber "phone-validator/pkg/phone_number"

	"github.com/go-chi/chi/v5"
)

const host = "127.0.0.1:7777"

type Response struct {
	Status     bool   `json:"status"`
	Normalized string `json:"normalized,omitempty"`
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func shutdownHandler(_ http.ResponseWriter, _ *http.Request) {
	os.Exit(0)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	decodedPhoneNum := params.Get("phone_number")
	phoneNum, err := url.QueryUnescape(decodedPhoneNum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if phoneNum == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, err := phonenumber.IsValidPhoneNum(phoneNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if !isValid {
		if err := json.NewEncoder(w).Encode(
			Response{Status: false},
		); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		normalized := phonenumber.NormalizePhoneNum(phoneNum)
		if err := json.NewEncoder(w).Encode(
			Response{
				Status:     true,
				Normalized: normalized,
			},
		); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	router := chi.NewRouter()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/shutdown", shutdownHandler)
	router.HandleFunc("/validatePhoneNumber", validateHandler)

	// TODO: implement graceful shutdown
	if err := http.ListenAndServe(host, router); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to shutdown server")
	}
}
