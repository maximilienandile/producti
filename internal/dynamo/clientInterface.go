package dynamo

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Client interface {
	Put(req PutRequest) error
	GetByKey(partitionKeyValue PartitionKey, sortKeyValue string) (map[string]*dynamodb.AttributeValue, error)
	GetAllByPK(pkValue PartitionKey) ([]map[string]*dynamodb.AttributeValue, error)
}
