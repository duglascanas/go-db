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
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID() (*Model, error)
	Delete(uint) error
}
