package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleCreateUser() http.HandlerFunc {
	type UserInput struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		Email     string `validate:"required,email"`
		Password  string `validate:"required"`
		Password2 string `validate:"required"`
	}

	type UserResult struct {
		Id        int
		FirstName string
		LastName  string
		Email     string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var input UserInput

		validate := validator.New()

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
		}

		var errors []string

		err = validate.Struct(&input)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				if err.Tag() == "required" {
					errors = append(errors, fmt.Sprintf("Field %s is required", err.Field()))
				}
			}

			returnJSON(w, errors, http.StatusBadRequest)
			return
		}

		_, err = s.UserRepository.Get(input.Email)

		if err == nil {
			returnError(w, "User with given Email address already exists.", http.StatusBadRequest)
			return
		}

		if input.Password != input.Password2 {
			returnError(w, "Passwords does not match", http.StatusBadRequest)
			return
		}

		hashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
		passwordHash := string(hashBytes)

		user, err := s.UserRepository.Create(input.FirstName, input.LastName, passwordHash, input.Email)
		result := &UserResult{
			Id:        user.Id,
			FirstName: user.Firstname,
			LastName:  user.LastName,
			Email:     user.Email,
		}
		returnJSON(w, result, http.StatusCreated)
	}
}
