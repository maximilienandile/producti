package storage

import (
	"errors"

	"github.com/maximilienandile/producti/internal/product"
)

// ErrNotFound is returned when no results are found
var ErrNotFound = errors.New("element not found")

// ProductStore is an interface that lists all methods to
// store and retrieve Products.
type ProductStore interface {
	// Create will persist a product
	Create(*product.Product) (*product.Product, error)
	// GetByID will retrieve in the persistence layer a product by it's ID
	GetByID(ID string) (*product.Product, error)
	// GetAll will retrieve all products
	GetAll() ([]*product.Product, error)
}
