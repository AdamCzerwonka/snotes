package routes

import (
	"notes-app/pkg/handlers"
	"notes-app/pkg/repositories"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	server := handlers.NewServer()
	server.UserRepository = &repositories.InMemoryUserRepository{}
	r.HandleFunc("/", server.HandleHome())
	r.HandleFunc("/api/user", server.HandleCreateUser()).Methods("POST")
}
