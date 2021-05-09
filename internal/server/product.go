package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/maximilienandile/producti/internal/product"
	"github.com/maximilienandile/producti/internal/storage"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (s *Server) GetProductById(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		_ = c.Error(errors.New("no id in path"))
		return
	}
	productFound, err := s.productStore.GetByID(productID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			c.Status(http.StatusNotFound)
			return
		} else {
			// other kind of error
			_ = c.Error(err)
			return
		}
	}
	c.JSON(http.StatusOK, productFound)
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
