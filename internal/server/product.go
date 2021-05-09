package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/maximilienandile/producti/internal/product"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (s *Server) GetProductById(c *gin.Context) {
	panic("todo")

}

func (s *Server) SearchProducts(c *gin.Context) {
	panic("todo")

}

func (s *Server) CreateProduct(c *gin.Context) {
	var newProduct product.Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		_ = c.Error(err)
		return
	}
	err = validate.Struct(newProduct)
	if err != nil {
		_ = c.Error(err)
		return
	}
	productCreated, err := s.productStore.Create(&newProduct)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, productCreated)
}

func (s *Server) GetAllProducts(c *gin.Context) {
	panic("todo")

}
