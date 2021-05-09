package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maximilienandile/producti/internal/storage"
)

type Server struct {
	productStore storage.ProductStore
	GinEngine    *gin.Engine
}

func New(conf *Config) *Server {
	s := Server{}
	s.productStore = conf.ProductStore
	// initialize new gin engine
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(s.CorsMiddleware())
	r.Use(s.ErrorHandlerMiddleware())
	// define routes
	AddRoutes(r, &s)
	s.GinEngine = r
	return &s
}
