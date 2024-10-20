package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var opCodes = []string{"982", "986", "912", "934"}
var formats = []string{
	`8%s\d{7}`,
	`8 (%s) \d{3}-\d{4}`,
	`8 %s \d{3} \d{2}[ ]?\d{2}`,
	`\+7%s\d{7}`,
	`\+7 (%s) \d{3}-\d{4}`,
	`\+7 %s \d{3} \d{2}[ ]?\d{2}`,
}

const host = "127.0.0.1:7777"

func isValidPhoneNum(phoneNum string) (bool, error) {
	for _, opCode := range opCodes {
		for _, format := range formats {
			isValid, err := regexp.MatchString(
				fmt.Sprintf(format, opCode),
				phoneNum,
			)
			if err != nil {
				return false, fmt.Errorf("cannot validate phone number: %w", err)
			}
			if isValid {
				return true, nil
			}
		}
	}

	return false, nil
}

func normalizePhoneNum(phoneNum string) string {
	for _, sym := range []string{" ", "-", "(", ")"} {
		phoneNum = strings.ReplaceAll(phoneNum, sym, "")
	}
	if strings.HasPrefix(phoneNum, "8") {
		phoneNum = strings.Replace(phoneNum, "8", "+7", 1)
	}
	phoneNum = phoneNum[:2] + "-" + phoneNum[2:5] + "-" + phoneNum[5:8] + "-" + phoneNum[8:]

	return phoneNum
}

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

	isValid, err := isValidPhoneNum(phoneNum)
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
		normalized := normalizePhoneNum(phoneNum)
		json.NewEncoder(w).Encode(
			Response{
				Status:     true,
				Normalized: normalized,
			},
		)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/shutdown", shutdownHandler)
	mux.HandleFunc("/validatePhoneNumber", validateHandler)

	http.ListenAndServe(host, mux)
}
