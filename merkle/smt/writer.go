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
	"context"

	"github.com/google/trillian/merkle/smt/node"
)

// NodeBatchAccessor reads and writes batches of Merkle tree node hashes. It is
// a batch interface for efficiency reasons, as it is designed to guard tree
// storage / database access. The Writer type operates on a per-shard basis,
// i.e. it calls Get and Set method exactly once for each shard.
type NodeBatchAccessor interface {
	// Get returns the hashes of the given nodes, as a map keyed by their IDs.
	// The returned hashes may be missing or be nil for empty subtrees.
	Get(ctx context.Context, ids []node.ID) (map[node.ID][]byte, error)
	// Set applies the given node hash updates.
	Set(ctx context.Context, nodes []Node) error
}

// Writer handles sharded writes to a sparse Merkle tree. The tree has two
// levels of shards: the single topmost shard spanning depths from 0 to split,
// and 2^split second-level shards each spanning levels from split to height.
// If the split height is 0 then effectively there is only one "global" shard.
type Writer struct {
	h      mapHasher
	height uint // The height of the tree.
	split  uint // The height of the top shard.
}

// NewWriter creates a new Writer for the specified tree of the given height,
// with two levels of sharding, where the upper shard is `split` levels high.
func NewWriter(treeID int64, hasher Hasher, height, split uint) *Writer {
	_ = "STUB: not implemented"
	return nil
}

// Split sorts and splits the given list of node hash updates into shards, i.e.
// the subsets belonging to different subtrees. The nodes must belong to the
// same tree level which is equal to the tree height.
func (w *Writer) Split(nodes []Node) ([][]Node, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO(pavelkalinnikov): Try estimating the capacity for this slice.

// The nodes are sorted, so we can split them by prefix.

// Check if this ID ends the shard.

// Write applies the given list of node updates to a single shard, and returns
// the resulting update of the shard root. It uses the given node accessor for
// reading and writing tree nodes.
//
// The typical usage pattern is as follows. For the lower shards, the input is
// the []Node slices returned from the Split method. For the top shard, the
// input is all the Node values from the lower shards Write calls.
//
// In another case, Write can be performed without Split if the shard split
// depth is 0, which effectively means that there is only one "global" shard.
func (w *Writer) Write(ctx context.Context, nodes []Node, acc NodeBatchAccessor) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

// shardTop returns the depth of a shard top based on its bottom depth.
func (w *Writer) shardTop(depth uint) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// newAccessor returns a NodeAccessor for HStar3 algorithm based on the set of
// preloaded node hashes.
func (w *Writer) newAccessor(nodes map[node.ID][]byte) *shardAccessor {
	_ = "STUB: not implemented"
	// For any node that HStar3 reads, it also writes its sibling. Therefore we
	// can pre-allocate this many items for the writes slice.
	// TODO(pavelkalinnikov): The actual number of written nodes will be slightly
	// bigger by at most the number of written leaves. Try allocating precisely.
	return nil
}

// shardAccessor provides read and write access to nodes used by HStar3. It
// operates entirely in-memory.
type shardAccessor struct {
	w      *Writer
	reads  map[node.ID][]byte
	writes []Node
}

// Get returns the hash of the given node from the preloaded map, or a hash of
// an empty subtree at this position if such node is not found.
func (s *shardAccessor) Get(id node.ID) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Set adds the given node hash update to the list of writes.
func (s *shardAccessor) Set(id node.ID, hash []byte) { _ = "STUB: not implemented"; return }
