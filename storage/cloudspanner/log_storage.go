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
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/storage/cache"
	"github.com/google/trillian/storage/cloudspanner/spannerpb"
	"golang.org/x/sync/semaphore"
)

const (
	leafDataTbl            = "LeafData"
	seqDataByMerkleHashIdx = "SequenceByMerkleHash"
	seqDataTbl             = "SequencedLeafData"
	unseqTable             = "Unsequenced"

	// t.TreeType: 1 = Log, 3 = PreorderedLog.
	// t.TreeState: 1 = Active, 5 = Draining.
	getActiveLogIDsSQL = `SELECT t.TreeID FROM TreeRoots t
WHERE (t.TreeType = 1 OR t.TreeType = 3)
AND (t.TreeState = 1 OR t.TreeState = 5)
AND t.Deleted=false`
)

// LogStorageOptions are tuning, experiments and workarounds that can be used.
type LogStorageOptions struct {
	TreeStorageOptions

	// DequeueAcrossMerkleBuckets controls whether DequeueLeaves will only dequeue
	// from within the chosen Time+Merkle bucket, or whether it will attempt to
	// continue reading from contiguous Merkle buckets until a sufficient number
	// of leaves have been dequeued, or the entire Time bucket has been read.
	DequeueAcrossMerkleBuckets bool
	// DequeueAcrossMerkleBucketsRangeFraction specifies the fraction of Merkle
	// keyspace to dequeue from when using multi-bucket-dequeue.
	DequeueAcrossMerkleBucketsRangeFraction float64
}

var (
	// Spanner DB columns:
	colLeafIdentityHash    = "LeafIdentityHash"
	colLeafValue           = "LeafValue"
	colExtraData           = "ExtraData"
	colMerkleLeafHash      = "MerkleLeafHash"
	colSequenceNumber      = "SequenceNumber"
	colQueueTimestampNanos = "QueueTimestampNanos"
)

type leafDataCols struct {
	TreeID              int64
	LeafIdentityHash    []byte
	LeafValue           []byte
	ExtraData           []byte
	QueueTimestampNanos int64
}

type sequencedLeafDataCols struct {
	TreeID                  int64
	SequenceNumber          int64
	LeafIdentityHash        []byte
	MerkleLeafHash          []byte
	IntegrateTimestampNanos int64
}

type unsequencedCols struct {
	TreeID              int64
	Bucket              int64
	QueueTimestampNanos int64
	MerkleLeafHash      []byte
	LeafIdentityHash    []byte
}

// NewLogStorage initialises and returns a new LogStorage.
func NewLogStorage(client *spanner.Client) storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

// NewLogStorageWithOpts initialises and returns a new LogStorage.
// The opts parameter can be used to enable custom workarounds.
func NewLogStorageWithOpts(client *spanner.Client, opts LogStorageOptions) storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

// This number is taken from the maximum number of in-flight
// transaction in the mutation pool. Add a field to opts if we decide to
// adopt this strategy.

// logStorage provides a Cloud Spanner backed trillian.LogStorage implementation.
// See third_party/golang/trillian/storage/log_storage.go for more details.
type logStorage struct {
	// ts provides the merkle-tree level primitives which are built upon by this
	// logStorage.
	ts *treeStorage

	// writeSem controls how many concurrent writes QueueLeaves/AddSequencedLeaves will do.
	writeSem *semaphore.Weighted

	// Additional options applied to this logStorage
	opts LogStorageOptions
}

func (ls *logStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ls *logStorage) readOnlyTX() *spanner.ReadOnlyTransaction {
	_ = "STUB: not implemented"
	return nil
}

func (ls *logStorage) GetActiveLogIDs(ctx context.Context) ([]int64, error) {
	_ = "STUB: not implemented"

	// We have to use SQL as Read() doesn't work against an index.
	return nil, nil
}

func newLogCache(tree *trillian.Tree) (*cache.SubtreeCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *logStorage) begin(ctx context.Context, tree *trillian.Tree, readonly bool, stx spanRead) (*logTX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Needed to generate ErrTreeNeedsInit in SnapshotForTree and other methods.

func (ls *logStorage) BeginForTree(ctx context.Context, treeID int64) (storage.LogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.LogTreeTX), nil
}

