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

// TileSet represents a set of Merkle tree tiles and the corresponding nodes.
// This type is not thread-safe.
//
// TODO(pavelkalinnikov): Make it immutable.
type TileSet struct {
	layout Layout
	tiles  map[node.ID]NodesRow
	hashes map[node.ID][]byte
	h      mapHasher
}

// NewTileSet creates an empty TileSet with the given tree parameters.
func NewTileSet(treeID int64, hasher Hasher, layout Layout) *TileSet {
	_ = "STUB: not implemented"
	return nil
}

// Hashes returns a map containing all node hashes keyed by node IDs.
func (t *TileSet) Hashes() map[node.ID][]byte {
	_ = "STUB: not implemented"

	// Add puts the given tile into the set. Not thread-safe.
	//
	// TODO(pavelkalinnikov): Take a whole list of Tiles instead.
	return nil
}

func (t *TileSet) Add(tile Tile) error { _ = "STUB: not implemented"; return nil }

// TileSetMutation accumulates tree tiles that need to be updated. This type is
// not thread-safe.
type TileSetMutation struct {
	read  *TileSet
	tiles map[node.ID][]Node
}

// NewTileSetMutation creates a mutation which is based off the provided
// TileSet. This means that each modification is checked against the hashes in
// this set, and is applied if it does change the hash.
func NewTileSetMutation(ts *TileSet) *TileSetMutation { _ = "STUB: not implemented"; return nil }

// Set updates the hash of the given tree node. Not thread-safe.
//
// TODO(pavelkalinnikov): Elaborate on the expected order of Set calls.
// Currently, Build method sorts nodes to allow any order, but it can be
// avoided.
func (t *TileSetMutation) Set(id node.ID, hash []byte) { _ = "STUB: not implemented"; return }

// Nothing changed.

// Not a leaf node of a tile.

// Build returns the full set of tiles modified by this mutation.
func (t *TileSetMutation) Build() ([]Tile, error) { _ = "STUB: not implemented"; return nil, nil }
