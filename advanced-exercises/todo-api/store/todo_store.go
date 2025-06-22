package store

import "github.com/giorgiodots/todo-go-api/models"

type TodoStore interface {
	Add(todo models.CreateTodoRequest) (models.Todo, error)
	List() ([]models.Todo, error)
	GetByID(id int) (models.Todo, error)
	Update(id int, todo models.UpdateTodoRequest) error
	Delete(id int) error
}
