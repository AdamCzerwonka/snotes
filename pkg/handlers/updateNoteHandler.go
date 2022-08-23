package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandleUpdateNote() http.HandlerFunc {
	type NoteInput struct {
		Title   string
		Content string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteId, err := strconv.Atoi(vars["note_id"])
		if err != nil {
			returnError(w, "Could not parse argument", http.StatusBadRequest)
			return
		}

		loggedInId := r.Context().Value("id").(int)

		var input NoteInput
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			returnError(w, "Could not parse request body", http.StatusInternalServerError)
			return
		}

		note, err := s.NotesService.Get(loggedInId, noteId)
		if err != nil {
			returnError(w, err.Error(), http.StatusNotFound)
			return
		}

		note, err = s.NotesService.Update(noteId, input.Title, input.Content)
		if err != nil {
			returnError(w, "Could not update the note", http.StatusInternalServerError)
			return
		}

		returnJSON(w, note, http.StatusOK)

	}
}
