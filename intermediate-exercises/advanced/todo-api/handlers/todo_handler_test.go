package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/giorgiodots/todo-go-api/handlers"
	"github.com/giorgiodots/todo-go-api/models"
	"github.com/giorgiodots/todo-go-api/store/memory"
)

func TestCreateAndListTodos(t *testing.T) {
	store := memory.NewInMemoryStore()
	handler := handlers.NewTodoHandler(store)

	// Simulate POST /todos
	body := bytes.NewBuffer([]byte(`{"text":"Test item","done":false}`))
	req := httptest.NewRequest(http.MethodPost, "/todos", body)
	w := httptest.NewRecorder()

	handler.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	// Simulate GET /todos
	req2 := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w2 := httptest.NewRecorder()
	handler.List(w2, req2)

	var todos []models.Todo
	if err := json.NewDecoder(w2.Body).Decode(&todos); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if len(todos) != 1 || todos[0].Text != "Test item" {
		t.Errorf("unexpected todo list: %+v", todos)
	}
}
