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

func TestPut(t *testing.T) {
	// setup mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)

	tableName := "myTable"
	testClient := client{
		dynamoDB:  mockedDynamo,
		tableName: tableName,
	}
	putReq := putRequest{
		object: product.Product{},
		sk:     "458888",
		pk:     productPk,
	}
	marshalled, err := testClient.marshallInput(putReq)
	assert.Nil(t, err)
	expectedInput := &dynamodb.PutItemInput{
		Item:      marshalled,
		TableName: aws.String(tableName),
	}
	// PutItem should be called
	mockedDynamo.EXPECT().PutItem(expectedInput).Return(&dynamodb.PutItemOutput{}, nil)
	err = testClient.put(putReq)
	assert.Nil(t, err)
}

func TestMarshallInput(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)

	testClient := client{
		dynamoDB:  mockedDynamo,
		tableName: "My Table",
	}
	skTest := "458888"
	putReq := putRequest{
		object: product.Product{},
		sk:     skTest,
		pk:     productPk,
	}

	marshalled, err := testClient.marshallInput(putReq)
	assert.Nil(t, err)
	// check that sk and pk are set correctly
	pkRetrieved, found := marshalled[pk]
	assert.True(t, found)
	assert.Equal(t, string(productPk), *pkRetrieved.S)

	skRetrieved, found := marshalled[sk]
	assert.True(t, found)
	assert.Equal(t, skTest, *skRetrieved.S)
}

func TestGetByKeyFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)

	testPk := productPk
	testSK := "42"
	tableName := "myTable"
	testClient := client{
		dynamoDB:  mockedDynamo,
		tableName: tableName,
	}
	mockedDynamo.EXPECT().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			pk: {
				S: aws.String(string(testPk)),
			},
			sk: {
				S: aws.String(testSK),
			},
		},
	}).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			pk: {S: aws.String(string(testPk))},
			sk: {S: aws.String(testSK)},
		},
	}, nil)
	out, err := testClient.getByKey(testPk, testSK)
	assert.Nil(t, err)
	assert.NotNil(t, out)
}

func TestGetAllByPK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedDynamo := mocks.NewMockrequestor(ctrl)
	tableName := "myTable"
	testClient := client{
		dynamoDB:  mockedDynamo,
		tableName: tableName,
	}
	query, err := testClient.queryInputGetByPK(productPk)
	assert.Nil(t, err)
	mockedDynamo.EXPECT().QueryPages(query, gomock.Any()).Return(nil)
	_, err = testClient.getAllByPK(productPk)
	assert.Nil(t, err)
}
