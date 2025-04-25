package bitset

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestClone(t *testing.T) {
	bitset := New(White, Black, White, Black, White, Black, Black, White, Black)
	clone := Clone(bitset)
	assert.Equal(t, bitset.bits, clone.bits)
}
