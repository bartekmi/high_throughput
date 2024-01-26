package main

import (
	"snippetstore/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	// Set up
	s := storage.NewDummy()
	s.Write(storage.KVPair{ID: "k1", Content: "Content 1"})
	s.Write(storage.KVPair{ID: "k2", Content: "Content 2", Title: "Title 2"})
	r := New(s)

	// Read k1
	p, ok, err := r.Read("k1")
	assert.True(t, ok)
	assert.Equal(t, "Content 1", p.Content)
	assert.Nil(t, err)

	// Read k2
	p, ok, err = r.Read("k2")
	assert.True(t, ok)
	assert.Equal(t, "Content 2", p.Content)
	assert.Nil(t, err)

	// Read Error
	s.Error = "Simulated Error"
	p, ok, err = r.Read("k1")
	assert.Equal(t, "Simulated Error", err.Error())
}
