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
* Install Serverless globally `npm install -g serverless`
* Install Go [see here](https://golang.org/dl/)
* Install make (should be already installed on MacOS and some Linux distros, installation necessary for Windows)
* 


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


## Mocks Generation


Interfaces of this project are mocked using [GoMock](https://github.com/golang/mock).

Mocks generated are in `interal/mocks`, in the package `mocks`.

To generate a mock of an interface use this command :

```
mockgen -source=internal/storage/productStore.go -destination=internal/mocks/productStore.go -package=mocks
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


## Tasks

- [ ] GET /products
- [ ] GET /product/{id} 
- [ ] GET /product?search=[NAME] 
- [ ] Api documentation
- [ ] Linter
- [ ] Deployment

Extra

- [ ] POST /product


## License
[MIT](https://choosealicense.com/licenses/mit/)