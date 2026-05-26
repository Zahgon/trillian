// Copyright 2024 Trillian Authors. All Rights Reserved.
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

package postgresql

import (
	"context"
	"sync"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/storage/tree"
	"github.com/google/trillian/types"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/transparency-dev/merkle/compact"
)

const (
	queueLeafSQL = "WITH insert_leaf AS (" +
		"INSERT INTO LeafData (TreeId,LeafIdentityHash,LeafValue,ExtraData,QueueTimestampNanos) " +
		"VALUES ($1,$2,$3,$4,$5) " +
		"ON CONFLICT DO NOTHING " +
		"RETURNING *" +
		") " +
		"INSERT INTO Unsequenced (TreeId,Bucket,LeafIdentityHash,MerkleLeafHash,QueueTimestampNanos,QueueID) " +
		"SELECT TreeId,0,LeafIdentityHash,$6,QueueTimestampNanos,$7 " +
		"FROM insert_leaf"
	createTempQueueLeavesTable = "CREATE TEMP TABLE TempQueueLeaves (" +
		" TreeId BIGINT," +
		" LeafIdentityHash BYTEA," +
		" LeafValue BYTEA," +
		" ExtraData BYTEA," +
		" MerkleLeafHash BYTEA," +
		" QueueTimestampNanos BIGINT," +
		" QueueID BYTEA," +
		" IsDuplicate BOOLEAN DEFAULT FALSE" +
		") ON COMMIT DROP"
	queueLeavesSQL = "SELECT * FROM queue_leaves()"

	createTempAddSequencedLeavesTable = "CREATE TEMP TABLE TempAddSequencedLeaves (" +
		" TreeId BIGINT," +
		" LeafIdentityHash BYTEA," +
		" LeafValue BYTEA," +
		" ExtraData BYTEA," +
		" MerkleLeafHash BYTEA," +
		" QueueTimestampNanos BIGINT," +
		" SequenceNumber BIGINT," +
		" IsDuplicateLeafData BOOLEAN DEFAULT FALSE," +
		" IsDuplicateSequencedLeafData BOOLEAN DEFAULT FALSE" +
		") ON COMMIT DROP"
	addSequencedLeavesSQL = "SELECT * FROM add_sequenced_leaves()"

	selectNonDeletedTreeIDByTypeAndStateSQL = "SELECT TreeId " +
		"FROM Trees " +
		"WHERE TreeType IN($1,$2)" +
		" AND TreeState IN($3,$4)" +
		" AND (Deleted IS NULL OR Deleted='false')"

	selectLatestSignedLogRootSQL = "SELECT TreeHeadTimestamp,TreeSize,RootHash,RootSignature " +
		"FROM TreeHead " +
		"WHERE TreeId=$1 " +
		"ORDER BY TreeHeadTimestamp DESC " +
		"LIMIT 1"

	selectLeavesByRangeSQL = "SELECT s.MerkleLeafHash,l.LeafIdentityHash,l.LeafValue,s.SequenceNumber,l.ExtraData,l.QueueTimestampNanos,s.IntegrateTimestampNanos " +
		"FROM SequencedLeafData s" +
		" INNER JOIN LeafData l ON (s.LeafIdentityHash=l.LeafIdentityHash AND s.TreeId=l.TreeId) " +
		"WHERE s.SequenceNumber>=$1" +
		" AND s.SequenceNumber<$2" +
		" AND l.TreeId=$3" + orderBySequenceNumberSQL

	selectLeavesByMerkleHashSQL = "SELECT s.MerkleLeafHash,l.LeafIdentityHash,l.LeafValue,s.SequenceNumber,l.ExtraData,l.QueueTimestampNanos,s.IntegrateTimestampNanos " +
		"FROM SequencedLeafData s" +
		" INNER JOIN LeafData l ON (s.LeafIdentityHash=l.LeafIdentityHash AND s.TreeId=l.TreeId) " +
		"WHERE s.MerkleLeafHash=ANY($1)" +
		" AND l.TreeId=$2"
	// TODO(robstradling): Per #1548, rework the code so the dummy hash isn't needed (e.g. this assumes hash size is 32)
	dummyMerkleLeafHash = "00000000000000000000000000000000"
	// This statement returns a dummy Merkle leaf hash value (which must be
	// of the right size) so that its signature matches that of the other
	// leaf-selection statements.
	selectLeavesByLeafIdentityHashSQL = "SELECT decode('" + dummyMerkleLeafHash + "','escape'),l.LeafIdentityHash,l.LeafValue,-1,l.ExtraData,l.QueueTimestampNanos,s.IntegrateTimestampNanos " +
		"FROM LeafData l" +
		" LEFT JOIN SequencedLeafData s ON (l.LeafIdentityHash=s.LeafIdentityHash AND l.TreeId=s.TreeId) " +
		"WHERE l.LeafIdentityHash=ANY($1)" +
		" AND l.TreeId=$2"

	// Same as above except with leaves ordered by sequence so we only incur this cost when necessary
	orderBySequenceNumberSQL                     = " ORDER BY s.SequenceNumber"
	selectLeavesByMerkleHashOrderedBySequenceSQL = selectLeavesByMerkleHashSQL + orderBySequenceNumberSQL

	logIDLabel = "logid"
)

var (
	once             sync.Once
	queuedCounter    monitoring.Counter
	queuedDupCounter monitoring.Counter
	dequeuedCounter  monitoring.Counter

	dequeueLatency       monitoring.Histogram
	dequeueSelectLatency monitoring.Histogram
	dequeueRemoveLatency monitoring.Histogram
)

