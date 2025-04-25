package bitset

import (
	"math/rand"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestLen(t *testing.T) {
	randomBools := make([]bool, 128)

	rng := rand.New(rand.NewSource(1))

	for i := range len(randomBools) {
		randomBools[i] = rng.Intn(2) == 1
	}

	for i := range len(randomBools) - 1 {
		result := New(Bits(randomBools[0:i]...)...)

		assert.Equal(t, result.Len(), i)
	}
}
