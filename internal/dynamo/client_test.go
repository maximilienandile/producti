package dynamo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/golang/mock/gomock"
	"github.com/maximilienandile/producti/internal/mocks"

	"github.com/maximilienandile/producti/internal/product"
)

const tableName = "myTable"

func TestPut(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)

	testClient := simpleClient{
		requestor: mockedDynamo,
		tableName: tableName,
	}
	putReq := PutRequest{
		object: product.Product{},
		sk:     "458888",
		pk:     ProductPk,
	}
	marshalled, err := testClient.marshallInput(putReq)
	assert.Nil(t, err)
	expectedInput := &dynamodb.PutItemInput{
		Item:      marshalled,
		TableName: aws.String(tableName),
	}
	// PutItem should be called
	mockedDynamo.EXPECT().PutItem(expectedInput).Return(&dynamodb.PutItemOutput{}, nil)
	err = testClient.Put(putReq)
	assert.Nil(t, err)
}

func TestMarshallInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)

	testClient := simpleClient{
		requestor: mockedDynamo,
		tableName: "My Table",
	}
	skTest := "458888"
	putReq := PutRequest{
		object: product.Product{},
		sk:     skTest,
		pk:     ProductPk,
	}

	marshalled, err := testClient.marshallInput(putReq)
	assert.Nil(t, err)
	// check that Sk and Pk are set correctly
	pkRetrieved, found := marshalled[Pk]
	assert.True(t, found)
	assert.Equal(t, string(ProductPk), *pkRetrieved.S)

	skRetrieved, found := marshalled[Sk]
	assert.True(t, found)
	assert.Equal(t, skTest, *skRetrieved.S)
}

func TestGetByKeyFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)

	testPk := ProductPk
	testSK := "42"
	tableName := "myTable"
	testClient := simpleClient{
		requestor: mockedDynamo,
		tableName: tableName,
	}
	mockedDynamo.EXPECT().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			Pk: {
				S: aws.String(string(testPk)),
			},
			Sk: {
				S: aws.String(testSK),
			},
		},
	}).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			Pk: {S: aws.String(string(testPk))},
			Sk: {S: aws.String(testSK)},
		},
	}, nil)
	out, err := testClient.GetByKey(testPk, testSK)
	assert.Nil(t, err)
	assert.NotNil(t, out)
}

func TestGetAllByPK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)
	tableName := "myTable"
	testClient := simpleClient{
		requestor: mockedDynamo,
		tableName: tableName,
	}
	query, err := testClient.queryInputGetByPK(ProductPk)
	assert.Nil(t, err)
	mockedDynamo.EXPECT().QueryPages(query, gomock.Any()).Return(nil)
	_, err = testClient.GetAllByPK(ProductPk)
	assert.Nil(t, err)
}
