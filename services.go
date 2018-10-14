package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/satori/go.uuid"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

func DoPing() Ping {
	return Ping{
		Version: "0.0.1",
		Health:  "HEALTHY",
		Msg:     "Looking good, beautiful",
	}
}

// Get a unique Bank record by the owningUserId and bankId composite key.
// Use the AWS DynamoDB service instance to query the UserBanks table in DynamoDB by the Primary & Sort key
func GetBank(owningUserId, bankId uuid.UUID) (bank *Bank, err error) {
	req := awsSvc.DynamoDbSvc().GetItemRequest(&dynamodb.GetItemInput{
		TableName: aws.String("UserBanks"),
		Key: map[string]dynamodb.AttributeValue{
			"owningUserId": {
				S: aws.String(owningUserId.String()),
			},
			"bankId": {
				S: aws.String(bankId.String()),
			},
		},
	})
	output, err := req.Send()
	if err != nil {
		return nil, err
	}
	// unmarshal map
	bank = new(Bank)
	err = dynamodbattribute.UnmarshalMap(output.Item, &bank)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

// Save the Bank record to the UserBanks table in DynamoDB.
// Use the AWS DynamoDB Service to marshal the bank and save the item in the table
func (bank *Bank) SaveBank() (*Bank, error) {
	bank.BankId = uuid.NewV4().String() // set bank id
	// marshal to map for dynamo input
	bankMap, err := dynamodbattribute.MarshalMap(bank)
	if err != nil {
		return nil, err
	}
	// build put item request
	input := &dynamodb.PutItemInput{
		Item:      bankMap,
		TableName: aws.String("UserBanks"),
	}
	// save item to db
	req := awsSvc.DynamoDbSvc().PutItemRequest(input)
	_, err = req.Send()
	if err != nil {
		return nil, err
	}
	return bank, nil
}
