package services

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
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
	Update(noteId int, title string, content string) (*Note, error)
}

type DbNotesService struct {
	db *sqlx.DB
}

func NewDbNotesService(db *sqlx.DB) *DbNotesService {
	return &DbNotesService{db: db}
}

func (s *DbNotesService) Create(title string, content string, ownerId int) (*Note, error) {
	sql := "INSERT INTO notes(title, content, user_id) VALUES ($1,$2,$3) RETURNING id"

	var id int
	err := s.db.Get(&id, sql, title, content, ownerId)
	if err != nil {
		return nil, err
	}

	sql = "SELECT id,title,content,user_id,created_at,updated_at,deleted_at FROM notes WHERE id=$1"

	note := &Note{}
	err = s.db.QueryRowx(sql, id).StructScan(note)
	if err != nil {
		return nil, err
	}

	return note, nil

}

func (s *DbNotesService) GetAll(userId int) []*Note {
	sql := "SELECT id,title,content,user_id,created_at,updated_at,deleted_at FROM notes WHERE user_id=$1 AND deleted_at IS NULL"

	notes := []*Note{}
	err := s.db.Select(&notes, sql, userId)
	if err != nil {
		log.Println(err)
	}

	return notes
}

func (s *DbNotesService) Get(userId int, noteId int) (*Note, error) {
	sql := "SELECT id,title,content,user_id,created_at,updated_at,deleted_at FROM notes WHERE user_id=$1 AND id=$2 AND deleted_at IS NULL"

	note := &Note{}
	err := s.db.QueryRowx(sql, userId, noteId).StructScan(note)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return note, nil
}

func (s *DbNotesService) Delete(noteId int) error {
	sql := "UPDATE notes SET deleted_at=$1 WHERE id=$2"
	_, err := s.db.Exec(sql, time.Now(), noteId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *DbNotesService) Update(noteId int, title string, content string) (*Note, error) {
	sql := "UPDATE notes SET title=$1, content=$2, updated_at=$3 WHERE id=$4"
	_, err := s.db.Exec(sql, title, content, time.Now(), noteId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sql = "SELECT id,title,content,user_id,created_at,updated_at,deleted_at FROM notes WHERE id=$1"

	note := &Note{}
	err = s.db.QueryRowx(sql, noteId).StructScan(note)
	if err != nil {
		return nil, err
	}

	return note, nil
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

func (s *InMemoryNotesService) Update(noteId int, title string, content string) (*Note, error) {
	var updatedNote *Note
	for _, note := range s.notes {
		if note.Id == noteId {
			note.Title = title
			note.Content = content
			note.UpdatedAt = time.Now()
			updatedNote = note
		}
	}

	return updatedNote, nil
}
