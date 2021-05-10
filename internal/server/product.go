package server

import (
	"errors"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/maximilienandile/producti/internal/product"
	"github.com/maximilienandile/producti/internal/storage"
	uuid "github.com/satori/go.uuid"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (s *Server) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	productFound, err := s.productStore.GetByID(productID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		// other kind of error
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, productFound)
}

func (s *Server) SearchProducts(c *gin.Context) {
	name := c.Query("search")
	if name == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	results, err := s.productIndexer.SearchProductByName(name)
	if err != nil {
		_ = c.Error(err)
		return
	}
	if len(results) == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	// TODO : those requests to the database can be avoided. Put all the data into the Index ? If so index becomes
	// the single source of trust, is it a desirable thing ?
	products := make([]*product.Product, 0, len(results))

	for _, result := range results {
		// retrieve product in database
		p, err := s.productStore.GetByID(result.ProductID)
		if err != nil {
			_ = c.Error(err)
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
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
	newProduct.ID = uuid.NewV4().String()
	productCreated, err := s.productStore.Create(&newProduct)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, productCreated)
}

func (s *Server) GetAllProducts(c *gin.Context) {
	products, err := s.productStore.GetAll()
	if err != nil {
		_ = c.Error(err)
		return
	}
	if len(products) == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	// sort by name
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
	c.JSON(http.StatusOK, products)
}
