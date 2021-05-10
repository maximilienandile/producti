package dynamo

// this type represents a partition Key used in our database
type PartitionKey string

const (
	ProductPk PartitionKey = "product"
)
