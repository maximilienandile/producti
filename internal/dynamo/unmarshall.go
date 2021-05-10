package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maximilienandile/producti/internal/product"
)

// unmarshallProduct will take a result from dynamodb and unmarshall it into a variable of type *product.Product
func UnmarshallProduct(dynamoOut map[string]*dynamodb.AttributeValue) (*product.Product, error) {
	productUnmarshalled := product.Product{}
	err := dynamodbattribute.UnmarshalMap(dynamoOut, &productUnmarshalled)
	if err != nil {
		return nil, fmt.Errorf("impossible to unmarshall product: %w", err)
	}
	return &productUnmarshalled, nil
}
