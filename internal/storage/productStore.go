package storage

import "github.com/maximilienandile/producti/internal/product"

// ProductStore is an interface that lists all methods to
// store and retrieve Products.
type ProductStore interface {
	// Create will persist a product
	Create(*product.Product) (*product.Product, error)
	// GetByID will retrieve in the persistence layer a product by it's ID
	GetByID(ID string) (*product.Product, error)
	// GetByName will search products by name (full-text search)
	GetByName(name string) ([]*product.Product, error)
}
