package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices(t *testing.T) {
	fmt.Println("In TestSlices")
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 4, 0}

	assert.Equal(t, a, b)    // Identical arrays are equal
	assert.NotEqual(t, a, c) // Different length => not equal

	s1 := a[:]
	s2 := c[:4]
	assert.Equal(t, s1, s2) // Slices equal if iteration has same data

	// Slices are just pointers to an underlying array
	assert.Equal(t, 2, a[1])
	assert.Equal(t, 2, s1[1])
	a[1] = 777
	assert.Equal(t, 777, s1[1])
}
