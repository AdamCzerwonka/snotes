package repositories

import (
	"errors"
	"time"
)

type User struct {
	Id           int
	Firstname    string
	LastName     string
	PasswordHash string
	Email        string
	LastLogin    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type UserRepository interface {
	Get(email string) (*User, error)
	Create(FirstName string, LastName string, Password string, Email string) (*User, error)
}

type InMemoryUserRepository struct {
	users []*User
}

func (repo *InMemoryUserRepository) Get(email string) (*User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("User with given email was not found in the db")
}

func (repo *InMemoryUserRepository) Create(FirstName string, LastName string, Password string, Email string) (*User, error) {
	newIdx := len(repo.users)
	newUser := &User{
		Id:           newIdx,
		Firstname:    FirstName,
		LastName:     LastName,
		PasswordHash: Password,
		Email:        Email,
		LastLogin:    time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    time.Now(),
	}

	repo.users = append(repo.users, newUser)
	return newUser, nil
}
