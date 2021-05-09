package server

import (
	"github.com/maximilienandile/producti/internal/secret"
	"github.com/maximilienandile/producti/internal/storage"
)

// Config will list all parameters needed to launch the server
type Config struct {
	Secrets      secret.Parameters
	ProductStore storage.ProductStore
	Stage        string
}
