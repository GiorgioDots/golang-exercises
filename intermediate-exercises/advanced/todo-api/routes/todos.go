package routes

import (
	"github.com/giorgiodots/todo-go-api/handlers"
	"github.com/giorgiodots/todo-go-api/store"
	"github.com/go-chi/chi/v5"
)

type todosResource struct {
	store store.TodoStore
}

func (rs todosResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware
	h := handlers.NewTodoHandler(rs.store)

	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/{id}", h.GetByID)
	r.Patch("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)

	return r
}

func NewTodosResource(store store.TodoStore) *todosResource {
	return &todosResource{store}
}
