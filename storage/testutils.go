package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func GetCount(t *testing.T, s Storage) int {
	count, err := s.Count()
	if err != nil {
		require.Fail(t, err.Error())
	}
	return count
}
