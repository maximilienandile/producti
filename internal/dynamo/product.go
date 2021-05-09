package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maximilienandile/producti/internal/product"
	"github.com/maximilienandile/producti/internal/storage"
	uuid "github.com/satori/go.uuid"
)

// An implementation of storage.ProductStore with DynamoDb
type ProductStore struct {
	client client
}

// Initialize the the product Store
func NewProductStore(tableName string, awsSession *session.Session) storage.ProductStore {
	repo := ProductStore{
		client: client{
			dynamodb.New(awsSession),
			tableName,
		},
	}
	return &repo
}

// Create will persist a product object into the database
// the field ID is set with an UUID v4
func (p ProductStore) Create(product *product.Product) (*product.Product, error) {
	product.ID = uuid.NewV4().String()
	err := p.client.put(putRequest{
		object: product,
		pk:     productPk,
		sk:     product.ID,
	})
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get a product by it's ID
func (p ProductStore) GetByID(ID string) (*product.Product, error) {
	out, err := p.client.getByKey(productPk, ID)
	if err != nil {
		return nil, fmt.Errorf("impossible to get product by id: %w", err)
	}
	return p.unmarshallProduct(out)
}

func (p ProductStore) GetByName(name string) ([]*product.Product, error) {
	panic("implement me")
}

// GetAll retrieve all the products stored
func (p ProductStore) GetAll() ([]*product.Product, error) {
	all, err := p.client.getAllByPK(productPk)
	if err != nil {
		return nil, fmt.Errorf("impossible to retrieve all products: %w", err)
	}
	allProducts := make([]*product.Product, 0, len(all))
	for k, v := range all {
		productUnmarshalled, err := p.unmarshallProduct(v)
		if err != nil {
			return nil, fmt.Errorf("impossible to unmarshall product at index %d: %w", k, err)
		}
		allProducts = append(allProducts, productUnmarshalled)
	}
	return allProducts, nil
}

// unmarshallProduct will take a result from dynamodb and unmarshall it into a variable of type *product.Product
func (p ProductStore) unmarshallProduct(out map[string]*dynamodb.AttributeValue) (*product.Product, error) {
	productUnmarshalled := product.Product{}
	err := dynamodbattribute.UnmarshalMap(out, &productUnmarshalled)
	if err != nil {
		return nil, fmt.Errorf("impossible to unmarshall product: %w", err)
	}
	return &productUnmarshalled, nil
}
