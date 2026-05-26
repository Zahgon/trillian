// Copyright 2016 Google LLC. All Rights Reserved.
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

package cache

import (
	"flag"

	"github.com/google/trillian/storage/storagepb"
	"github.com/google/trillian/storage/tree"
	"github.com/transparency-dev/merkle"
	"github.com/transparency-dev/merkle/compact"
)

// TODO(al): move this up the stack
var populateConcurrency = flag.Int("populate_subtree_concurrency", 256, "Max number of concurrent workers concurrently populating subtrees")

// TODO(pavelkalinnikov): Rename subtrees to tiles.

// GetSubtreesFunc describes a function which can return a number of Subtrees from storage.
type GetSubtreesFunc func(ids [][]byte) ([]*storagepb.SubtreeProto, error)

// SubtreeCache provides a caching access to Subtree storage. Currently there are assumptions
// in the code that all subtrees are multiple of 8 in depth and that log subtrees are always
// of depth 8. It is not possible to just change the constants above and have things still
// work. This is because of issues like byte packing of node IDs.
//
// SubtreeCache is not thread-safe: GetNodes, SetNodes and Flush methods must
// be called sequentially.
type SubtreeCache struct {
	hasher merkle.LogHasher

	// subtrees contains the Subtree data read from storage, and is updated by
	// calls to SetNodes.
	subtrees map[string]*storagepb.SubtreeProto
	// dirtyPrefixes keeps track of all Subtrees which need to be written back
	// to storage.
	dirtyPrefixes map[string]bool

	// populateConcurrency sets the amount of concurrency when repopulating subtrees.
	populateConcurrency int
}

// NewLogSubtreeCache creates and returns a SubtreeCache appropriate for use with a log
// tree. The caller must supply a suitable LogHasher.
func NewLogSubtreeCache(hasher merkle.LogHasher) *SubtreeCache {
	_ = "STUB: not implemented"
	return nil
}

// preload calculates the set of subtrees required to know the hashes of the
// passed in node IDs, uses getSubtrees to retrieve them, and finally populates
// the cache structures with the data. Returns the list of tile IDs not found.
func (s *SubtreeCache) preload(ids []compact.NodeID, getSubtrees GetSubtreesFunc) ([]string, error) {
	_ = "STUB: not implemented"
	// Figure out the set of subtrees we need.
	return nil, nil
}

// Don't make a read request for zero subtrees.

// wait for a token before starting work

// return it when done

// TODO(mhutchinson): This error should be propagated.

// Note: This never blocks because len(ch) == len(subtrees).

func (s *SubtreeCache) cacheSubtree(t *storagepb.SubtreeProto) error {
	_ = "STUB: not implemented"
	return nil
}

// GetNodes returns the requested nodes, calling the getSubtrees function if
// they are not already cached.
func (s *SubtreeCache) GetNodes(ids []compact.NodeID, getSubtrees GetSubtreesFunc) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getNodeHash returns a single node hash from the cache.
func (s *SubtreeCache) getNodeHash(id compact.NodeID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Look up the hash in the appropriate map.
// The leaf hashes are stored in a separate map to the internal nodes so that
// we can easily dump (and later reconstruct) the internal nodes. As log subtrees
// have a fixed depth if the suffix has the same number of significant bits as the
// subtree depth then this is a leaf. For example if the subtree is depth 8 its leaves
// have 8 significant suffix bits.

// SetNodes sets hashes for the given nodes in the cache.
func (s *SubtreeCache) SetNodes(nodes []tree.Node, getSubtrees GetSubtreesFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// Store the hash to the containing tile, and mark it as dirty if the hash
// differs from the previously stored one.

// This is a leaf node.

// This is an internal node.

// UpdatedTiles returns all updated tiles that need to be written to storage.
func (s *SubtreeCache) UpdatedTiles() ([]*storagepb.SubtreeProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
