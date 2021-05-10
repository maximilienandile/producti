package server

import (
	"github.com/maximilienandile/producti/internal/indexing"
	"github.com/maximilienandile/producti/internal/secret"
	"github.com/maximilienandile/producti/internal/storage"
)

// Config will list all parameters needed to launch the server.
type Config struct {
	Secrets        secret.Parameters
	ProductStore   storage.ProductStore
	ProductIndexer indexing.ProductIndexer
	Stage          string
}
