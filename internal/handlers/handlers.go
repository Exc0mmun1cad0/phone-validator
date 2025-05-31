package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	phonenumber "phone-validator/pkg/phone_number"
	"strings"
)

type Request struct {
	PhoneNumber string `json:"phone_number"`
}

type Response struct {
	Status     bool   `json:"status"`
	Normalized string `json:"normalized,omitempty"`
}

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ShutdownHandler(_ http.ResponseWriter, _ *http.Request) {
	os.Exit(0)
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			"Failed to read request body",
			http.StatusInternalServerError,
		)
		return
	}

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(
			w,
			"Failed to unmarshal request body",
			http.StatusBadRequest,
		)
		return
	}

	phoneNumber := req.PhoneNumber
	if phoneNumber == "" {
		http.Error(
			w,
			"No phone_number to validate",
			http.StatusBadRequest,
		)
		return
	}

	isValid, err := phonenumber.IsValidPhoneNum(phoneNumber)
	if err != nil {
		http.Error(
			w,
			"Failed to validate phone_number",
			http.StatusInternalServerError,
		)
		return
	}

	resp := Response{Status: isValid}
	if isValid {
		resp.Normalized = phonenumber.Normalize(phoneNumber)
	}

	msg, err := json.Marshal(resp)
	if err != nil {
		http.Error(
			w,
			"Failed to marshal response",
			http.StatusInternalServerError,
		)
		return
	}

	if _, err := w.Write(msg); err != nil {
		http.Error(
			w,
			"Failed to write response",
			http.StatusInternalServerError,
		)
	}
}
