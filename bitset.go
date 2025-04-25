package bitset

// Bitset is a set of bits.
type Bitset struct {
	// numBits is the number of bits in the set.
	numBits int
	// bits is the underlying byte array.
	bits    []byte
}

// New returns a new Bitset.
func New(v ...Bit) *Bitset {
	b := &Bitset{numBits: 0, bits: make([]byte, 0)}
	b.AppendBits(v...)

	return b
}
