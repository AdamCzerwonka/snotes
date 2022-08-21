package middleware

import (
	"context"
	"log"
	"net/http"
	"notes-app/pkg/services"
)

func IsLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := services.AuthService{}
		id, err := auth.ValidateTokenFromRequest(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "id", id)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
