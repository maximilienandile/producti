# Producti

Producti is a product API written in Go for e-commerce websites.


## Compilation

For the `api` lambda :
```
$ go build -o api cmd/api/main.go
```


For the `indexer` lambda :
```
$ go build -o indexer cmd/indexer/main.go
```

## Deployment

### Prerequisites

* Create an AWS Account
* Create an IAM user
* Save your AWS credentials into a named profile. [Instructions Here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html)
* Install Node.js and NPM: [see here](https://nodejs.org/en/download/)
* Install `serverless` globally `npm install -g serverless` (used for the deployment)
* Install `go` [see here](https://golang.org/dl/)
* Install `make` (should be already installed on MacOS and some Linux distros, installation necessary for Windows)
* Install `golangci-lint` [see here](https://golangci-lint.run/usage/install/) (used for the linter)
* Install `newman` [see here](https://www.npmjs.com/package/newman#getting-started) with npm : `npm install -g newman` (used for the functional tests)

``
### Manual Env Setup in AWS

* Create two KMS keys (Customer Managed Keys)
    * For the **develop** env
        * Alias : `ssm-encryption-key-producti-develop`
        * Key Spec : Symmetric
        * Note the Key Id
    * For the **prod** env
        * Alias : `ssm-encryption-key-producti-prod`
        * Key Spec : Symmetric
        * Note the Key Id
* Create two Parameters in SSM (System Manager > Application Management > Parameter Store)
    * For the **develop** env
        * Name : `/producti/develop/secrets`
        * Tier : standard
        * Tpe : SecureString
        * Data Type : text
        * KMS Kay Source : 'My Current Account' => Select key with alias : `ssm-encryption-key-producti-develop`
        * Value : see `doc/secrets.example.json`
    * For the **prod** env
        * Name : `/producti/prod/secrets`
        * Tier : standard
        * Tpe : SecureString
        * Data Type : text
        * KMS Kay Source : 'My Current Account' => Select key with alias : `ssm-encryption-key-producti-prod`
        * Value : see `doc/secrets.example.json`
     
### Deployment 

Note that the DynamoDb table will be created !


#### Prod

```
$ make deploy_prod aws-profile=profileName kms-ssm-key-id=XXXX
```

### Develop

```
$ make deploy_develop aws-profile=profileName kms-ssm-key-id=XXXX
```


## Architecture

The project is deployed to the AWS cloud with the Serverless framework.

![Architecture AWS](doc/architecture-aws.png)

* An API gateway will accept incoming requests
* Those requests are sent to an AWS Lambda (`api`) that will treat incoming requests
* Products are stored into a DynamoDb database
* Thanks to DynamoDb streams, for each new record another Lambda is launched (`indexer`), this lambda will create add an object
 to an Algolia Index
    * to get more info about DynamoDb Streams : [see here](https://www.serverless.com/blog/event-driven-architecture-dynamodb) 
* When a search by product name occurs, a queery is made to the Algolia Index

## Tests

### Unit Tests


To launch unit tests execute :

```
$ make unit_tests 
```

### Functional tests

To launch functional tests first deploy to a web server the API then execute this command :

Example for the develop env (on my AWS account)
```
$ make functional_tests base-url=https://benr4nyn7k.execute-api.eu-central-1.amazonaws.com/develop
```

Example Output :

````
newman

Producti Test

→ Create Product
  POST https://benr4nyn7k.execute-api.eu-central-1.amazonaws.com/develop/product [201 Created, 1.16KB, 193ms]
  ✓  Status test
  ✓  response must be valid and have a body JSON

→ Get Product By ID
  GET https://benr4nyn7k.execute-api.eu-central-1.amazonaws.com/develop/product/8bf888af-079e-43dd-92a0-5e0c6d861d6a [200 OK, 1.16KB, 66ms]
  ✓  Status test
  ✓  response must be valid and have a body JSON

→ Get All Products
  GET https://benr4nyn7k.execute-api.eu-central-1.amazonaws.com/develop/products [200 OK, 1.82KB, 60ms]
  ✓  Status test
  ✓  response must be valid and have a body JSON
  ✓  response is an array

→ Search Product
  GET https://benr4nyn7k.execute-api.eu-central-1.amazonaws.com/develop/product?search=Delacroix [200 OK, 1.14KB, 91ms]
  ✓  Status test
  ✓  response must be valid and have a body JSON
  ✓  response is an array

┌─────────────────────────┬────────────────────┬───────────────────┐
│                         │           executed │            failed │
├─────────────────────────┼────────────────────┼───────────────────┤
│              iterations │                  1 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│                requests │                  4 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│            test-scripts │                  4 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│      prerequest-scripts │                  0 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│              assertions │                 10 │                 0 │
├─────────────────────────┴────────────────────┴───────────────────┤
│ total run duration: 538ms                                        │
├──────────────────────────────────────────────────────────────────┤
│ total data received: 2.8KB (approx)                              │
├──────────────────────────────────────────────────────────────────┤
│ average response time: 102ms [min: 60ms, max: 193ms, s.d.: 53ms] │

````

### Linter

We use golangci-lint (https://golangci-lint.run/usage/quick-start/)

Before launching the linter you need to install it. (see here for instructions : https://golangci-lint.run/usage/install/)

To launch the linter execute :

```
$ make lint
```

## Mocks Generation


Interfaces of this project are mocked using [GoMock](https://github.com/golang/mock).

Mocks generated are in `interal/mocks`, in the package `mocks`.

To generate a mock of an interface use this command :

```
$ make mocks
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


## Tasks

- [x] GET /products
- [x] Sort by Name result of GET /products
- [x] GET /product/{id} 
- [x] GET /product?search=[NAME] 
- [x] API documentation
- [x] Linter
- [x] Deployment

Extra

- [x] POST /product
- [x] Postman Collection for functional tests
- [x] Add Unit tests
- [ ] CI setup Github


## License
[MIT](https://choosealicense.com/licenses/mit/)