package todo

type Repository interface {
	Create(todo *Todo) error

	GetAll() ([]Todo, error)

	GetByID(id string) (*Todo, error)

	Update(todo *Todo) error

	Delete(id string) error
}
