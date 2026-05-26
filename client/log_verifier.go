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

package client

import (
	"github.com/google/trillian"
	"github.com/google/trillian/types"
	"github.com/transparency-dev/merkle"
)

// LogVerifier allows verification of output from Trillian Logs, both regular
// and pre-ordered; it is safe for concurrent use (as its contents are fixed
// after construction).
type LogVerifier struct {
	// hasher is the hash strategy used to compute nodes in the Merkle tree.
	hasher merkle.LogHasher
}

// NewLogVerifier returns an object that can verify output from Trillian Logs.
func NewLogVerifier(hasher merkle.LogHasher) *LogVerifier { _ = "STUB: not implemented"; return nil }

// NewLogVerifierFromTree creates a new LogVerifier using the algorithms
// specified by a Trillian Tree object.
func NewLogVerifierFromTree(config *trillian.Tree) (*LogVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyRoot verifies that newRoot is a valid append-only operation from
// trusted. If trusted.TreeSize is zero, a consistency proof is not needed.
func (c *LogVerifier) VerifyRoot(trusted *types.LogRootV1, newRoot *trillian.SignedLogRoot, consistency [][]byte) (*types.LogRootV1, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Implicitly trust the first root we get.

// Verify consistency proof.

// VerifyInclusionByHash verifies that the inclusion proof for the given Merkle leafHash
// matches the given trusted root.
func (c *LogVerifier) VerifyInclusionByHash(trusted *types.LogRootV1, leafHash []byte, pf *trillian.Proof) error {
	_ = "STUB: not implemented"
	return nil
}
