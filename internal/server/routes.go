package server

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.Engine, s *Server) {
	r.GET("/product/:id", s.GetProductById)
	r.GET("/product", s.SearchProducts)
	r.GET("/products", s.GetAllProducts)
	r.POST("/product", s.CreateProduct)
}
