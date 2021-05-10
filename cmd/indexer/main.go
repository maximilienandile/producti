package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/maximilienandile/producti/internal/indexing"
	"github.com/maximilienandile/producti/internal/secret"

	"github.com/maximilienandile/producti/internal/product"

	"github.com/maximilienandile/producti/internal/dynamo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var indexer indexing.ProductIndexer

func handler(ctx context.Context, event events.DynamoDBEvent) error {
	for _, record := range event.Records {
		// get PK
		pkRetrieved, found := record.Change.Keys[dynamo.Pk]
		if !found {
			return errors.New("impossible to handle stream PK is net set")
		}
		if pkRetrieved.String() != string(dynamo.ProductPk) {
			return errors.New("only PK == product is supported yet for dynamo streams")
		}
		// get the product Name
		productID, found := record.Change.NewImage[dynamo.Sk]
		if !found {
			return errors.New("impossible to handle stream 'SK' is net set")
		}
		productName, found := record.Change.NewImage["name"]
		if !found {
			return errors.New("impossible to handle stream 'name' is net set")
		}
		productIndexed := product.Indexed{
			ProductID: productID.String(),
			Name:      productName.String(),
		}
		err := indexer.AddProduct(&productIndexed)
		if err != nil {
			return fmt.Errorf("impossible to add product to index: %w", err)
		}
	}
	return nil
}

func main() {
	// load secrets from SSM parameter store
	parameterStoreName, found := os.LookupEnv("PARAMETER_STORE_NAME")
	if !found {
		log.Fatal("impossible to start server no PARAMETER_STORE_NAME env value")
	}
	outSSM, err := ssm.New(session.Must(session.NewSession())).GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameterStoreName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("impossible to get SSM parameters %s", err)
	}
	// unmarshall secrets
	var secrets secret.Parameters
	err = json.Unmarshal([]byte(*outSSM.Parameter.Value), &secrets)
	if err != nil {
		log.Fatalf("impossible to unmarshall SSM parameters %s", err)
	}
	indexer = indexing.NewAlgoliaProductIndexer(secrets.Algolia)
	lambda.Start(handler)
}
