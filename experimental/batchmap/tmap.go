// Copyright 2020 Google LLC. All Rights Reserved.
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

// Package batchmap is a library to be used within Beam pipelines to construct
// verifiable data structures.
package batchmap

//go:generate go install github.com/apache/beam/sdks/v2/go/cmd/starcgen
//go:generate starcgen --package=batchmap --identifiers=entryToNodeHashFn,partitionByPrefixLenFn,tileHashFn,leafShardFn,tileToNodeHashFn,tileUpdateFn

import (
	"context"
	"crypto"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"

	"github.com/google/trillian/merkle/coniks"
	"github.com/google/trillian/merkle/smt"
	"github.com/google/trillian/merkle/smt/node"
)

var (
	cntTilesHashed  = beam.NewCounter("batchmap", "tiles-hashed")
	cntTilesCopied  = beam.NewCounter("batchmap", "tiles-copied")
	cntTilesCreated = beam.NewCounter("batchmap", "tiles-created")
	cntTilesUpdated = beam.NewCounter("batchmap", "tiles-updated")
)

func init() {
	register.DoFn1x2[nodeHash, []byte, nodeHash](&leafShardFn{})
	register.DoFn3x2[context.Context, []byte, func(*nodeHash) bool, *Tile, error](&tileHashFn{})
	register.DoFn4x2[context.Context, []byte, func(**Tile) bool, func(*nodeHash) bool, *Tile, error](&tileUpdateFn{})
	register.Function5x1(createStratum)
	register.Function6x1(updateStratum)
	register.Function1x2(tilePathFn)
}

// Create builds a new map from the given PCollection of *Entry. Outputs
// the resulting Merkle tree tiles as a PCollection of *Tile.
//
// The keys in the input PCollection must be 256-bit, uniformly distributed,
// and unique within the input.
// The values in the input PCollection must be 256-bit.
// treeID should be a unique ID for the lifetime of this map. This is used as
// part of the hashing algorithm to provide preimage resistance. If the tiles
// are to be imported into Trillian for serving, this must match the tree ID
// within Trillian.
// The internal hash algorithm can be picked between SHA256 and SHA512_256.
// The internal nodes will use this algorithm via the CONIKS strategy.
// prefixStrata is the number of 8-bit prefix strata. Any path from root to leaf
// will have prefixStrata+1 tiles.
func Create(s beam.Scope, entries beam.PCollection, treeID int64, hash crypto.Hash, prefixStrata int) (beam.PCollection, error) {
	_ = "STUB: not implemented"
	return *new(beam.PCollection), nil
}

// Construct the map pipeline starting with the leaf tiles.

// Collate all of the strata together and return them.

// Update takes an existing base map (PCollection of *Tile), applies the
// delta (PCollection of *Entry) and returns the resulting map as a
// PCollection of *Tile.
// The deltas can add new keys to the map or overwrite existing keys. Keys
// cannot be deleted (though their value can be set to a sentinel value).
//
// treeID, hash, and prefixStrata must match the values passed into the
// original call to Create that started the base map.
func Update(s beam.Scope, base, delta beam.PCollection, treeID int64, hash crypto.Hash, prefixStrata int) (beam.PCollection, error) {
	_ = "STUB: not implemented"
	return *new(beam.PCollection), nil
}

// Tile sets returned from this library have tiles present at all byte
// lengths from [0..prefixStrata]. This makes this a perfect partition fn.

// Construct the map pipeline starting with the leaf tiles.

// Collate all of the strata together and return them.

// createStratum creates the tiles for the stratum at the given rootDepth bytes.
// leaves is a PCollection of nodeHash that are the leaves of this layer.
// output is a PCollection of *Tile.
func createStratum(s beam.Scope, leaves beam.PCollection, treeID int64, hash crypto.Hash, rootDepth int) beam.PCollection {
	_ = "STUB: not implemented"
	return *new(beam.PCollection)
}

// updateStratum updates the tiles for the stratum at the given bytes depth.
// base is a PCollection of *Tile which is the tiles in the stratum
// to be updated.
// deltas is a PCollection of nodeHash that are the updated leaves of this layer.
// output is a PCollection of *Tile.
func updateStratum(s beam.Scope, base, deltas beam.PCollection, treeID int64, hash crypto.Hash, rootDepth int) beam.PCollection {
	_ = "STUB: not implemented"
	return *new(beam.PCollection)
}

