package repositories

import (
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
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

type DbUserRespository struct {
	db *sqlx.DB
}

func NewDbUserRepository(db *sqlx.DB) *DbUserRespository {
	return &DbUserRespository{db: db}
}

func (repo *DbUserRespository) Get(email string) (*User, error) {
	var count int
	err := repo.db.Get(&count, "SELECT count(id) FROM users WHERE email=$1", email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("User with given email was not found in the db")
	}

	user := User{}
	err = repo.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *DbUserRespository) Create(FirstName string, LastName string, Password string, Email string) (*User, error) {
	result, err := repo.db.Exec("INSERT INTO users (firstname,lastname, passwordhash, email) VALUES (?, ?,?,?)", FirstName, LastName, Password, Email)

	if err != nil {
		log.Println(err)
	}
	user := User{}
	id, _ := result.LastInsertId()
	err = repo.db.Get(&user, "SELECT * FROM users where id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
