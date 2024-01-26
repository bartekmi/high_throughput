package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDB struct {
	client *dynamodb.Client
}

func NewDynamoDB() *dynamoDB {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &dynamoDB{
		client: client,
	}
}

func (d *dynamoDB) Write(key, content string) error {
	return nil
}

func (d *dynamoDB) Read(key string) (KVPair, bool, error) {
	return KVPair{}, false, nil
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
