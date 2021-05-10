package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/maximilienandile/producti/internal/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// an internal DynamoDb client designed to
// make the interaction with the database easier
type client struct {
	requestor requestor
	tableName string
}

// a standard Put request
type PutRequest struct {
	object interface{}
	pk     PartitionKey
	sk     string
}

// this is the name of the partition key field
const Pk = "PK"

// this is the name of the sort key field
const Sk = "SK"

// put will insert an item into the database
// the item will be marshalled before insertion
func (c client) Put(req PutRequest) error {
	marshalled, err := c.marshallInput(req)
	if err != nil {
		return fmt.Errorf("impossible to marshall item: %w", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      marshalled,
		TableName: aws.String(c.tableName),
	}
	_, err = c.requestor.PutItem(input)
	if err != nil {
		return fmt.Errorf("impossible to Put item in dynamo: %w", err)
	}
	return nil
}

// marshallInput generates a marshalled version of a put request
// Pk and Sk are added to the object
func (c client) marshallInput(req PutRequest) (map[string]*dynamodb.AttributeValue, error) {
	marshalled, err := dynamodbattribute.MarshalMap(req.object)
	if err != nil {
		return nil, fmt.Errorf("error encountered when attempt to marshall object before Put: %w", err)
	}
	// add Pk and Sk
	marshalled[Pk] = &dynamodb.AttributeValue{S: aws.String(string(req.pk))}
	marshalled[Sk] = &dynamodb.AttributeValue{S: aws.String(req.sk)}
	return marshalled, nil
}

// this method allow you to get an item in dynamo by PK and SK
func (c *client) GetByKey(partitionKeyValue PartitionKey, sortKeyValue string) (map[string]*dynamodb.AttributeValue, error) {
	result, err := c.requestor.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			Pk: {
				S: aws.String(string(partitionKeyValue)),
			},
			Sk: {
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

func (c *client) queryInputGetByPK(pkValue PartitionKey) (*dynamodb.QueryInput, error) {
	keyCondition := expression.Key(Pk).Equal(expression.Value(pkValue))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}
	params := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(c.tableName),
	}
	return params, nil
}

// this method allow you to retrieve all items with a given PK
// it will get ALL the elements with the provided PK.
func (c *client) GetAllByPK(pkValue PartitionKey) ([]map[string]*dynamodb.AttributeValue, error) {
	params, err := c.queryInputGetByPK(pkValue)
	if err != nil {
		return nil, err
	}
	itemsRetrieved := make([]map[string]*dynamodb.AttributeValue, 0)
	pageNum := 0
	//exec query
	err = c.requestor.QueryPages(params, func(page *dynamodb.QueryOutput, lastPage bool) bool {
		itemsRetrieved = append(itemsRetrieved, page.Items...)
		pageNum++
		return true
	})
	if err != nil {
		return nil, err
	}
	return itemsRetrieved, nil
}
