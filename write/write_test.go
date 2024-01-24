package write

import (
	"fmt"
	"regexp"
	"snippetstore/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	// Set up
	s := storage.NewDummy()
	w := New(s)

	// Test
	k1 := w.Write("Content 1")
	k2 := w.Write("Content 2")

	// Verify
	assert.Equal(t, 2, s.Count())

	pair, ok := s.Read(k1)
	assert.True(t, ok)
	assert.Equal(t, "Content 1", pair.Content)

	pair, ok = s.Read(k2)
	assert.True(t, ok)
	assert.Equal(t, "Content 2", pair.Content)
}

func TestGeneratedUniqueID(t *testing.T) {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9]{3}-[a-zA-Z0-9]{3}-[a-zA-Z0-9]{3}$`)
	for ii := 0; ii < 1000; ii++ {
		id := generatedUniqueID()
		assert.True(t, regex.MatchString(id),
			fmt.Sprintf("Did not match regex: %s", id))
	}
}
