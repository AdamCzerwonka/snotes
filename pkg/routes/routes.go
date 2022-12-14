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
	server.Config = services.LoadConfig("config.json")

	db, err := sqlx.Connect("postgres", "user=postgres dbname=notes-app sslmode=disable password=example")
	if err != nil {
		log.Fatalln(err)
	}

	server.Db = db
	server.UserRepository = repositories.NewDbUserRepository(server.Db)
	server.AuthService = services.NewAuthService(server.Config.Get("JwtSecretKey"))
	server.NotesService = services.NewDbNotesService(server.Db)

	r.HandleFunc("/", server.HandleHome())
	r.HandleFunc("/api/user", server.HandleCreateUser()).Methods("POST")
	r.HandleFunc("/api/notes", middleware.IsLoggedIn(server.HandleCreateNote())).Methods("POST")
	r.HandleFunc("/api/notes", middleware.IsLoggedIn(server.HandleGetAllNotes())).Methods("GET")
	r.HandleFunc("/api/notes/{note_id}", middleware.IsLoggedIn(server.HandlerGetNote())).Methods("GET")
	r.HandleFunc("/api/notes/{note_id}", middleware.IsLoggedIn(server.HandleDeleteNote())).Methods("DELETE")
	r.HandleFunc("/api/notes/{note_id}", middleware.IsLoggedIn(server.HandleUpdateNote())).Methods("PUT")
	r.HandleFunc("/api/login", server.HandleLogin()).Methods("POST")
}
