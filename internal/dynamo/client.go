package dynamo

import (
	"fmt"

	"github.com/maximilienandile/producti/internal/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// an internal DynamoDb client designed to
// make the interaction with the database easier
type client struct {
	dynamoDB  requestor
	tableName string
}

// a standard put request
type putRequest struct {
	object interface{}
	pk     partitionKey
	sk     string
}

// this is the name of the partition key field
const pk = "PK"

// this is the name of the sort key field
const sk = "SK"

// put will insert an item into the database
// the item will be marshalled before insertion
func (c client) put(req putRequest) error {
	marshalled, err := c.marshallInput(req)
	if err != nil {
		return fmt.Errorf("impossible to marshall item: %w", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      marshalled,
		TableName: aws.String(c.tableName),
	}
	_, err = c.dynamoDB.PutItem(input)
	if err != nil {
		return fmt.Errorf("impossible to put item in dynamo: %w", err)
	}
	return nil
}

// marshallInput generates a marshalled version of a put request
// pk and sk are added to the object
func (c client) marshallInput(req putRequest) (map[string]*dynamodb.AttributeValue, error) {
	marshalled, err := dynamodbattribute.MarshalMap(req.object)
	if err != nil {
		return nil, fmt.Errorf("error encountered when attempt to marshall object before put: %w", err)
	}
	// add pk and sk
	marshalled[pk] = &dynamodb.AttributeValue{S: aws.String(string(req.pk))}
	marshalled[sk] = &dynamodb.AttributeValue{S: aws.String(req.sk)}
	return marshalled, nil
}

func (c *client) getByKey(partitionKeyValue partitionKey, sortKeyValue string) (map[string]*dynamodb.AttributeValue, error) {
	result, err := c.dynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			pk: {
				S: aws.String(string(partitionKeyValue)),
			},
			sk: {
				S: aws.String(sortKeyValue),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("impossible to get item in db: %w", err)
	}
	if len(result.Item) == 0 {
		return nil, storage.ErrNotFound
	}
	return result.Item, nil
}
