package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book *Book) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Book, error) {
	return s.repository.FindByID(id)
}

func (s *service) Create(book *Book) error {
	return s.repository.Create(book)
}