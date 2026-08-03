package todo

import (
	"errors"
	"strings"
)

type Service struct {
	todos []Todo
	nextID int
}

func New() *Service {
	return &Service{
		nextID: 1,
	}
}

func (s *Service) List() []Todo {
	result := make([]Todo, len(s.todos))
	copy(result, s.todos)
	return result
}

func (s *Service) Add(description string) error {
	if strings.TrimSpace(description) == "" {
		return errors.New("empty description")
	}

	todo := Todo{
    ID: s.nextID,
    Description: description,
    IsCompleted: false,
	}

	s.todos = append(s.todos, todo)
	s.nextID++

	return nil
}

func (s *Service) Done(id int) error {
	for i := range s.todos {
    if s.todos[i].ID == id {
      s.todos[i].IsCompleted = true
      return nil
    }
	}

	return ErrTodoNotFound
}

func (s *Service) Delete(id int) error {
	for i := range s.todos {
    if s.todos[i].ID == id {
      s.todos = append(s.todos[:i], s.todos[i+1:]...)
      return nil
    }
	}

	return ErrTodoNotFound
}
