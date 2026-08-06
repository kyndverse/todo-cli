package todo

type Storage interface {
	Load() ([]Todo, error)
	Save([]Todo) error
}