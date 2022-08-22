package routes

import (
	"log"
	"notes-app/pkg/handlers"
	"notes-app/pkg/middleware"
	"notes-app/pkg/repositories"
	"notes-app/pkg/services"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func RegisterRoutes(r *mux.Router) {
	server := handlers.NewServer()
	db, err := sqlx.Connect("postgres", "user=postgres dbname=notes-app sslmode=disable password=example")
	if err != nil {
		log.Fatalln(err)
	}

	server.Db = db
	server.UserRepository = repositories.NewDbUserRepository(server.Db)
	server.AuthService = services.AuthService{}
	server.NotesService = services.NewInMemoryNotesService()

	r.HandleFunc("/", server.HandleHome())
	r.HandleFunc("/api/user", server.HandleCreateUser()).Methods("POST")
	r.HandleFunc("/api/notes", middleware.IsLoggedIn(server.HandleCreateNote())).Methods("POST")
	r.HandleFunc("/api/login", server.HandleLogin()).Methods("POST")
}
