package invoiceitem

import "time"

// Model of invoiceitem

type Model struct {
	ID            uint
	invoiceheader uint
	ProductID     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//Models slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetByID() (*Model, error)
	//Delete(uint) error
}

//Service of invoice
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceitem
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
