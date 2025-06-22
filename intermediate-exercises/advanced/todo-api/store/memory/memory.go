package memory

import (
	"fmt"
	"slices"
	"sync"

	"github.com/giorgiodots/todo-go-api/models"
)

type InMemoryTodoStore struct {
	mu     sync.Mutex
	todos  []models.Todo
	nextID int
}

func NewInMemoryStore() *InMemoryTodoStore {
	return &InMemoryTodoStore{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

func (s *InMemoryTodoStore) Add(createData models.CreateTodoRequest) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo := models.NewTodo()
	todo.ID = s.nextID
	s.nextID++
	createData.CopyTo(&todo)
	s.todos = append(s.todos, todo)
	return todo, nil
}

func (s *InMemoryTodoStore) List() ([]models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	copy := append([]models.Todo(nil), s.todos...)
	if copy == nil {
		copy = []models.Todo{}
	}
	return deepCopyTodos(copy), nil
}

func (s *InMemoryTodoStore) GetByID(id int) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.todos {
		if s.todos[i].ID == id {
			return s.todos[i], nil
		}
	}
	return models.Todo{}, fmt.Errorf("Could not find todo with id %d", id)
}

func (s *InMemoryTodoStore) Update(id int, updateData models.UpdateTodoRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.todos {
		if s.todos[i].ID == id {
			updateData.CopyTo(&s.todos[i])
			return nil
		}
	}
	return fmt.Errorf("Could not find todo with id %d", id)
}

func (s *InMemoryTodoStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.todos = slices.DeleteFunc(s.todos, func(t models.Todo) bool {
		return t.ID == id
	})
	return nil
}

func deepCopyTodos(todos []models.Todo) []models.Todo {
	copy := make([]models.Todo, len(todos))
	for i, t := range todos {
		copy[i] = t // value copy (deep enough for simple types)
	}
	return copy
}
