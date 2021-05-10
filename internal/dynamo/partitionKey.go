package dynamo

// this type represents a partition Key used in our database
type PartitionKey string

const (
	// Used to store products
	ProductPk PartitionKey = "product"
)