func (ls *logStorage) ReadWriteTransaction(ctx context.Context, tree *trillian.Tree, f storage.LogTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

/* readonly */

func (ls *logStorage) SnapshotForTree(ctx context.Context, tree *trillian.Tree) (storage.ReadOnlyLogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyLogTreeTX), nil
}

/* readonly */

func (ls *logStorage) QueueLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, qTimestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture values of i and l for later reference in the MutationResultFunc below.

// The insert of the leafdata and the unsequenced work item must happen atomically.

// implicit OK status

// Wait for all of our mutations to apply (or fail):

// Finally, read back any leaves which failed with an already exists error
// when we tried to insert them:

func (ls *logStorage) AddSequencedLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, ts time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture the values for later reference in the MutationResultFunc below.

// The insert of the LeafData and SequencedLeafData must happen atomically.

// If failed because of a duplicate insert, set the status correspondingly.

// Skip this error, we only need one.

// Wait for all of our mutations to apply (or fail).

// Check if any failed, and return the first error if so.

// No error.

// readDupeLeaves reads the leaves whose ids are passed as keys in the dupes map,
// and stores them in results.
func (ls *logStorage) readDupeLeaves(ctx context.Context, logID int64, dupes map[string][]int, results []*trillian.QueuedLogLeaf) error {
	_ = "STUB: not implemented"
	return nil
}

// logTX is a concrete implementation of the Trillian storage.LogStorage
// interface.
type logTX struct {
	// treeTX embeds the merkle-tree level transactional actions.
	*treeTX

	// logStorage is the logStorage which begat this logTX.
	ls *logStorage

	// numSequenced holds the number of leaves sequenced by this transaction.
	numSequenced int64

	// dequeued is a map of LeafIdentityHash to QueuedEntry containing entries for
	// everything dequeued by this transaction.
	// This is required to recover the primary key for the unsequenced entry in
	// UpdateSequencedLeaves.
	dequeued map[string]*QueuedEntry
}

func (tx *logTX) getLogStorageConfig() *spannerpb.LogStorageConfig {
	_ = "STUB: not implemented"
	return nil
}

