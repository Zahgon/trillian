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

package testonly

import (
	"context"

	"github.com/google/trillian/storage/tree"
	"github.com/transparency-dev/merkle/compact"
)

// This is a fake implementation of a NodeReader intended for use in testing Merkle path code.
// Building node sets for tests by hand is onerous and error prone, especially when trying
// to test code reading from multiple tree revisions. It cannot live in the main testonly
// package as this creates import cycles.

// FakeNodeReader is an implementation of storage.NodeReader that's preloaded with a set of
// NodeID -> Node mappings and will return only those. Requesting any other nodes results in
// an error. For use in tests only, does not implement any other storage APIs.
type FakeNodeReader struct {
	nodeMap map[compact.NodeID]tree.Node
}

// NewFakeNodeReader creates and returns a FakeNodeReader with the supplied nodes
// assuming that all the nodes are at a specified tree revision. All the node IDs
// must be distinct.
func NewFakeNodeReader(nodes []tree.Node) *FakeNodeReader { _ = "STUB: not implemented"; return nil }

// Duplicate mapping - the test data is invalid so don't continue.

// GetMerkleNodes implements the corresponding NodeReader API.
func (f FakeNodeReader) GetMerkleNodes(ids []compact.NodeID) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FakeNodeReader) hasID(id compact.NodeID) bool { _ = "STUB: not implemented"; return false }

// MultiFakeNodeReader can provide nodes at multiple revisions. It delegates to a number of
// FakeNodeReaders, each set up to handle one revision.
type MultiFakeNodeReader struct {
	readers []FakeNodeReader
}

// LeafBatch describes a set of leaves to be loaded into a MultiFakeNodeReader via a compact
// merkle tree. As each batch is added to the tree a set of node updates are collected
// and recorded in a FakeNodeReader for that revision. The expected root should be the
// result of calling CurrentRoot() on the compact Merkle tree encoded by hex.EncodeToString().
type LeafBatch struct {
	TreeRevision int64
	Leaves       []string
	ExpectedRoot []byte
}

// NewMultiFakeNodeReader creates a MultiFakeNodeReader delegating to a number of FakeNodeReaders
func NewMultiFakeNodeReader(readers []FakeNodeReader) *MultiFakeNodeReader {
	_ = "STUB: not implemented"
	return nil
}

// NewMultiFakeNodeReaderFromLeaves uses a compact Merkle tree to set up the nodes at various
// revisions. It collates all node updates from a batch of leaf data into one FakeNodeReader.
// This has the advantage of not needing to manually create all the data structures but the
// disadvantage is that a bug in the compact tree could be reflected in test using this
// code. To help guard against this we check the tree root hash after each batch has been
// processed. The supplied batches should be in ascending order of tree revision.
func NewMultiFakeNodeReaderFromLeaves(batches []LeafBatch) *MultiFakeNodeReader {
	_ = "STUB: not implemented"
	return nil
}

// Store the new leaf node, and all new perfect nodes.

// TODO(pavelkalinnikov): Use testing.T.Fatalf instead of panics.

// Sanity check the tree root hash against the one we expect to see.

// Unroll the update map to []tree.Node to retain the most recent node update within
// the batch for each ID. Use that to create a new FakeNodeReader.

func (m MultiFakeNodeReader) readerForNodeID(id compact.NodeID) *FakeNodeReader {
	_ = "STUB: not implemented"
	// Work backwards and use the first reader where the node is present.
	return nil
}

// GetMerkleNodes implements the corresponding NodeReader API.
func (m MultiFakeNodeReader) GetMerkleNodes(ctx context.Context, ids []compact.NodeID) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	// Find the correct reader for the supplied tree revision. This must be done for each node
	// as earlier revisions may still be relevant
	return nil, nil
}
