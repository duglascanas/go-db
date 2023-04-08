package product

import "time"

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
