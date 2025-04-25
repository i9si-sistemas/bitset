package bitset

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestSubstr(t *testing.T) {
	data := []Bit{White, Black, White, Black, White, Black, Black, White}

	tests := []struct {
		start    int
		end      int
		expected []Bit
	}{
		{
			0,
			8,
			[]Bit{White, Black, White, Black, White, Black, Black, White},
		},
		{
			0,
			0,
			[]Bit{},
		},
		{
			0,
			1,
			[]Bit{White},
		},
		{
			2,
			4,
			[]Bit{White, Black},
		},
	}

	for _, test := range tests {
		b := New()
		b.AppendBits(data...)

		result := b.Substr(test.start, test.end)

		expected := New()
		expected.AppendBits(test.expected...)

		assert.True(t, result.Equals(expected))
	}
}
