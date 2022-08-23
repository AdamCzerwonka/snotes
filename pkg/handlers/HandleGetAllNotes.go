package handlers

import (
	"log"
	"net/http"
)

func (s *Server) HandleGetAllNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedInUserId := r.Context().Value("id").(int)

		user, err := s.UserRepository.GetById(loggedInUserId)
		if err != nil {
			log.Println(err)
		}

		notes := s.NotesService.GetAll(user.Id)

		returnJSON(w, notes, http.StatusOK)
	}
}
