package bitset

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestEnsureCapacity(t *testing.T) {
	bitset := New()
	bitset.ensureCapacity(100)
	assert.Equal(t, len(bitset.bits), 13)
}
