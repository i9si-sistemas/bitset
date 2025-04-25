package bitset

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestAt(t *testing.T) {
	test := []Bit{White, Black, White, Black, White, Black, Black, White, Black}

	bitset := New(test...)
	for i := range test {
		result := bitset.At(i)
		assert.Equal(t, result, test[i].Bool())
	}
}
