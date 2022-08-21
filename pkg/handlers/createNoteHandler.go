package handlers

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleCreateNote() http.HandlerFunc {
	type NoteInput struct {
		Title   string
		Content string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		input := NoteInput{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			returnError(w, "Could not parse input", http.StatusBadRequest)
			return
		}

	}
}
