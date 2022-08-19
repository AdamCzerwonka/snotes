package handlers

import (
	"net/http"
	"net/http/httptest"
	"notes-app/pkg/repositories"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerWithCorrectInput(t *testing.T) {
	payload := strings.NewReader(`{
	"firstName":"Adam",
	"lastName":"Czerwonka",
	"Password": "Test",
	"Password2": "Test",
	"Email":"test1@test.com"
	}`)
	req, err := http.NewRequest("POST", "/api/user", payload)
	if err != nil {
		t.Fatal(err)
	}

	server := NewServer()
	server.UserRepository = &repositories.InMemoryUserRepository{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleCreateUser())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestHandlerWithNotCompleteInput(t *testing.T) {
	payload := strings.NewReader(`{
	"lastName":"Czerwonka",
	"Password": "Test",
	"Password2": "Test",
	"Email":"test1@test.com"
	}`)
	req, err := http.NewRequest("POST", "/api/user", payload)
	if err != nil {
		t.Fatal(err)
	}

	server := NewServer()
	server.UserRepository = &repositories.InMemoryUserRepository{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleCreateUser())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandlerWithNotMatchingPasswords(t *testing.T) {
	payload := strings.NewReader(`{
		"firstName":"Adam",
		"lastName":"Czerwonka",
		"Password": "Test123",
		"Password2": "Test",
		"Email":"test1@test.com"
	}`)
	req, err := http.NewRequest("POST", "/api/user", payload)
	if err != nil {
		t.Fatal(err)
	}

	server := NewServer()
	server.UserRepository = &repositories.InMemoryUserRepository{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleCreateUser())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandlerWithWrongEmail(t *testing.T) {
	payload := strings.NewReader(`{
	"lastName":"Czerwonka",
	"Password": "Test",
	"Password2": "Test",
	"Email":"test"
	}`)
	req, err := http.NewRequest("POST", "/api/user", payload)
	if err != nil {
		t.Fatal(err)
	}

	server := NewServer()
	server.UserRepository = &repositories.InMemoryUserRepository{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleCreateUser())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandlerWhileUserAlreadyExists(t *testing.T) {
	payload := strings.NewReader(`{
	"firstName":"Joe",
	"lastName":"Czerwonka",
	"Password": "Test",
	"Password2": "Test",
	"Email":"test1@test.com"
	}`)
	req, err := http.NewRequest("POST", "/api/user", payload)
	if err != nil {
		t.Fatal(err)
	}

	server := NewServer()
	server.UserRepository = &repositories.InMemoryUserRepository{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleCreateUser())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
