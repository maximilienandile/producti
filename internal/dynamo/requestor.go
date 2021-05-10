package dynamo

import "github.com/aws/aws-sdk-go/service/dynamodb"

// an interface that lists methods used by the
// dynamoDB client to interact with the database
// it has been added to facilitate mocking
type requestor interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	QueryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error
}
