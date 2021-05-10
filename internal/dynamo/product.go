package dynamo

import (
	"fmt"

	awsClient "github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maximilienandile/producti/internal/product"
	"github.com/maximilienandile/producti/internal/storage"
)

// An implementation of storage.ProductStore with DynamoDb
type ProductStore struct {
	client Client
}

// Initialize the the product Store
func NewProductStore(tableName string, provider awsClient.ConfigProvider) storage.ProductStore {
	repo := ProductStore{
		client: &simpleClient{
			dynamodb.New(provider),
			tableName,
		},
	}
	return &repo
}

// Create will persist a product object into the database
func (p ProductStore) Create(product *product.Product) (*product.Product, error) {
	err := p.client.Put(PutRequest{
		object: product,
		pk:     ProductPk,
		sk:     product.ID,
	})
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get a product by it's ID
func (p ProductStore) GetByID(ID string) (*product.Product, error) {
	out, err := p.client.GetByKey(ProductPk, ID)
	if err != nil {
		return nil, fmt.Errorf("impossible to get product by id: %w", err)
	}
	return UnmarshallProduct(out)
}

// GetAll retrieve all the products stored
func (p ProductStore) GetAll() ([]*product.Product, error) {
	all, err := p.client.GetAllByPK(ProductPk)
	if err != nil {
		return nil, fmt.Errorf("impossible to retrieve all products: %w", err)
	}
	allProducts := make([]*product.Product, 0, len(all))
	for k, v := range all {
		productUnmarshalled, err := UnmarshallProduct(v)
		if err != nil {
			return nil, fmt.Errorf("impossible to unmarshall product at index %d: %w", k, err)
		}
		allProducts = append(allProducts, productUnmarshalled)
	}
	return allProducts, nil
}
