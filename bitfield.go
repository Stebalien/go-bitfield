package bitfield

import (
	newbf "github.com/ipfs/go-bitfield"
)

// Deprecated: Use github.com/ipfs/go-bitfield.Bitfield
type Bitfield = newbf.Bitfield

// NewBitfield creates a new fixed-sized Bitfield (allocated up-front).
//
// Panics if size is not a multiple of 8.
//
// Deprecated: Use github.com/ipfs/go-bitfield.NewBitfield
func NewBitfield(size int) Bitfield {
	return newbf.NewBitfield(size)
}

// FromBytes constructs a new bitfield from a serialized bitfield.
//
// Deprecated: Use github.com/ipfs/go-bitfield.FromBytes
func FromBytes(size int, bits []byte) Bitfield {
	return newbf.FromBytes(size, bits)
}
