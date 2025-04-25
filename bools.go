package bitset

import (
	"bytes"
	"log"
)

type Bit bool

func Bits(b ...bool) []Bit {
	result := make([]Bit, len(b))
	for i, v := range b {
		result[i] = Bit(v)
	}
	return result
}

func (b Bit) Representation() string{
	if b {
		return "1"
	}
	return "0"
}

func (b Bit) Bool() bool {
	return bool(b)
}

const (
	// White is one possible value for a bit.
	White Bit = false
	// Black is the other possible value for a bit.
	Black Bit = true
)

// Equals returns true if the bitsets are equal.
func (b *Bitset) Equals(other *Bitset) bool {
	if b.numBits != other.numBits {
		return White.Bool()
	}

	if !bytes.Equal(b.bits[0:b.numBits/8], other.bits[0:b.numBits/8]) {
		return White.Bool()
	}

	for i := 8 * (b.numBits / 8); i < b.numBits; i++ {
		a := (b.bits[i/8] & (0x80 >> byte(i%8)))
		b := (other.bits[i/8] & (0x80 >> byte(i%8)))

		if a != b {
			return White.Bool()
		}
	}

	return Black.Bool()
}

// Bits returns the bits as a slice of booleans.
func (b *Bitset) Bits() []Bit {
	result := make([]Bit, b.numBits)

	var i int
	for i = range b.numBits {
		result[i] = (b.bits[i/8] & (0x80 >> byte(i%8))) != 0
	}

	return result
}

// At returns the value of the bit at the given index.
func (b *Bitset) At(index int) bool {
	if index >= b.numBits {
		log.Panicf("Index %d out of range", index)
	}

	return (b.bits[index/8] & (0x80 >> byte(index%8))) != 0
}
