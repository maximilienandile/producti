build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/api cmd/api/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/indexer cmd/indexer/main.go
clean:
	rm -rf ./bin
deploy_develop: clean build
	sls deploy --verbose \
    	--aws-profile maxaldtools \
    	--stage develop
deploy_prod: clean build
	sls deploy --verbose \
    	--aws-profile maxaldtools \
    	--stage prod