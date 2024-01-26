package main

import (
	"fmt"
	"math/rand"
	"snippetstore/storage"
	"strings"
)

type Writer struct {
	storage storage.Storage
}

func New(s storage.Storage) *Writer {
	return &Writer{storage: s}
}

// Writes content and returns an id, guaranteed to be
// reasonably short and unique.
// data.ID is ignored.
func (w *Writer) Write(data storage.KVPair) (string, error) {
	id := generatedUniqueID()
	data.ID = id
	// TODO... Check for prior existence!!!

	err := w.storage.Write(data)
	if err != nil {
		return "", fmt.Errorf("Error writing content: %s", err)
	}

	return id, nil
}

// We want to accommodate 1T unique id's with a 1:1000
// probability of duplicate when one is generated at random.
// ID's are made up of [a-zA-Z0-9] = 26 + 26 + 10 = 62 characters
// 9 charactes will do the trick. For human readability, we will
// lay out the chars as 123-456-789
func generatedUniqueID() string {
	sb := &strings.Builder{}

	generateRandomRunes(sb, 3)
	sb.WriteRune('-')
	generateRandomRunes(sb, 3)
	sb.WriteRune('-')
	generateRandomRunes(sb, 3)

	return sb.String()
}

func generateRandomRunes(sb *strings.Builder, repeat int) {
	max := 26*2 + 10

	for i := 0; i < repeat; i++ {
		r := rand.Intn(max)
		var ascii int

		if r < 26 {
			ascii = 'a' + r
		} else if r < 26*2 {
			ascii = 'A' + (r - 26)
		} else {
			ascii = '0' + (r - 26*2)
		}

		sb.WriteRune(rune(ascii))
	}
}
