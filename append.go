package bitset

import "log"

// Append appends the given Bitset to the current Bitset.
func (b *Bitset) Append(other *Bitset) {
	b.ensureCapacity(other.numBits)

	for i := range other.numBits {
		if other.At(i) {
			b.bits[b.numBits/8] |= 0x80 >> uint(b.numBits%8)
		}
		b.numBits++
	}
}

// AppendBits appends the given bits to the current Bitset.
func (b *Bitset) AppendBits(bits ...Bit) {
	b.ensureCapacity(len(bits))

	for _, v := range bits {
		if v {
			b.bits[b.numBits/8] |= 0x80 >> uint(b.numBits%8)
		}
		b.numBits++
	}
}

// AppendBools appends the given bits to the current Bitset.
func (b *Bitset) AppendBools(bits ...bool) {
	b.ensureCapacity(len(bits))

	for _, v := range bits {
		if v {
			b.bits[b.numBits/8] |= 0x80 >> uint(b.numBits%8)
		}
		b.numBits++
	}
}

// AppendNummBools appends the given number of bits to the current Bitset.
func (b *Bitset) AppendNumBools(num int, value bool) {
	for range num {
		b.AppendBools(value)
	}
}

// AppendBytes appends the given bytes to the current Bitset.
func (b *Bitset) AppendBytes(data []byte) {
	for _, d := range data {
		b.AppendByte(d, 8)
	}
}

// AppendByte appends the given byte to the current Bitset.
func (b *Bitset) AppendByte(value byte, numBits int) {
	b.ensureCapacity(numBits)

	if numBits > 8 {
		log.Panicf("numBits %d out of range 0-8", numBits)
	}

	for i := numBits - 1; i >= 0; i-- {
		if value&(1<<uint(i)) != 0 {
			b.bits[b.numBits/8] |= 0x80 >> uint(b.numBits%8)
		}

		b.numBits++
	}
}

// AppendUint32 appends the given uint32 to the current Bitset.
func (b *Bitset) AppendUint32(value uint32, numBits int) {
	b.ensureCapacity(numBits)

	if numBits > 32 {
		log.Panicf("numBits %d out of range 0-32", numBits)
	}

	for i := numBits - 1; i >= 0; i-- {
		if value&(1<<uint(i)) != 0 {
			b.bits[b.numBits/8] |= 0x80 >> uint(b.numBits%8)
		}

		b.numBits++
	}
}
