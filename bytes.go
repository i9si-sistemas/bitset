package bitset

import "log"

// ByteAt returns the byte at the given index.
func (b *Bitset) ByteAt(index int) byte {
	if index < 0 || index >= b.numBits {
		log.Panicf("Index %d out of range", index)
	}

	var result byte

	for i := index; i < index+8 && i < b.numBits; i++ {
		result <<= 1
		if b.At(i) {
			result |= 1
		}
	}

	return result
}
