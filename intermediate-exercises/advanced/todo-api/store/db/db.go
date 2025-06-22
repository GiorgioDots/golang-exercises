package db

import (
	"database/sql"

	"github.com/giorgiodots/todo-go-api/models"
)

type DBTodoStore struct {
	db *sql.DB // or GORM, etc.
}

func (s *DBTodoStore) Add(todo models.Todo) (models.Todo, error) {
	// insert into db, return ID
	panic("Not implemented")
}

func (s *DBTodoStore) List() ([]models.Todo, error) {
	// select * from todos
	panic("Not implemented")
}
