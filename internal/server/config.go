package server

import (
	"github.com/maximilienandile/producti/internal/indexing"
	"github.com/maximilienandile/producti/internal/storage"
)

// Config will list all parameters needed to launch the server.
type Config struct {
	ProductStore   storage.ProductStore
	ProductIndexer indexing.ProductIndexer
}
