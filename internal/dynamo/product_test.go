package dynamo

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/stretchr/testify/assert"

	"github.com/maximilienandile/producti/internal/product"

	"github.com/golang/mock/gomock"
)

func TestProductStore_Create(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockClient(ctrl)
	testProductStore := ProductStore{client: m}
	input := product.Product{
		ID:   "12544",
		Name: "Tot Bag",
	}
	m.EXPECT().Put(PutRequest{
		object: &input,
		pk:     ProductPk,
		sk:     "12544",
	}).Return(nil)

	productCreated, err := testProductStore.Create(&input)
	assert.Nil(t, err)
	assert.Equal(t, "12544", productCreated.ID)
	assert.Equal(t, "Tot Bag", productCreated.Name)
}

func TestProductStore_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockClient(ctrl)
	testProductStore := ProductStore{client: m}
	dynamoOut := []map[string]*dynamodb.AttributeValue{
		{
			"id":   {S: aws.String("1245")},
			"name": {S: aws.String("Tot Bag")},
		},
	}
	m.EXPECT().GetAllByPK(ProductPk).Return(dynamoOut, nil)
	actual, err := testProductStore.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(actual))
}

func TestProductStore_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockClient(ctrl)
	searchID := "42"
	dynamoOut := map[string]*dynamodb.AttributeValue{
		"id":   {S: aws.String(searchID)},
		"name": {S: aws.String("Tot Bag")},
	}

	testProductStore := ProductStore{client: m}
	m.EXPECT().GetByKey(ProductPk, searchID).Return(dynamoOut, nil)
	actual, err := testProductStore.GetByID(searchID)
	assert.Nil(t, err)
	assert.Equal(t, "42", actual.ID)
	assert.Equal(t, "Tot Bag", actual.Name)
}
