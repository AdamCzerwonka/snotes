package services

import (
	"database/sql"
	"errors"
	"time"
)

type Note struct {
	Id        int           `db:"id"`
	Title     string        `db:"title"`
	Content   string        `db:"content"`
	OwnerId   int           `db:"user_id"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
	DeletedAt *sql.NullTime `db:"deleted_at"`
}

type NotesService interface {
	Create(title string, content string, ownerId int) (*Note, error)
	GetAll(userId int) []*Note
	Get(userId int, noteId int) (*Note, error)
	Delete(noteId int) error
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
		OwnerId:   ownerId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	s.notes = append(s.notes, newNote)
	return newNote, nil
}

func (s *InMemoryNotesService) GetAll(userId int) []*Note {
	var result []*Note

	for _, note := range s.notes {
		if note.OwnerId == userId {
			result = append(result, note)
		}
	}
	return result
}

func (s *InMemoryNotesService) Get(userId int, noteId int) (*Note, error) {
	for _, note := range s.notes {
		if note.OwnerId == userId && note.Id == noteId {
			return note, nil
		}
	}

	return nil, errors.New("Note not find")
}

func (s *InMemoryNotesService) Delete(noteId int) error {
	var noteToDelete int
	for idx, note := range s.notes {
		if note.Id == noteId {
			noteToDelete = idx
			break
		}
	}

	s.notes = append(s.notes[:noteToDelete], s.notes[noteToDelete+1:]...)
	return nil

}
