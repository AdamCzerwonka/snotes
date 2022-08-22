package services

import (
	"database/sql"
	"time"
)

type Note struct {
	Id        int           `db:"id"`
	Title     string        `db:"title"`
	Content   string        `db:"content"`
	OnwerId   int           `db:"user_id"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
	DeletedAt *sql.NullTime `db:"deleted_at"`
}

type NotesService interface {
	Create(title string, content string, ownerId int) (*Note, error)
}

type InMemoryNotesService struct {
	notes []*Note
}

func NewInMemoryNotesService() *InMemoryNotesService {
	return &InMemoryNotesService{}
}

func (s *InMemoryNotesService) Create(title string, content string, ownerId int) (*Note, error) {
	idx := len(s.notes)
	newNote := &Note{
		Id:        idx,
		Title:     title,
		Content:   content,
		OnwerId:   ownerId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	s.notes = append(s.notes, newNote)
	return newNote, nil
}
