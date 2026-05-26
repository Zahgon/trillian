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

package memory

import (
	"context"
	"sync"
	"time"

	"github.com/google/btree"
	"github.com/google/trillian"
	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	stree "github.com/google/trillian/storage/tree"
	"github.com/google/trillian/types"
	"github.com/transparency-dev/merkle/compact"
)

const logIDLabel = "logid"

var (
	once            sync.Once
	queuedCounter   monitoring.Counter
	dequeuedCounter monitoring.Counter
)

func createMetrics(mf monitoring.MetricFactory) { _ = "STUB: not implemented"; return }

func labelForTX(t *logTreeTX) string { _ = "STUB: not implemented"; return "" }

// unseqKey formats a key for use in a tree's BTree store.
// The associated Item value will be a list of unsequenced entries.
func unseqKey(treeID int64) btree.Item { _ = "STUB: not implemented"; return *new(btree.Item) }

// seqLeafKey formats a key for use in a tree's BTree store.
// The associated Item value will be the leaf at the given sequence number.
func seqLeafKey(treeID, seq int64) btree.Item { _ = "STUB: not implemented"; return *new(btree.Item) }

// hashToSeqKey formats a key for use in a tree's BTree store.
// The associated Item value will be the sequence number for the leaf with
// the given hash.
func hashToSeqKey(treeID int64) btree.Item { _ = "STUB: not implemented"; return *new(btree.Item) }

// sthKey formats a key for use in a tree's BTree store.
// The associated Item value will be the STH with the given timestamp.
func sthKey(treeID int64, timestamp uint64) btree.Item {
	_ = "STUB: not implemented"
	return *new(btree.Item)
}

// revKey formats a key for use in a tree's BTree store. The associated Item
// value will be the revision number for the given timestamp.
func revKey(treeID int64, timestamp uint64) btree.Item {
	_ = "STUB: not implemented"
	return *new(btree.Item)
}

type memoryLogStorage struct {
	*TreeStorage
	metricFactory monitoring.MetricFactory
}

// NewLogStorage creates an in-memory LogStorage instance.
func NewLogStorage(ts *TreeStorage, mf monitoring.MetricFactory) storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

func (m *memoryLogStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"

	// GetActiveLogIDs returns the IDs of all logs that are currently in a state
	// that requires sequencing (e.g. ACTIVE, DRAINING).
	return nil
}

func (m *memoryLogStorage) GetActiveLogIDs(ctx context.Context) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *memoryLogStorage) beginInternal(ctx context.Context, tree *trillian.Tree, readonly bool) (*logTreeTX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *memoryLogStorage) ReadWriteTransaction(ctx context.Context, tree *trillian.Tree, f storage.LogTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

/* readonly */

func (m *memoryLogStorage) AddSequencedLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *memoryLogStorage) SnapshotForTree(ctx context.Context, tree *trillian.Tree) (storage.ReadOnlyLogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyLogTreeTX), nil
}

/* readonly */

func (m *memoryLogStorage) QueueLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* readonly */

// Ensure we don't leak the transaction. For example if we get an
// ErrTreeNeedsInit from beginInternal() or if QueueLeaves fails
// below.

type logTreeTX struct {
	treeTX
	ls   *memoryLogStorage
	root types.LogRootV1
	slr  *trillian.SignedLogRoot
}

// GetMerkleNodes returns the requested nodes at (or below) the read revision.
func (t *logTreeTX) GetMerkleNodes(ctx context.Context, ids []compact.NodeID) ([]stree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) DequeueLeaves(ctx context.Context, limit int, cutoffTime time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(al): consider cutoffTime

func (t *logTreeTX) QueueLeaves(ctx context.Context, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	// Don't accept batches if any of the leaves are invalid.
	return nil, nil
}

// No deduping in this storage!

func (t *logTreeTX) AddSequencedLeaves(ctx context.Context, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) GetLeavesByRange(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) GetLeavesByHash(ctx context.Context, leafHashes [][]byte, orderBySequence bool) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) LatestSignedLogRoot(ctx context.Context) (*trillian.SignedLogRoot, error) {
	_ = "STUB: not implemented"

	// fetchLatestRoot reads the latest SignedLogRoot from the DB and returns it.
	return nil, nil
}

func (t *logTreeTX) fetchLatestRoot(ctx context.Context) (*trillian.SignedLogRoot, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (t *logTreeTX) StoreSignedLogRoot(ctx context.Context, slr *trillian.SignedLogRoot) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(alcutter): this breaks the transactional model

func (t *logTreeTX) UpdateSequencedLeaves(ctx context.Context, leaves []*trillian.LogLeaf) error {
	_ = "STUB: not implemented"
	return nil
}

// This should fail on insert but catch it early

// insert sequenced leaf:

// update merkle-to-seq mapping:
