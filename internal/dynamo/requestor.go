package dynamo

import "github.com/aws/aws-sdk-go/service/dynamodb"

type requestor interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
