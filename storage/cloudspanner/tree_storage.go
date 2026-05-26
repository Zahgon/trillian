// Copyright 2018 Google LLC. All Rights Reserved.
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

package cloudspanner

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/storage/cache"
	"github.com/google/trillian/storage/cloudspanner/spannerpb"
	"github.com/google/trillian/storage/storagepb"
	"github.com/google/trillian/storage/tree"
	"github.com/transparency-dev/merkle/compact"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	// ErrNotFound is returned when a read/lookup fails because there was no such
	// item.
	ErrNotFound = status.Errorf(codes.NotFound, "not found")

	// ErrNotImplemented is returned by any interface methods which have not been
	// implemented yet.
	ErrNotImplemented = errors.New("not implemented")

	// ErrTransactionClosed is returned by interface methods when an operation is
	// attempted on a transaction whose Commit or Close methods have
	// previously been called.
	ErrTransactionClosed = errors.New("transaction is closed")

	// ErrWrongTXType is returned when, somehow, a write operation is attempted
	// with a read-only transaction.  This should not even be possible.
	ErrWrongTXType = errors.New("mutating method called on read-only transaction")
)

const (
	subtreeTbl   = "SubtreeData"
	colSubtree   = "Subtree"
	colSubtreeID = "SubtreeID"
	colTreeID    = "TreeID"
	colRevision  = "Revision"
)

// treeStorage provides a shared base for the concrete CloudSpanner-backed
// implementation of the Trillian storage.LogStorage and storage.MapStorage
// interfaces.
type treeStorage struct {
	admin  storage.AdminStorage
	opts   TreeStorageOptions
	client *spanner.Client
}

// TreeStorageOptions holds various levers for configuring the tree storage instance.
type TreeStorageOptions struct {
	// ReadOnlyStaleness controls how far in the past a read-only snapshot
	// transaction will read.
	// This is intended to allow Spanner to use local replicas for read requests
	// to help with performance.
	// See https://cloud.google.com/spanner/docs/timestamp-bounds for more details.
	ReadOnlyStaleness time.Duration
}

func newTreeStorageWithOpts(client *spanner.Client, opts TreeStorageOptions) *treeStorage {
	_ = "STUB: not implemented"
	return nil
}

type spanRead interface {
	Query(context.Context, spanner.Statement) *spanner.RowIterator
	Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadUsingIndex(ctx context.Context, table, index string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
	ReadWithOptions(ctx context.Context, table string, keys spanner.KeySet, columns []string, opts *spanner.ReadOptions) (ri *spanner.RowIterator)
}

// latestSTH reads and returns the newest STH.
func (t *treeStorage) latestSTH(ctx context.Context, stx spanRead, treeID int64) (*spannerpb.TreeHead, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type newCacheFn func(*trillian.Tree) (*cache.SubtreeCache, error)

func (t *treeStorage) getTreeAndConfig(ctx context.Context, tree *trillian.Tree) (*trillian.Tree, proto.Message, error) {
	_ = "STUB: not implemented"
	return nil, *new(proto.Message), nil
}

// begin returns a newly started tree transaction for the specified tree.
func (t *treeStorage) begin(ctx context.Context, tree *trillian.Tree, newCache newCacheFn, stx spanRead) (*treeTX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getLatestRoot populates this TX with the newest tree root visible (when
// taking read-staleness into account) by this transaction.
func (t *treeTX) getLatestRoot(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// treeTX is a concrete implementation of the part of storage.LogTreeTX
// interface formerly known as storage.TreeTX.
type treeTX struct {
	treeID   int64
	treeType trillian.TreeType

	ts *treeStorage

	// mu guards the nil setting/checking of stx as part of the open checking.
	mu sync.RWMutex
	// stx is the underlying Spanner transaction in which all operations will be
	// performed.
	stx spanRead

	// config holds the StorageSettings proto acquired from the trillian.Tree.
	// Varies according to tree_type (LogStorageConfig vs MapStorageConfig).
	config proto.Message

	// currentSTH holds a copy of the latest known STH at the time the
	// transaction was started, or nil if there was no STH.
	_currentSTH    *spannerpb.TreeHead
	_currentSTHErr error

	// writeRev is the tree revision at which any writes will be made.
	_writeRev int64

	cache *cache.SubtreeCache

	getLatestRootOnce sync.Once
}

func (t *treeTX) currentSTH(ctx context.Context) (*spannerpb.TreeHead, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *treeTX) writeRev(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// storeSubtrees adds buffered writes to the in-flight transaction to store the
// passed in subtrees.
func (t *treeTX) storeSubtrees(ctx context.Context, sts []*storagepb.SubtreeProto) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *treeTX) flushSubtrees(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Commit attempts to apply all actions perfomed to the underlying Spanner
// transaction.  If this call returns an error, any values READ via this
// transaction MUST NOT be used.
// On return from the call, this transaction will be in a closed state.
func (t *treeTX) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close aborts any operations perfomed on the underlying Spanner transaction.
// On return from the call, this transaction will be in a closed state.
func (t *treeTX) Close() error { _ = "STUB: not implemented"; return nil }

// readRevision returns the tree revision at which the currently visible (taking
// into account read-staleness) STH was stored.
func (t *treeTX) readRevision(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// getSubtree retrieves the most recent subtree specified by id at (or below)
// the requested revision.
// If no such subtree exists it returns nil.
func (t *treeTX) getSubtree(ctx context.Context, rev int64, id []byte) (p *storagepb.SubtreeProto, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If this is a subtree with a zero-length prefix, we'll need to create an
// empty Prefix field:

// GetMerkleNodes returns the requested set of nodes at, or before, the
// transaction read revision.
func (t *treeTX) GetMerkleNodes(ctx context.Context, ids []compact.NodeID) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSubtreesAtRev returns a GetSubtreesFunc which reads at the passed in rev.
func (t *treeTX) getSubtreesAtRev(ctx context.Context, rev int64) cache.GetSubtreesFunc {
	_ = "STUB: not implemented"
	return *new(cache.GetSubtreesFunc)
}

// Request the various subtrees in parallel.
// c will carry any retrieved subtrees

// Spawn goroutines for each request

// Now wait for the goroutines to signal their completion, and collect
// the results.

// SetMerkleNodes stores the provided merkle nodes at the writeRevision of the
// transaction.
func (t *treeTX) SetMerkleNodes(ctx context.Context, nodes []tree.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func checkDatabaseAccessible(ctx context.Context, client *spanner.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't care about freshness here, being able to read *something* is enough
