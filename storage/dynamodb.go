package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamoDB struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDB(tableName string) *dynamoDB {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &dynamoDB{
		client:    client,
		tableName: tableName,
	}
}

func (d *dynamoDB) Write(data KVPair) error {
	item, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}
	_, err = d.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

func (d *dynamoDB) Read(id string) (KVPair, bool, error) {
	key, err := toKey(id)
	if err != nil {
		return KVPair{}, false, err
	}

	response, err := d.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(d.tableName),
	})
	if err != nil {
		return KVPair{}, false, err
	}
	if len(response.Item) == 0 {
		return KVPair{}, false, nil
	}

	data := KVPair{ID: id}
	err = attributevalue.UnmarshalMap(response.Item, &data)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return data, true, err
}

func toKey(ID string) (map[string]types.AttributeValue, error) {
	val, err := attributevalue.Marshal(ID)
	if err != nil {
		return nil, err
	}
	return map[string]types.AttributeValue{"id": val}, nil
}

func (d *dynamoDB) Count() int {
	return 0
}

func (d *dynamoDB) listTables() ([]string, error) {
	// List DynamoDB tables
	result, err := d.client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		return nil, unwrapError(err)
	}

	tableNames := []string{}
	for _, tableName := range result.TableNames {
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}

func unwrapError(err error) error {
	var builder strings.Builder
	for err != nil {
		builder.WriteString(err.Error() + "\n")
		err = errors.Unwrap(err)
	}

	return fmt.Errorf(builder.String())
}
