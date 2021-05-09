# Producti

Producti is a product API written in Go for e-commerce websites.

## Installation


## Usage


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



## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)