package indexing

import "github.com/maximilienandile/producti/internal/product"

// Implement this interface to plug search capability on products
// It should be able to add a product to the index and search a product by name
type ProductIndexer interface {
	// AddProduct adds a product in it's indexed representation to the indexClient
	AddProduct(product *product.Indexed) error
	// SearchProductByName will perform a search on products by their name
	SearchProductByName(name string) ([]*product.Indexed, error)
}