func createMetrics(mf monitoring.MetricFactory) { _ = "STUB: not implemented"; return }

func labelForTX(t *logTreeTX) string { _ = "STUB: not implemented"; return "" }

func observe(hist monitoring.Histogram, duration time.Duration, label string) {
	_ = "STUB: not implemented"
	return
}

type postgreSQLLogStorage struct {
	*postgreSQLTreeStorage
	admin         storage.AdminStorage
	metricFactory monitoring.MetricFactory
}

// NewLogStorage creates a storage.LogStorage instance for the specified PostgreSQL URL.
// It assumes storage.AdminStorage is backed by the same PostgreSQL database as well.
func NewLogStorage(db *pgxpool.Pool, mf monitoring.MetricFactory) storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

func (m *postgreSQLLogStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *postgreSQLLogStorage) GetActiveLogIDs(ctx context.Context) ([]int64, error) {
	_ = "STUB: not implemented"
	// Include logs that are DRAINING in the active list as we're still
	// integrating leaves into them.
	return nil, nil
}

func (m *postgreSQLLogStorage) beginInternal(ctx context.Context, tree *trillian.Tree) (*logTreeTX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(robstradling): This and many other methods of this storage
// implementation can leak a specific sql.ErrTxDone all the way to the client,
// if the transaction is rolled back as a result of a canceled context. It must
// return "generic" errors, and only log the specific ones for debugging.
func (m *postgreSQLLogStorage) ReadWriteTransaction(ctx context.Context, tree *trillian.Tree, f storage.LogTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *postgreSQLLogStorage) AddSequencedLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure we don't leak the transaction. For example if we get an
// ErrTreeNeedsInit from beginInternal() or if AddSequencedLeaves fails
// below.

func (m *postgreSQLLogStorage) SnapshotForTree(ctx context.Context, tree *trillian.Tree) (storage.ReadOnlyLogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyLogTreeTX), nil
}

func (m *postgreSQLLogStorage) QueueLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure we don't leak the transaction. For example if we get an
// ErrTreeNeedsInit from beginInternal() or if QueueLeaves fails
// below.

// Queue leave(s), using a more efficient implementation when the batch size is 1.

type logTreeTX struct {
	treeTX
	ls       *postgreSQLLogStorage
	root     types.LogRootV1
	slr      *trillian.SignedLogRoot
	dequeued map[string]dequeuedLeaf
}

// GetMerkleNodes returns the requested nodes.
func (t *logTreeTX) GetMerkleNodes(ctx context.Context, ids []compact.NodeID) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) DequeueLeaves(ctx context.Context, limit int, cutoffTime time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(robstradling): Optimize this by fetching only the required
// fields of LogLeaf. We can avoid joining with LeafData table here.

// dupe, user probably called DequeueLeaves more than once.

func (t *logTreeTX) QueueLeaf(ctx context.Context, leaf *trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare details to store, but don't accept leaf if invalid.

// Create the leaf data record and work queue entry, unless the leaf already exists.

// Leaf did not already exist.

// Replace the requested leaf with the actual leaf.

func (t *logTreeTX) QueueLeaves(ctx context.Context, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare rows to copy, but don't accept batches if any of the leaves are invalid.

// Create temporary table.

// Copy rows to temporary table.

// Create the leaf data records, work queue entries, and obtain a deduplicated list of existing leaves.

// Remember the duplicate leaf, using the requested leaf for now.

// Replace the requested leaves with the actual leaves.

func (t *logTreeTX) AddSequencedLeaves(ctx context.Context, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare rows to copy.

// This should fail on insert, but catch it early.

// Create temporary table.

// Copy rows to temporary table.

// Create the leaf data records and sequenced leaf data records, returning details of which records already existed.

// TODO(robstradling): Support opting out from duplicates detection.
// TODO(robstradling): Update IntegrateTimestamp on integrating the leaf.
// TODO(robstradling): Load LeafData for conflicting entries.

func (t *logTreeTX) GetLeavesByRange(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) getLeavesByRangeInternal(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure no entries queried/returned beyond the tree.

// TODO(robstradling): Further clip `count` to a safe upper bound like 64k.

func (t *logTreeTX) GetLeavesByHash(ctx context.Context, leafHashes [][]byte, orderBySequence bool) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getLeafDataByIdentityHash retrieves leaf data by LeafIdentityHash, returned
// as a slice of LogLeaf objects for convenience.  However, note that the
// returned LogLeaf objects will not have a valid MerkleLeafHash, LeafIndex, or IntegrateTimestamp.
func (t *logTreeTX) getLeafDataByIdentityHash(ctx context.Context, leafHashes [][]byte) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) LatestSignedLogRoot(ctx context.Context) (*trillian.SignedLogRoot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetchLatestRoot reads the latest root from the DB.
func (t *logTreeTX) fetchLatestRoot(ctx context.Context) (*trillian.SignedLogRoot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// It's possible there are no roots for this tree yet

// Put logRoot back together. Fortunately LogRoot has a deterministic serialization.

func (t *logTreeTX) StoreSignedLogRoot(ctx context.Context, root *trillian.SignedLogRoot) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *logTreeTX) getLeavesByHashInternal(ctx context.Context, leafHashes [][]byte, query string, desc string) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The tree could include duplicates so we don't know how many results will be returned

// We might be using a LEFT JOIN in our statement, so leaves which are
// queued but not yet integrated will have a NULL IntegrateTimestamp
// when there's no corresponding entry in SequencedLeafData, even though
// the table definition forbids that, so we use a nullable type here and
// check its validity below.
