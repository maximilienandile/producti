package dynamo

import (
	"testing"

	"github.com/maximilienandile/producti/internal/product"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestUnmarshallProduct(t *testing.T) {
	input := map[string]*dynamodb.AttributeValue{
		"id":         {S: aws.String("15487-545")},
		"name":       {S: aws.String("Leather")},
		"followers":  {N: aws.String("42")},
		"daysOnline": {N: aws.String("154")},
	}
	actual, err := UnmarshallProduct(input)
	assert.Nil(t, err)
	expected := &product.Product{
		ID:         "15487-545",
		Name:       "Leather",
		Followers:  42,
		DaysOnline: 154,
	}
	assert.Equal(t, expected, actual)
}
