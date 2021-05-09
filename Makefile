# Todo : aws profile dynamic
build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/api cmd/api/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/indexer cmd/indexer/main.go
clean:
	rm -rf ./bin
deploy_develop: clean build
	sls deploy --verbose \
    	--aws-profile $(aws-profile) \
    	--kms-ssm-key-id $(kms-ssm-key-id) \
    	--stage develop
deploy_prod: clean build
	sls deploy --verbose \
    	--aws-profile $(aws_profile) \
    	--kms-ssm-key-id $(kms-ssm-key-id) \
    	--stage prod
unit_tests:
	go test ./...
lint:
	golangci-lint run --enable-all
mocks:
	mockgen -source=internal/storage/productStore.go -destination=internal/mocks/productStore.go -package=mocks
	mockgen -source=internal/dynamo/requestor.go -destination=internal/mocks/requestor.go -package=mocks