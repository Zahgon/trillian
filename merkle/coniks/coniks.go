// Copyright 2017 Google LLC. All Rights Reserved.
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

// Package coniks provides CONIKS hashing for maps.
package coniks

import (
	"bytes"
	"crypto"

	"github.com/google/trillian/merkle/smt/node"
)

// Domain separation prefixes
var (
	leafIdentifier  = []byte("L")
	emptyIdentifier = []byte("E")
	// Default is the standard CONIKS hasher.
	Default = New(crypto.SHA512_256)
	// Some zeroes, to avoid allocating temporary slices.
	zeroes = make([]byte, 32)
)

// Hasher implements the sparse merkle tree hashing algorithm specified in the CONIKS paper.
type Hasher struct {
	crypto.Hash
}

// New creates a new hashers.TreeHasher using the passed in hash function.
func New(h crypto.Hash) *Hasher { _ = "STUB: not implemented"; return nil }

// EmptyRoot returns the root of an empty tree.
func (m *Hasher) EmptyRoot() []byte { _ = "STUB: not implemented"; return nil }

// HashEmpty returns the hash of an empty subtree with the given root.
func (m *Hasher) HashEmpty(treeID int64, root node.ID) []byte {
	_ = "STUB: not implemented"
	return nil

	// bytes.Buffer never returns errors so we can ignore them below.
}

// HashLeaf calculate the merkle tree leaf value:
// H(Identifier || treeID || depth || index || dataHash)
func (m *Hasher) HashLeaf(treeID int64, id node.ID, leaf []byte) []byte {
	_ = "STUB: not implemented"
	return nil

	// bytes.Buffer never returns errors so we can ignore them below.
}

// HashChildren returns the internal Merkle tree node hash of the two child nodes l and r.
// The hashed structure is  H(l || r).
func (m *Hasher) HashChildren(l, r []byte) []byte { _ = "STUB: not implemented"; return nil }

// BitLen returns the number of bits in the hash function.
func (m *Hasher) BitLen() int { _ = "STUB: not implemented"; return 0 }

// writeMaskedNodeID writes the node ID bits to the buffer, padded with zero
// bits to the byte Size() of the hashes in use by this hasher.
//
// TODO(pavelkalinnikov): We must not use BitLen() and Size() interchangeably.
// The tree height and hash size could be different.
// TODO(pavelkalinnikov): Padding with zeroes doesn't buy us anything, as the
// depth is also written to the Buffer.
func (m *Hasher) writeMaskedNodeID(b *bytes.Buffer, id node.ID) { _ = "STUB: not implemented"; return }

// Write the complete bytes.

// Mask off unwanted bits in the last byte, if there is an incomplete one.

// Pad to the correct length with zeroes. Allow for future hashers that might
// be > 256 bits.
// TODO(pavelkalinnikov): YAGNI. Simplify this until that actually happens.

// Use the pre-allocated zeroes to avoid allocating them each time.
