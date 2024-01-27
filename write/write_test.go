package main

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

	// First Write
	k, err := w.Write(storage.KVPair{Content: "Content 1"})
	assert.Equal(t, 1, storage.GetCount(t, s))
	pair, ok, err := s.Read(k)
	assert.Equal(t, "Content 1", pair.Content)
	assert.True(t, ok)
	assert.Nil(t, err)

	// Second Write
	k, err = w.Write(storage.KVPair{Content: "Content 2", Title: "Title 2"})
	assert.Equal(t, 2, storage.GetCount(t, s))
	pair, ok, err = s.Read(k)
	assert.Equal(t, "Content 2", pair.Content)
	assert.Equal(t, "Title 2", pair.Title)
	assert.True(t, ok)
	assert.Nil(t, err)

	// Error Write
	s.Error = "Simulated Error"
	k, err = w.Write(storage.KVPair{Content: "Content Error"})
	assert.Equal(t, 2, storage.GetCount(t, s))
	assert.Equal(t, "Error writing content: Simulated Error", err.Error())
}

func TestGeneratedUniqueID(t *testing.T) {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9]{3}-[a-zA-Z0-9]{3}-[a-zA-Z0-9]{3}$`)
	for ii := 0; ii < 1000; ii++ {
		id := generatedUniqueID()
		assert.True(t, regex.MatchString(id),
			fmt.Sprintf("Did not match regex: %s", id))
	}
}
