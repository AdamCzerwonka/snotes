package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) HandleCreateNote() http.HandlerFunc {
	type NoteInput struct {
		Title   string
		Content string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		loggedInUserId := r.Context().Value("id").(int)
		input := NoteInput{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			returnError(w, "Could not parse input", http.StatusBadRequest)
			return
		}

		user, err := s.UserRepository.GetById(loggedInUserId)
		if err != nil {
			log.Println(err)
		}

		returnJSON(w, user, http.StatusOK)

	}
}
