package todo

import (
	"errors"
	"strings"
)

type Service struct {
	storage Storage
	todos []Todo
	nextID int
}

func New(s Storage) (*Service, error) {
	todos, err  := s.Load()
	if err != nil {
		return nil, err
	}

	return &Service{
		storage: s,
		todos: todos,
		nextID: calculateNextID(todos),
	}, nil
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

	err := s.storage.Save(s.todos)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Done(id int) error {
	for i := range s.todos {
    if s.todos[i].ID == id {
      s.todos[i].IsCompleted = true

			err := s.storage.Save(s.todos)
			if err != nil {
				return err
			}

      return nil
    }
	}

	return ErrTodoNotFound
}

func (s *Service) Delete(id int) error {
	for i := range s.todos {
    if s.todos[i].ID == id {
      s.todos = append(s.todos[:i], s.todos[i+1:]...)

			err := s.storage.Save(s.todos)
			if err != nil {
				return err
			}

      return nil
    }
	}

	return ErrTodoNotFound
}

func calculateNextID(todos []Todo) int {
	if len(todos) == 0 {
		return  1
	}

	id := 0

	for _, todo := range todos {
		if todo.ID > id {
			id = todo.ID
		}
	}

	return id + 1
}