package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

func (p ProductStore) GetByID(ID string) (*product.Product, error) {
	panic("implement me")
}

func (p ProductStore) GetByName(name string) ([]*product.Product, error) {
	panic("implement me")
}
