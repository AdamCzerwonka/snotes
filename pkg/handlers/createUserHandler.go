package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func (s *Server) HandleCreateUser() http.HandlerFunc {
	type UserInput struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		Email     string `validate:"required,email"`
		Password  string `validate:"required"`
		Password2 string `validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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
					fmt.Printf("Field %s is required", err.Field())
				}
				fmt.Println(err.Tag())
				fmt.Println(err.Type())
			}
			json, err := json.Marshal(errors)
			if err != nil {
				log.Println("Failed to parse errors to json")
			}
			fmt.Fprintf(w, string(json))
			return
		}

		_, err = s.UserRepository.Get(input.Email)

		if err == nil {
			http.Error(w, "User with given Email addr already exists!", http.StatusBadRequest)
		}
		user, err := s.UserRepository.Create(input.FirstName, input.LastName, input.Password, input.Email)
		json, err := json.Marshal(user)
		fmt.Fprintf(w, string(json))
	}
}
