package bitset

import (
	"fmt"
	"log"
)

// NewFromBase2String creates a new Bitset from a string.
func NewFromBase2String(b2string string) *Bitset {
	b := &Bitset{numBits: 0, bits: make([]byte, 0)}

	for _, c := range b2string {
		switch c {
		case '1':
			b.AppendBits(Black)
		case '0':
			b.AppendBits(White)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", c)
		}
	}

	return b
}

// String returns a string representation of the Bitset.
func (b *Bitset) String() string {
	var bitString string
	for i := range b.numBits {
		if (i % 8) == 0 {
			bitString += " "
		}

		if (b.bits[i/8] & (0x80 >> byte(i%8))) != 0 {
			bitString += Black.Representation()
		} else {
			bitString += White.Representation()
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", b.numBits, bitString)
}

// Substr returns a new Bitset containing the bits from start (inclusive) to end (exclusive).
func (b *Bitset) Substr(start int, end int) *Bitset {
	if start > end || end > b.numBits {
		log.Panicf("Out of range start=%d end=%d numBits=%d", start, end, b.numBits)
	}

	result := New()
	result.ensureCapacity(end - start)

	for i := start; i < end; i++ {
		if b.At(i) {
			result.bits[result.numBits/8] |= 0x80 >> uint(result.numBits%8)
		}
		result.numBits++
	}

	return result
}