func tilePathFn(t *Tile) ([]byte, *Tile) {
	_ = "STUB: not implemented"

	// nodeHash describes a leaf to be included in a tile.
	// This is logically the same as smt.Node however it has public fields so is
	// serializable by the default Beam coder. Also, it allows changes to be made
	// to smt.Node without affecting this, which improves decoupling.
	return nil, nil
}

type nodeHash struct {
	// Path from root of the map to this node. Equivalent to node.ID, but with
	// the significant benefit that it will be serialized properly without
	// writing a custom coder for nodeHash.
	Path []byte
	Hash []byte
}

func partitionByPrefixLenFn(t *Tile) int { _ = "STUB: not implemented"; return 0 }

func tileToNodeHashFn(t *Tile) nodeHash { _ = "STUB: not implemented"; return *new(nodeHash) }

func entryToNodeHashFn(e *Entry) nodeHash { _ = "STUB: not implemented"; return *new(nodeHash) }

// leafShardFn groups nodeHashs together based on the first RootDepthBytes
// bytes of their path. This groups all leaves from the same tile together.
type leafShardFn struct {
	RootDepthBytes int
}

func (fn *leafShardFn) ProcessElement(leaf nodeHash) ([]byte, nodeHash) {
	_ = "STUB: not implemented"
	return nil, *new(nodeHash)
}

type tileHashFn struct {
	TreeID int64
	Hash   crypto.Hash
	th     *tileHasher
}

func (fn *tileHashFn) Setup() { _ = "STUB: not implemented"; return }

func (fn *tileHashFn) ProcessElement(ctx context.Context, rootPath []byte, leaves func(*nodeHash) bool) (*Tile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertNodes consumes the Beam-style iterator of nodeHash and returns the
// corresponding slice of smt.Node. Nothing clever is attempted to ensure that
// the data structure will fit in memory. If the iterator has too many elements
// then this will cause an out of memory panic. It is up to the library client
// to configure the map with an appropriate number of prefix strata such that
// this does not occur.
func convertNodes(leaves func(*nodeHash) bool) ([]smt.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tileUpdateFn merges the base tile from the original map with the deltas that
// represent the changes to the map. Note this only supports additions or
// overwrites. There is no ability to delete a leaf.
type tileUpdateFn struct {
	TreeID int64
	Hash   crypto.Hash
	th     *tileHasher
}

func (fn *tileUpdateFn) Setup() { _ = "STUB: not implemented"; return }

func (fn *tileUpdateFn) ProcessElement(ctx context.Context, rootPath []byte, bases func(**Tile) bool, deltas func(*nodeHash) bool) (*Tile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If there are no deltas, then the base tile is unchanged.

func (fn *tileUpdateFn) updateTile(rootPath []byte, base *Tile, deltas []smt.Node) (*Tile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tileHasher is an smt.NodeAccessor used for computing node hashes of a tile.
// This is not serializable and must be constructed within each worker stage.
type tileHasher struct {
	treeID int64
	h      *coniks.Hasher
}

func (th *tileHasher) construct(rootPath []byte, nodes []smt.Node) (*Tile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// N.B. This needs to be done after Prepare but BEFORE HStar3 because it
// fiddles around with the nodes and makes their IDs invalid afterwards.

func (th *tileHasher) update(rootPath []byte, baseNodes, deltaNodes []smt.Node) (*Tile, error) {
	_ = "STUB: not implemented"
	// We add new values first and then update with base to easily check for duplicates in deltas.
	return nil, nil
}

// Only add base values if they haven't been updated.

// hashTile computes the root hash of the root given the prepared leaves.
// The leaves slice MUST NOT be used after calling this method.
func (th *tileHasher) hashTile(depthBits uint, leaves []smt.Node) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get returns hash of an empty subtree for the given root node ID.
func (th tileHasher) Get(id node.ID) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (th tileHasher) Set(id node.ID, hash []byte) { _ = "STUB: not implemented"; return }

func nodeID2Encode(n node.ID) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func nodeID2Decode(bs []byte) (node.ID, error) {
	_ = "STUB: not implemented"
	return *new(node.ID), nil
}

// getOptionalTile consumes the Beam-style iterator and returns:
// - nil if there were no entries
// - the single tile if there was only one entry
// - an error if there were multiple entries
func getOptionalTile(iter func(**Tile) bool) (*Tile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only at most one entry is found.
// Note: Returns nil if found nothing.
