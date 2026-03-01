package storage

import (
	"fmt"

	"github.com/amuif/Go-Todo/internal/todo"
)

type MemoryStorage struct {
	todos map[string]todo.Todo
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[string]todo.Todo),
	}
}

func (m *MemoryStorage) Create(t *todo.Todo) error {
	if t == nil {
		return fmt.Errorf("todo cannot be nil")
	}

	m.todos[t.ID] = *t
	return nil
}

func (m *MemoryStorage) GetAll() ([]todo.Todo, error) {
	todos := make([]todo.Todo, 0, len(m.todos))
	for _, t := range m.todos {
		todos = append(todos, t)
	}
	return todos, nil
}

func (m *MemoryStorage) GetByID(id string) (*todo.Todo, error) {
	t, exists := m.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo with ID %s not found", id)
	}
	return &t, nil
}

func (m *MemoryStorage) Update(t *todo.Todo) error {
	if t == nil {
		return fmt.Errorf("todo cannot be nil")
	}

	_, exists := m.todos[t.ID]
	if !exists {
		return fmt.Errorf("todo with ID %s not found", t.ID)
	}

	m.todos[t.ID] = *t
	return nil
}

func (m *MemoryStorage) Delete(id string) error {
	_, exists := m.todos[id]
	if !exists {
		return fmt.Errorf("todo with ID %s not found", id)
	}

	delete(m.todos, id)
	return nil
}
