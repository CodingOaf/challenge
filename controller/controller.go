package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// ASANA_API_URL is the base URL for the Asana API
	ASANA_API_URL = "https://app.asana.com/api/1.0"
)

func NewRouter(args string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/_health", HealthCheck).Methods("GET")
	r.HandleFunc("/users", usersHandler)
	r.HandleFunc("/projects", projectsHandler)
	//r.HandleFunc("/users/me", meHandler)
	//r.HandleFunc("/users/{id}", userHandler).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	return r
}