// LatestSignedLogRoot returns the freshest SignedLogRoot for this log at the
// time the transaction was started.
func (tx *logTX) LatestSignedLogRoot(ctx context.Context) (*trillian.SignedLogRoot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put logRoot back together. Fortunately LogRoot has a deterministic serialization.

// We already read the latest root as part of starting the transaction (in
// order to calculate the writeRevision), so we just return that data here:

// StoreSignedLogRoot stores the provided root.
// This method will return an error if the caller attempts to store more than
// one root per log for a given tree size.
func (tx *logTX) StoreSignedLogRoot(ctx context.Context, root *trillian.SignedLogRoot) error {
	_ = "STUB: not implemented"
	return nil
}

func readLeaves(ctx context.Context, stx *spanner.ReadOnlyTransaction, logID int64, ids [][]byte, f func(*trillian.LogLeaf)) error {
	_ = "STUB: not implemented"
	return nil
}

// DequeueLeaves removes [0, limit) leaves from the to-be-sequenced queue.
// The leaves returned are not guaranteed to be in any particular order.
// The caller should assign sequence numbers and pass the updated leaves as
// arguments to the UpdateSequencedLeaves method.
//
// The LogLeaf structs returned by this method will not be fully populated;
// only the LeafIdentityHash and MerkleLeafHash fields will contain data, this
// should be sufficient for assigning sequence numbers with this storage impl.
//
// TODO(al): cutoff is currently ignored.
func (tx *logTX) DequeueLeaves(ctx context.Context, limit int, cutoff time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Special case pre-ordered logs.

// Decide which bucket(s) to dequeue from.
// The high 8 bits of the bucket key is a time based ring - at any given
// moment, FEs queueing entries will be adding them to different buckets
// than we're dequeuing from here - the low 8 bits are the first byte of the
// merkle hash of the entry.

// Select a prefix that is likley to be on a different span server to spread load.

// Choose a starting point in the merkle prefix range, and calculate the
// start/limit of the merkle range we'll dequeue from.
// It seems to be much better to tune for keeping this range small, and allow
// the signer to run multiple times per second than try to dequeue a large batch
// which spans a large number of merkle prefixes.

// The range is too big and wraps around, overflowing a byte value, so we'll
// start the second range at 0 and end at the upper limit modulo suffixBuckets:

// XXX: When suffixFraction = 1, this produces an overlapping range at suffixStart

// dupe, user probably called DequeueLeaves more than once.

// If we've already got enough leaves, don't wrap around for any further reads.

// UpdateSequencedLeaves stores the sequence numbers assigned to the leaves,
// and integrates them into the tree.
func (tx *logTX) UpdateSequencedLeaves(ctx context.Context, leaves []*trillian.LogLeaf) error {
	_ = "STUB: not implemented"
	return nil
}

// We need the latest root to know what the next sequence number to use below is.

// Add the sequence mapping...

// leafmap is a map of LogLeaf by sequence number which knows how to populate
// itself directly from Spanner Rows.
type leafmap map[int64]*trillian.LogLeaf

// addFullRow appends the leaf data in row to the array
func (l leafmap) addFullRow(seqLeaves map[string]sequencedLeafDataCols) func(r *spanner.Row) error {
	_ = "STUB: not implemented"
	return nil
}

// leavesByHash is a map of []LogLeaf (keyed by value hash) which knows how to
// populate itself from Spanner Rows.
type leavesByHash map[string][]*trillian.LogLeaf

// addRow adds the contents of the Spanner Row to this map.
func (b leavesByHash) addRow(r *spanner.Row) error { _ = "STUB: not implemented"; return nil }

// populateLeafData populates the partial LogLeaf structs held in the passed in
// map of LeafIdentityHash to []LogLeaf by reading the remaining LogLeaf data from
// Spanner.
// The value of byHash is an []LogLeaf because the underlying leaf data could
// be sequenced into multiple tree leaves if the log allows duplication.
func (tx *logTX) populateLeafData(ctx context.Context, byHash leavesByHash) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRange(start, count, treeSize int64) error { _ = "STUB: not implemented"; return nil }

// GetLeavesByRange returns the leaves corresponding to the given index range.
func (tx *logTX) GetLeavesByRange(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	// We need the latest root to validate the indices are within range.
	return nil, nil
}

// Allow requesting entries beyond the tree size.

// TODO: replace with INNER JOIN when spannertest supports JOINs
// https://github.com/googleapis/google-cloud-go/tree/master/spanner/spannertest

// Results need to be returned in order [start, end), all of which
// should be available (as we restricted xend/count to TreeSize).

// leafSlice is a slice of LogLeaf which knows how to populate itself from
// Spanner Rows.
type leafSlice []*trillian.LogLeaf

// addRow appends the leaf data in Row to the array.
func (l *leafSlice) addRow(r *spanner.Row) error { _ = "STUB: not implemented"; return nil }

// getUsingIndex returns a slice containing the LogLeaf structs corresponding
// to the requested keys.
// The entries in key are used in constructing a primary key (treeID, keyElem)
// for the specified Spanner index.
// If bySeq is true, the returned slice will be order by LogLeaf.LeafIndex.
func (tx *logTX) getUsingIndex(ctx context.Context, idx string, keys [][]byte, bySeq bool) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Now we can fetch & combine the actual leaf data:

// GetLeavesByHash returns the leaves corresponding to the given merkle hashes.
// Any unknown hashes will simply be ignored, and the caller should inspect the
// returned leaves to determine whether this has occurred.
// TODO(al): Currently, this method does not populate the IntegrateTimestamp
//
//	member of the returned leaves. We should convert this method to use SQL
//	rather than denormalising IntegrateTimestampNanos into the index too.
func (tx *logTX) GetLeavesByHash(ctx context.Context, hashes [][]byte, bySeq bool) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueuedEntry represents a leaf which was dequeued.
// It's used to store some extra info which is necessary for rebuilding the
// leaf's primary key when it's passed back in to UpdateSequencedLeaves.
type QueuedEntry struct {
	// leaf is partially populated with the Merkle and LeafValue hashes only.
	leaf      *trillian.LogLeaf
	bucket    int64
	timestamp int64
}

// LogLeaf sorting boilerplate below.

type byIndex []*trillian.LogLeaf

func (b byIndex) Len() int { _ = "STUB: not implemented"; return 0 }

func (b byIndex) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (b byIndex) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
