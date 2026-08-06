package todo

import (
	"encoding/json"
	"os"
)

type JSONStorage struct {
	path string
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{path: path}
}

func (j *JSONStorage) Load() ([]Todo, error) {
	data, err := os.ReadFile(j.path)

	if os.IsNotExist(err) {
    return []Todo{}, nil
	}
	
	if err != nil {
    return nil, err
}

	if len(data) == 0 {
    return nil, nil
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)

	if(err != nil) {
		return nil, err
	}

	return todos, nil
}

func (j *JSONStorage) Save(t []Todo) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}	
	
	err = os.WriteFile(j.path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}