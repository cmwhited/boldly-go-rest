# Boldly Go REST

Simple RESTful service to demonstrate the ease of setting a HTTP service using Go.

## Data Storage

This service uses DynamoDB to persist data. Check out the [docs](https://aws.amazon.com/dynamodb/) for more information.

### AWS Access

To access your AWS DynamoDB tables, you will need an AWS account with an IAM user that has access to Read, Write DynamoDB tables.
Once the IAM user is created, get the access key and secret and store them in environment variables:

- `AWS_ACCESS_KEY_ID`: The IAM user access key
- `AWS_SECRET_KEY`: The IAM user secret

## Dependency Management

This service uses [go dep](https://github.com/golang/dep) for the dependency management tool. After pulling the code down,
run `dep ensure`; this will install necessary dependencies to the project and get it ready for running.

## Endpoints

- Ping: Health Check for the service.
    - endpoint: `GET /api/v1/ping`
    - CURL Example
        ```bash
        curl -XGET http://localhost:5002/api/v1/ping
        ```
    - Example Response:
        ```json
        {
            "version": "0.0.1",
            "health": "HEALTHY",
            "msg": "Looking good, beautiful"
        }
        ```
- Get Bank: Get a User Bank Record
    - endpoint: `GET /api/v1/user/{owningUserId}/bank/{bankId}`; where 
        - `{owningUserId}` is the id of the user that the bank belongs to & 
        - `{bankId}` is the unique id of the bank.
    - CURL example
        ```bash
        curl -XGET http://localhost:5002/api/v1/user/4b7b2def-e76e-48bf-993b-8ec2b193b855/bank/01e173f4-02a2-4310-a7cc-e2b919f13aac
        ```
    - Example Response:
        ```json
        {
            "owningUserId": "4b7b2def-e76e-48bf-993b-8ec2b193b855",
            "bankId": "01e173f4-02a2-4310-a7cc-e2b919f13aac",
            "bankName": "US Bank",
            "accountNumber": "2112"
        }
        ```
- Save Bank: Save a new User Bank record
    - endpoint: `POST /api/v1/bank`
    - Example Request Body
        ```json
        {
        	"owningUserId": "4b7b2def-e76e-48bf-993b-8ec2b193b855",
        	"bankName": "BANK NAME",
        	"accountNumber": "1234"
        }
        ```
    - CURL example
        ```bash
        curl -X POST -H "Content-Type: application/json" \
        -d '{"owningUserId": "4b7b2def-e76e-48bf-993b-8ec2b193b855", "bankName": "BANK NAME", "accountNumber": "1234"}' \
        http://localhost:5002/api/v1/bank
        ```
    - Example Response
        ```json
        {
          "owningUserId": "4b7b2def-e76e-48bf-993b-8ec2b193b855",
          "bankId": "b920cfc7-c455-4ac6-b856-f9d3a416d9d1",
          "bankName": "BANK NAME",
          "accountNumber": "1234"
        }
        ```