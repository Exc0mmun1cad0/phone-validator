package main

import (
	"encoding/json"
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
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
		json.NewEncoder(w).Encode(
			Response{Status: false},
		)
	} else {
		normalized := phonenumber.NormalizePhoneNum(phoneNum)
		json.NewEncoder(w).Encode(
			Response{
				Status:     true,
				Normalized: normalized,
			},
		)
	}
}

func main() {
	router := chi.NewRouter()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/shutdown", shutdownHandler)
	router.HandleFunc("/validatePhoneNumber", validateHandler)

	http.ListenAndServe(host, router)
}
