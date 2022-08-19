package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/pkg/repositories"
)

type Server struct {
	UserRepository repositories.UserRepository
}

func NewServer() *Server {
	return &Server{}
}

func returnJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(jsonResponse)
}

func returnError(w http.ResponseWriter, error string, code int) {
	returnJSON(w, map[string]string{"error": error}, code)
}
