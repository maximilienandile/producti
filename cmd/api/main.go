package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/maximilienandile/producti/internal/indexing"

	"github.com/maximilienandile/producti/internal/dynamo"

	"github.com/maximilienandile/producti/internal/secret"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/maximilienandile/producti/internal/server"
)

var ginLambda *ginadapter.GinLambda

func init() {
	tableName, found := os.LookupEnv("DYNAMO_TABLE_NAME")
	if !found {
		log.Fatal("impossible to start server no TABLE_NAME env value")
	}
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

	activateAwsLogsStr, ok := os.LookupEnv("ACTIVATE_AWS_LOGS")
	if !ok {
		log.Fatal("impossible to start server no LOG_LEVEL_AWS env value")
	}
	activateAwsLogs, err := strconv.ParseBool(activateAwsLogsStr)
	if err != nil {
		log.Fatal("impossible to parse to boolean ACTIVATE_AWS_LOGS")
	}
	logLevel := aws.LogOff
	if activateAwsLogs {
		logLevel = aws.LogDebug
	}
	awsSession := session.Must(session.NewSession(&aws.Config{
		LogLevel: &logLevel,
	}))
	config := server.Config{
		ProductStore:   dynamo.NewProductStore(tableName, awsSession),
		ProductIndexer: indexing.NewAlgoliaProductIndexer(secrets.Algolia),
	}
	s := server.New(&config)
	ginLambda = ginadapter.New(s.GinEngine)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
