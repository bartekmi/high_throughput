package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnection(t *testing.T) {
	// Set up
	dd := NewDynamoDB("high_throughput_test.snippets")
	tableNames, err := dd.listTables()

	if err != nil {
		require.Fail(t, err.Error())
	}
	require.Equal(t, []string{"high_throughput.snippets", "high_throughput_test.snippets"}, tableNames)
}

func MyFunc() {
	// do nothing
}

func TestReadWrite(t *testing.T) {
	// Set up
	dd := NewDynamoDB(DYNAMODB_TABLE_TEST)
	require.Nil(t, dd.deleteAll())

	// First Write
	err := dd.Write(KVPair{ID: "ID1", Content: "Content 1"})
	require.Nil(t, err)
	require.Equal(t, 1, GetCount(t, dd))
	pair, ok, err := dd.Read("ID1")
	require.Nil(t, err)
	require.True(t, ok)
	require.Equal(t, "Content 1", pair.Content)
	require.True(t, ok)
	require.Nil(t, err)

	// Second Write
	err = dd.Write(KVPair{ID: "ID2", Content: "Content 2", Title: "Title 2"})
	require.Nil(t, err)
	require.Equal(t, 2, GetCount(t, dd))
	pair, ok, err = dd.Read("ID2")
	require.Nil(t, err)
	require.True(t, ok)
	require.Equal(t, "Content 2", pair.Content)
	require.Equal(t, "Title 2", pair.Title)
	require.True(t, ok)
	require.Nil(t, err)
}

func TestReadMissingID(t *testing.T) {
	dd := NewDynamoDB(DYNAMODB_TABLE_TEST)

	_, ok, err := dd.Read("DOES NOT EXIST")
	require.Nil(t, err)
	require.False(t, ok)
}
