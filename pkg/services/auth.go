package services

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const EXPIERATION_PERIOD = 1

type AuthService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) AuthService {
	return AuthService{jwtSecret: jwtSecret}
}

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
	tokenString, err := token.SignedString([]byte(auth.jwtSecret))
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
		return []byte(auth.jwtSecret), nil
	})

	if err != nil {
		return -1, err
	}

	if !tkn.Valid {
		return -1, errors.New("Token not valid")
	}

	return claims.Id, nil
}
