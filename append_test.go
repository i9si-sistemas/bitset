package bitset

import (
	"math/rand"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestAppend(t *testing.T) {
	randomBools := make([]bool, 128)

	rng := rand.New(rand.NewSource(1))

	for i := range len(randomBools) {
		randomBools[i] = rng.Intn(2) == 1
	}

	for i := range len(randomBools)-1 {
		a := New(Bits(randomBools[0:i]...)...)
		b := New(Bits(randomBools[i:]...)...)

		a.Append(b)

		assert.Equal(t, a.Bits(), Bits(randomBools...))
	}
}

func TestAppendByte(t *testing.T) {
	tests := []struct {
		initial  *Bitset
		value    byte
		numBits  int
		expected *Bitset
	}{
		{
			New(),
			0x01,
			1,
			New(Black),
		},
		{
			New(Black),
			0x01,
			1,
			New(Black, Black),
		},
		{
			New(White),
			0x01,
			1,
			New(White, Black),
		},
		{
			New(Black, White, Black, White, Black, White, Black),
			0xAA,
			2,
			New(Black, White, Black, White, Black, White, Black, Black, White),
		},
		{
			New(Black, White, Black, White, Black, White, Black),
			0xAA,
			8,
			New(Black, White, Black, White, Black, White, Black, Black, White, Black, White, Black, White, Black, White),
		},
	}

	for _, test := range tests {
		test.initial.AppendByte(test.value, test.numBits)
		assert.Equal(t, test.initial.Bits(), test.expected.Bits())
	}
}

func TestAppendUint32(t *testing.T) {
	tests := []struct {
		initial  *Bitset
		value    uint32
		numBits  int
		expected *Bitset
	}{
		{
			New(),
			0xAAAAAAAF,
			4,
			New(Black, Black, Black, Black),
		},
		{
			New(),
			0xFFFFFFFF,
			32,
			New(Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black,
				Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black, Black),
		},
		{
			New(),
			0x0,
			32,
			New(White, White, White, White, White, White, White, White, White, White, White, White, White, White, White, White, White,
				White, White, White, White, White, White, White, White, White, White, White, White, White, White, White),
		},
		{
			New(),
			0xAAAAAAAA,
			32,
			New(Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black,
				White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White),
		},
		{
			New(),
			0xAAAAAAAA,
			31,
			New(White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black,
				White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White, Black, White),
		},
	}

	for _, test := range tests {
		test.initial.AppendUint32(test.value, test.numBits)
		assert.Equal(t, test.initial.Bits(), test.expected.Bits())
	}
}

func TestAppendBools(t *testing.T) {
	randomBools := make([]bool, 128)

	rng := rand.New(rand.NewSource(1))

	for i := range len(randomBools) {
		randomBools[i] = rng.Intn(2) == 1
	}

	for i := range len(randomBools) - 1 {
		result := New(Bits(randomBools[0:i]...)...)
		result.AppendBools(randomBools[i:]...)

		assert.Equal(t, result.Bits(), Bits(randomBools...))
	}
}

func BenchmarkShortAppend(b *testing.B) {
	bitset := New()

	for range b.N {
		bitset.AppendBits(White, Black, White, Black, White, Black, White)
	}
}
