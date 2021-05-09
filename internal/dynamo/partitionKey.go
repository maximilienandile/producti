package dynamo

// this type represents a partition Key used in our database
type partitionKey string

const (
	productPk partitionKey = "product"
)
