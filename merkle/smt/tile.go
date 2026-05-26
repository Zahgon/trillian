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

package smt

import (
	"github.com/google/trillian/merkle/smt/node"
)

// Tile represents a sparse Merkle tree tile, i.e. a dense set of tree nodes
// located under a single "root" node at a distance not exceeding the tile
// height. A tile is identified by the ID of its root node. The Tile struct
// contains the list of non-empty leaf nodes, which can be used to reconstruct
// all the remaining inner nodes of the tile.
//
// Invariants of this structure that must be preserved at all times:
//   - ID is a prefix of Leaves' IDs, i.e. the nodes are in the same subtree.
//   - IDs of Leaves have the same length, i.e. the nodes are at the same level.
//   - Leaves are ordered by ID from left to right.
//   - IDs of Leaves are unique.
//
// Algorithms that create Tile structures must ensure that these invariants
// hold. Use NewNodesRow function for ordering nodes correctly.
type Tile struct {
	ID     node.ID
	Leaves NodesRow
}

// Merge returns a new tile which is a combination of this tile with the given
// updates. The resulting tile contains all the nodes from the updates, and all
// the nodes from the original tile not present in the updates.
func (t Tile) Merge(updates NodesRow) (Tile, error) {
	_ = "STUB: not implemented"
	return *new(Tile), nil
}

// merge merges two sorted slices of nodes into one sorted slice. If a node ID
// exists in both slices, then the one from the updates slice is taken, i.e. it
// overrides the node from the nodes slice.
func merge(nodes, updates NodesRow) NodesRow { _ = "STUB: not implemented"; return *new(NodesRow) }

// scan visits all non-empty nodes of the tile except the root. The order of
// node visits is arbitrary.
func (t Tile) scan(l Layout, h mapHasher, visit func(Node)) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(pavelkalinnikov): Remove HStar3 side effects, to avoid copying here.
// Currently, the Update method modifies the nodes given to NewHStar3.

// emptyHashes is a NodeAccessor used for computing node hashes of a tile.
type emptyHashes struct {
	h     mapHasher
	visit func(Node)
}

// Get returns an empty hash for the given root node ID.
func (e emptyHashes) Get(id node.ID) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Set calls the visitor callback for the given node and hash.
func (e emptyHashes) Set(id node.ID, hash []byte) { _ = "STUB: not implemented"; return }
