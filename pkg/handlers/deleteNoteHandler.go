package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandleDeleteNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedInId := r.Context().Value("id").(int)
		vars := mux.Vars(r)
		noteId, err := strconv.Atoi(vars["note_id"])
		if err != nil {
			returnError(w, "Could not parse argument", http.StatusNotFound)
			return
		}

		note, err := s.NotesService.Get(loggedInId, noteId)
		if err != nil {
			returnError(w, err.Error(), http.StatusNotFound)
			return
		}

		err = s.NotesService.Delete(note.Id)
		if err != nil {
			returnError(w, "There was an error during the operation", http.StatusBadRequest)
			return
		}

		returnJSON(w, nil, http.StatusNoContent)
	}
}
