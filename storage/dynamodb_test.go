package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDynamoDB(t *testing.T) {
	// Set up
	dd := NewDynamoDB()
	tableNames, err := dd.listTables()

	if err != nil {
		require.Fail(t, err.Error())
	}
	require.Equal(t, []string{"high_throughput.snippets"}, tableNames)
}
