package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandlerGetNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedInId := r.Context().Value("id").(int)
		vars := mux.Vars(r)
		noteId, err := strconv.Atoi(vars["note_id"])
		if err != nil {
			returnError(w, "Wrong parameter passed", http.StatusNotFound)
		}

		note, err := s.NotesService.Get(loggedInId, noteId)
		if err != nil {
			returnJSON(w, nil, http.StatusNotFound)
			return
		}

		returnJSON(w, note, http.StatusOK)
	}
}
