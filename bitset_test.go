package bitset

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestNewBitset(t *testing.T) {
	tests := [][]Bit{
		{},
		{Black},
		{White},
		{Black, White},
		{Black, White, Black},
		{White, White, Black},
	}

	for _, v := range tests {
		result := New(v...)

		assert.Equal(t, result.Bits(), v)
	}
}

func TestExample(t *testing.T) {
	b := New()
	b.AppendBits(Black, Black, White)
	b.AppendBits(Black)
	b.AppendByte(0x02, 4)

	expected := []Bit{Black, Black, White, Black, White, White, Black, White}

	assert.Equal(t, b.Bits(), expected)
}
