package repositories

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id           int           `db:"id"`
	Firstname    string        `db:"firstname"`
	LastName     string        `db:"lastname"`
	PasswordHash string        `db:"passwordhash"`
	Email        string        `db:"email"`
	LastLogin    *sql.NullTime `db:"lastlogin"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	DeletedAt    *sql.NullTime `db:"deletedat"`
}

type UserRepository interface {
	Get(email string) (*User, error)
	Create(FirstName string, LastName string, Password string, Email string) (*User, error)
	GetById(id int) (*User, error)
	UpdateLastLogin(id int) error
}

type InMemoryUserRepository struct {
	users []*User
}

func (repo *InMemoryUserRepository) UpdateLastLogin(id int) error {
	return nil
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
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	repo.users = append(repo.users, newUser)
	return newUser, nil
}

func (repo *InMemoryUserRepository) GetById(id int) (*User, error) {
	for _, user := range repo.users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New("User with given id was not found")
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
	var id int
	err := repo.db.Get(&id, "INSERT INTO users (firstname,lastname, passwordhash, email) VALUES ($1, $2,$3,$4) RETURNING id", FirstName, LastName, Password, Email)

	if err != nil {
		log.Println(err)
	}

	log.Println(id)
	user := User{}
	err = repo.db.QueryRowx("SELECT id,email,firstname,lastname,passwordhash,lastlogin,created_at,updated_at,deletedat FROM users WHERE id=$1", id).StructScan(&user)
	log.Println(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *DbUserRespository) GetById(id int) (*User, error) {
	user := User{}
	err := repo.db.QueryRowx("SELECT id,email,firstname,lastname,passwordhash,lastlogin,created_at,updated_at,deletedat FROM users WHERE id=$1 AND deletedat IS NULL", id).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *DbUserRespository) UpdateLastLogin(id int) error {
	_, err := repo.db.Exec("UPDATE users SET lastlogin=$1 WHERE id=$2", time.Now(), id)
	return err
}
