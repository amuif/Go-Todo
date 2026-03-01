package todo

import "time"

type Service struct{ repo Repository }

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTodo(title string, description *string) (*Todo, error) {
	todo := NewToDo(title)
	if description != nil {
		todo.Description = description
	}
	err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, err
}

func (s *Service) GetAllTodos() ([]Todo, error) {
	return s.repo.GetAll()
}

func (s *Service) GetTodo(id string) (*Todo, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateTodo(id string, title *string, description *string, status *Status, dueDate *time.Time) (*Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if title != nil {
		todo.Title = *title
	}
	if description != nil {
		todo.Description = description
	}
	if status != nil {
		todo.Status = *status
		if *status == StatusCompleted && todo.CompletedAt == nil {
			now := time.Now()
			todo.CompletedAt = &now
		}
	}
	if dueDate != nil {
		todo.DueDate = dueDate
	}
	todo.UpdatedAt = time.Now()
	err = s.repo.Update(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *Service) DeleteTodo(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) CompleteTodo(id string) (*Todo, error) {
	completed := StatusCompleted
	return s.UpdateTodo(id, nil, nil, &completed, nil)
}
