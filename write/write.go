package write

import (
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

func (w *Writer) Write(content string) string {
	id := generatedUniqueID()
	w.storage.Write(id, content)
	return id
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
