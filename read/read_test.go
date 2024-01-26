package main

import (
	"snippetstore/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	// Set up
	s := storage.NewDummy()
	s.Write("k1", "Content 1")
	s.Write("k2", "Content 2")
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
