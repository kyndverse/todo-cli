package main

import (
	"log"
	"os"
	"path/filepath"
	"todo-cli/internal/todo"
	"todo-cli/internal/tui"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dataPath := filepath.Join(
		filepath.Dir(exe),
		"todos.json",
	)

	storage := todo.NewJSONStorage(dataPath)

	service, err := todo.New(storage)
	if err != nil {
		log.Fatal(err)
	}
	
	err = tui.RunApp(service)
	if err != nil {
		log.Fatal(err)
	}	
}