package handlers

import "notes-app/pkg/repositories"

type Server struct {
	UserRepository repositories.UserRepository
}

func NewServer() *Server {
	return &Server{}
}
