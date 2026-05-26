// Copyright 2019 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package node implements a sparse Merkle tree node.
package node

// ID identifies a node of a Merkle tree. It is a bit string that counts the
// node down from the tree root, i.e. 0 and 1 bits represent going to the left
// or right child correspondingly.
//
// ID is immutable, comparable, and can be used as a Golang map key. It also
// incurs zero memory allocations in transforming methods like Prefix and
// Sibling.
//
// The internal structure of ID is driven by its use-cases:
//   - To make ID objects immutable and comparable, the Golang string type is
//     used for storing the bit string bytes.
//   - To make Sibling and Prefix operations fast, the last byte is stored
//     separately from the rest of the bytes, so that it can be "amended".
//   - To make ID objects comparable, there is only one (canonical) way to encode
//     an ID. For example, if the last byte is used partially, its unused bits
//     are always unset. See invariants next to field definitions below.
//
// Constructors and methods of ID make sure its invariants are always met.
//
// For example, an 11-bit node ID [1010,1111,001] is structured as follows:
// - path string contains 1 byte, which is [1010,1111].
// - last byte is [0010,0000]. Note the unset lower 5 bits.
// - bits is 3, so effectively only the upper 3 bits [001] of last are used.
type ID struct {
	path string
	last byte  // Invariant: Lowest (8-bits) bits of the last byte are unset.
	bits uint8 // Invariant: 1 <= bits <= 8, or bits == 0 for the empty ID.
}

// NewID creates a node ID from the given path bytes truncated to the specified
// number of bits if necessary. Panics if the number of bits is more than the
// byte string contains.
func NewID(path string, bits uint) ID { _ = "STUB: not implemented"; return *new(ID) }

// Note: Getting the substring is cheap because strings are immutable in Go.

// NewIDWithLast creates a node ID from the given path bytes and the additional
// last byte, of which only the specified number of most significant bits is
// used. The number of bits must be between 1 and 8, and can be 0 only if the
// path bytes string is empty; otherwise the function panics.
func NewIDWithLast(path string, last byte, bits uint8) ID {
	_ = "STUB: not implemented"
	return *new(ID)
}

// newMaskedID constructs a node ID ensuring its invariants are met. The last
// byte is masked so that the given number of upper bits are in use, and the
// others are unset.
func newMaskedID(path string, last byte, bits uint8) ID { _ = "STUB: not implemented"; return *new(ID) }

// Unset the unused bits.

// BitLen returns the length of the ID in bits.
func (n ID) BitLen() uint { _ = "STUB: not implemented"; return 0 }

// FullBytes returns the ID bytes that are complete. Note that there might
// still be up to 8 extra bits, which can be obtained with the LastByte method.
func (n ID) FullBytes() string {
	_ = "STUB: not implemented"

	// LastByte returns the terminating byte of the ID, with the number of upper
	// bits that it uses (between 1 and 8, and 0 if the ID is empty). The remaining
	// unused lower bits are always unset.
	return ""
}

func (n ID) LastByte() (byte, uint8) {
	_ = "STUB: not implemented"
	return 0,

		// Prefix returns the prefix of the node ID with the given number of bits.
		0
}

func (n ID) Prefix(bits uint) ID {
	_ = "STUB: not implemented"
	// Note: This code is very similar to NewID, and it's tempting to return
	// NewID(n.path, bits). But there is a difference: NewID expects all the
	// bytes to be in the path string, while here the last byte is not.
	return *new(ID)
}

// Sibling returns the ID of the nodes's sibling in a binary tree, i.e. the ID
// of the parent node's other child. If the node is the root then the returned
// ID is the same.
func (n ID) Sibling() ID { _ = "STUB: not implemented"; return *new(ID) }

// String returns a human-readable bit string.
func (n ID) String() string { _ = "STUB: not implemented"; return "" }

// Trim the brackets.

// split returns the decomposition of an ID with the given number of bits. The
// first int returned is the number of full bytes stored in the dynamically
// allocated part. The second one is the number of bits in the tail byte.
func split(bits uint) (bytes uint, tailBits uint8) { _ = "STUB: not implemented"; return 0, 0 }
