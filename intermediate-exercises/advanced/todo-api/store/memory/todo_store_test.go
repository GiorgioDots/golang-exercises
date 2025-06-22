package memory

import (
	"testing"

	"github.com/giorgiodots/todo-go-api/models"
)

func TestAddAndListTodos(t *testing.T) {
	store := NewInMemoryStore()

	text := "Test todo"
	done := false
	todo, err := store.Add(models.CreateTodoRequest{
		Text: &text,
		Done: &done,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	todos, err := store.List()
	if err != nil {
		t.Fatalf("Unexpected error in list method %v", err)
	}
	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	if todos[0].ID != todo.ID {
		t.Errorf("expected ID %d, got %d", todo.ID, todos[0].ID)
	}
}

func TestGetTodoByID(t *testing.T) {
	store := NewInMemoryStore()

	text := "Find me"
	done := true
	todo, _ := store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	found, err := store.GetByID(todo.ID)
	if err != nil {
		t.Fatalf("expected to find todo, got error %v", err)
	}

	if found.Text != text {
		t.Errorf("expected text %q, got %v", text, found.Text)
	}
}

func TestUpdateTodo(t *testing.T) {
	store := NewInMemoryStore()

	text := "Original"
	done := false
	todo, _ := store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	newText := "Updated"
	newDone := true
	err := store.Update(todo.ID, models.UpdateTodoRequest{
		Text: &newText,
		Done: &newDone,
	})
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated, _ := store.GetByID(todo.ID)
	if updated.Text != newText || updated.Done != newDone {
		t.Errorf("update failed: got %+v", updated)
	}
}

func TestDeleteTodo(t *testing.T) {
	store := NewInMemoryStore()

	text := "To be deleted"
	done := false
	todo, _ := store.Add(models.CreateTodoRequest{Text: &text, Done: &done})

	err := store.Delete(todo.ID)
	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	_, err = store.GetByID(todo.ID)
	if err == nil {
		t.Fatalf("expected error when getting deleted todo")
	}
}
