package main

import (
	"errors"
	"log"
	"net/http"
	"phone-validator/internal/handlers"

	"github.com/go-chi/chi/v5"
)

const host = ":7777"

func main() {
	router := chi.NewRouter()

	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/ping", handlers.PingHandler)
	router.HandleFunc("/shutdown", handlers.ShutdownHandler)
	router.HandleFunc("/validate", handlers.ValidateHandler)

	// TODO: implement graceful shutdown
	if err := http.ListenAndServe(host, router); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to shutdown server")
	}
}
