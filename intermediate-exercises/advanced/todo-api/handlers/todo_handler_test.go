package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/giorgiodots/todo-go-api/models"
	"github.com/giorgiodots/todo-go-api/store/memory"
	"github.com/go-chi/chi/v5"
)

func TestCreateTodoHandler(t *testing.T) {
	store := memory.NewInMemoryStore()
	handler := NewTodoHandler(store)

	text := "Test item"
	done := false
	body, _ := json.Marshal(models.CreateTodoRequest{
		Text: &text,
		Done: &done,
	})

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Create(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201, got %d", res.StatusCode)
	}
}

func TestListTodosHandler(t *testing.T) {
	store := memory.NewInMemoryStore()
	handler := NewTodoHandler(store)

	// Add a todo
	text := "List me"
	done := true
	_, _ = store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()

	handler.List(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	var todos []models.Todo
	if err := json.NewDecoder(res.Body).Decode(&todos); err != nil {
		t.Fatalf("decoding failed: %v", err)
	}

	if len(todos) != 1 || todos[0].Text != "List me" {
		t.Errorf("unexpected todos: %+v", todos)
	}
}

func TestGetTodoByIDHandler(t *testing.T) {
	store := memory.NewInMemoryStore()
	handler := NewTodoHandler(store)

	text := "Find me"
	done := false
	todo, _ := store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	req := httptest.NewRequest(http.MethodGet, "/todos/"+strconv.Itoa(todo.ID), nil)
	w := httptest.NewRecorder()

	router := chi.NewRouter()
	router.Get("/todos/{id}", handler.GetByID)
	router.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	var got models.Todo
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("decoding failed: %v", err)
	}

	if got.ID != todo.ID {
		t.Errorf("expected ID %v, got %v", todo.ID, got.ID)
	}
}

func TestUpdateTodoHandler(t *testing.T) {
	store := memory.NewInMemoryStore()
	handler := NewTodoHandler(store)

	text := "Initial"
	done := false
	todo, _ := store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	newText := "Updated"
	newDone := true
	body, _ := json.Marshal(models.UpdateTodoRequest{Text: &newText, Done: &newDone})

	req := httptest.NewRequest(http.MethodPatch, "/todos/"+strconv.Itoa(todo.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := chi.NewRouter()
	router.Patch("/todos/{id}", handler.Update)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	updated, _ := store.GetByID(todo.ID)
	if updated.Text != newText || updated.Done != newDone {
		t.Errorf("todo not updated: %+v", updated)
	}
}

func TestDeleteTodoHandler(t *testing.T) {
	store := memory.NewInMemoryStore()
	handler := NewTodoHandler(store)

	text := "Delete me"
	done := true
	todo, _ := store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	req := httptest.NewRequest(http.MethodDelete, "/todos/"+strconv.Itoa(todo.ID), nil)
	w := httptest.NewRecorder()

	router := chi.NewRouter()
	router.Delete("/todos/{id}", handler.Delete)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", w.Code)
	}

	_, err := store.GetByID(todo.ID)
	if err == nil {
		t.Error("expected error when fetching deleted todo")
	}
}
