package todo

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	service := New()

	if service == nil {
		t.Fatal("expected service, got nil")
	}

	todos := service.List()

	if len(todos) != 0 {
    t.Fatal("expected length 1, got 0")
	}
}

func TestAdd(t *testing.T) {
	
	t.Run("valid description", func(t *testing.T) {
		service := New()

		err := service.Add("Belajar Go")
		
		if err != nil {
			t.Fatal("expected error nil, got error not nil")
		} 
			
		todos := service.List()
		if len(todos) == 0 {
			t.Errorf("expected 0, got %d", len(todos)) 
		} else if todos[0].Description != "Belajar Go" {
			t.Errorf("expected Belajar Go, got %q", todos[0].Description)
		} else if todos[0].IsCompleted != false {
			t.Error("expected false, got true")
		} else if todos[0].ID != 1 {
			t.Errorf("expected 1, got %d", todos[0].ID)
		} 
  })

  t.Run("empty description", func(t *testing.T) {
		service := New()
		
		err := service.Add("")
		todos := service.List()
		
		if err == nil || len(todos) != 0 {
			t.Error("expected error not nil, got nil")
		} 
  })
}

func TestDone(t *testing.T)  {
	t.Run("valid id", func(t *testing.T) {
		service := New()
		service.Add("Belajar golang")

		todos := service.List()
		err := service.Done(todos[0].ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		
		todos = service.List()
		if todos[0].IsCompleted != true {
			t.Error("expected isComplete true, got false")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		service := New()
		service.Add("Belajar golang")
		
		err := service.Done(99)
		if !errors.Is(err, ErrTodoNotFound) {
			t.Fatal("expected ErrTodoNotFound")
		}
	})
	
	t.Run("double call", func(t *testing.T) {
		service := New()
		service.Add("Belajar golang")
		
		todos := service.List()	
		err := service.Done(todos[0].ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		
		err = service.Done(todos[0].ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		todos = service.List()
		if !todos[0].IsCompleted {
			t.Error("expected todo to remain completed")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		service := New()
		service.Add("Belajar Golang")
		
		todos := service.List()
		err := service.Delete(todos[0].ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		service := New()
		service.Add("Belajar Golang")
		
		err := service.Delete(99)
		if !errors.Is(err, ErrTodoNotFound) {
			t.Fatal("expected ErrTodoNotFound")
		}
	})
}