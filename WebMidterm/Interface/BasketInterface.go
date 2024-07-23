package Interface

import (
	"WebMidterm/Model"
)

type Repository interface {
	Migrate() error
	Create(basket Model.Basket) (*Model.Basket, error)
	All() ([]Model.Basket, error)
	GetById(id int64) (*Model.Basket, error)
	Update(id int64, updated Model.Basket) (*Model.Basket, error)
	Delete(id int64) error
}
