package dynamo

import "github.com/aws/aws-sdk-go/service/dynamodb"

// DynamoDb Client interface
// on this interface are listed all methods used in the AWS go module to
// interact with the database
type Client interface {
	Put(req PutRequest) error
	GetByKey(partitionKeyValue PartitionKey, sortKeyValue string) (map[string]*dynamodb.AttributeValue, error)
	GetAllByPK(pkValue PartitionKey) ([]map[string]*dynamodb.AttributeValue, error)
}
