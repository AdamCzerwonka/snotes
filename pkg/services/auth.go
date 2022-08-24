package services

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const EXPIERATION_PERIOD = 1
const SECRET_KEY string = "very_secret_password_123!"

type AuthService struct{}

type Claims struct {
	FirstName string `json:"firstName"`
	Id        int    `json:"id"`
	jwt.StandardClaims
}

func (auth *AuthService) GenerateToken(id int, firstName string) (string, error) {
	expirationTime := time.Now().Add(EXPIERATION_PERIOD * time.Hour)

	claims := &Claims{
		Id:        id,
		FirstName: firstName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (auth *AuthService) ValidateTokenFromRequest(r *http.Request) (int, error) {
	token := r.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer", "", 1)
	token = strings.TrimSpace(token)
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return -1, err
	}

	if !tkn.Valid {
		return -1, errors.New("Token not valid")
	}

	return claims.Id, nil
}
