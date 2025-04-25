package bitset

// Clone returns a copy of the given bitset.
func Clone(from *Bitset) *Bitset {
	return &Bitset{numBits: from.numBits, bits: from.bits[:]}
}