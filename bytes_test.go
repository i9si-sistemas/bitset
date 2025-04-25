package bitset

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestByteAt(t *testing.T) {
	white := White.Bool()
	black := Black.Bool()
	data := []bool{
		white, 
		black, 
		white, 
		black, 
		white, 
		black, 
		black, 
		white, 
		black,
	}

	tests := []struct {
		index    int
		expected byte
	}{
		{
			0,
			0x56,
		},
		{
			1,
			0xad,
		},
		{
			2,
			0x2d,
		},
		{
			5,
			0x0d,
		},
		{
			8,
			0x01,
		},
	}

	for _, test := range tests {
		b := New()
		b.AppendBools(data...)

		result := b.ByteAt(test.index)
		assert.Equal(t, result, test.expected)
	}
}
