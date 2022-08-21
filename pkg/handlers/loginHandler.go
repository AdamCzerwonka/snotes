package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleLogin() http.HandlerFunc {
	type LoginInput struct {
		Email    string
		Password string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		input := LoginInput{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			returnError(w, "Could not parse the input", http.StatusBadRequest)
			return
		}

		user, err := s.UserRepository.Get(input.Email)
		if err != nil {
			returnError(w, "Wrong email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
		if err != nil {
			returnError(w, "Wrong email or password", http.StatusUnauthorized)
			return
		}

		token, err := s.AuthService.GenerateToken(user.Id, user.Firstname)
		if err != nil {
			returnError(w, "Could not perform auth action", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if err = s.UserRepository.UpdateLastLogin(user.Id); err != nil {
			log.Println(err)
			return
		}

		returnJSON(w, map[string]string{"token": token}, http.StatusOK)
	}
}
