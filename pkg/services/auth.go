package services

import (
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
