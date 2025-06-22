package handlers

import (
	"net/http"

	"github.com/giorgiodots/todo-go-api/models"
	"github.com/giorgiodots/todo-go-api/store"
)

type TodoHandler struct {
	store store.TodoStore
}

func NewTodoHandler(s store.TodoStore) *TodoHandler {
	return &TodoHandler{s}
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	todo, err := ExtractJSON[models.CreateTodoRequest](r)
	if err != nil {
		BadRequest(w, err)
		return
	}
	h.store.Add(todo)
	RespondMessage(w, http.StatusCreated, "Todo created")
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.List()
	if err != nil {
		BadRequest(w, err)
		return
	}
	RespondJSON(w, http.StatusOK, list)
}

func (h *TodoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDParam(r, "id")
	if err != nil {
		BadRequest(w, err)
		return
	}
	todo, err := h.store.GetByID(id)
	if err != nil {
		BadRequest(w, err)
		return
	}
	RespondJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDParam(r, "id")
	if err != nil {
		BadRequest(w, err)
		return
	}
	todo, err := ExtractJSON[models.UpdateTodoRequest](r)
	if err != nil {
		BadRequest(w, err)
		return
	}
	err = h.store.Update(id, todo)
	if err != nil {
		BadRequest(w, err)
		return
	}
	RespondMessage(w, http.StatusCreated, "Todo updated")
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDParam(r, "id")
	if err != nil {
		BadRequest(w, err)
		return
	}
	err = h.store.Delete(id)
	if err != nil {
		BadRequest(w, err)
		return
	}
	RespondMessage(w, http.StatusCreated, "Todo Deleted")
}
