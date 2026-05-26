// Copyright 2016 Trillian Authors
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

package crdb

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/transparency-dev/merkle/compact"

	"github.com/google/trillian"
	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/storage/tree"
	"github.com/google/trillian/types"
)

const (
	valuesPlaceholder5 = "($1,$2,$3,$4,$5)"

	insertLeafDataSQL      = "INSERT INTO LeafData(TreeId,LeafIdentityHash,LeafValue,ExtraData,QueueTimestampNanos) VALUES" + valuesPlaceholder5
	insertSequencedLeafSQL = "INSERT INTO SequencedLeafData(TreeId,LeafIdentityHash,MerkleLeafHash,SequenceNumber,IntegrateTimestampNanos) VALUES"

	selectNonDeletedTreeIDByTypeAndStateSQL = `
		SELECT TreeId FROM Trees
		  WHERE TreeType IN($1,$2)
		  AND TreeState IN($3,$4)
		  AND (Deleted IS NULL OR Deleted = 'false')`

	selectLatestSignedLogRootSQL = `SELECT TreeHeadTimestamp,TreeSize,RootHash,TreeRevision,RootSignature
			FROM TreeHead WHERE TreeId=$1
			ORDER BY TreeHeadTimestamp DESC LIMIT 1`

	selectLeavesByRangeSQL = `SELECT s.MerkleLeafHash,l.LeafIdentityHash,l.LeafValue,s.SequenceNumber,l.ExtraData,l.QueueTimestampNanos,s.IntegrateTimestampNanos
			FROM LeafData l,SequencedLeafData s
			WHERE l.LeafIdentityHash = s.LeafIdentityHash
			AND s.SequenceNumber >= $1 AND s.SequenceNumber < $2 AND l.TreeId = $3 AND s.TreeId = l.TreeId` + orderBySequenceNumberSQL

	// These statements need to be expanded to provide the correct number of parameter placeholders.
	// Note that this uses the MySQL-specific marker syntax here, but is eventually replaced with
	// the postgres syntax in getStmt.
	selectLeavesByMerkleHashSQL = `SELECT s.MerkleLeafHash,l.LeafIdentityHash,l.LeafValue,s.SequenceNumber,l.ExtraData,l.QueueTimestampNanos,s.IntegrateTimestampNanos
			FROM LeafData l,SequencedLeafData s
			WHERE l.LeafIdentityHash = s.LeafIdentityHash
			AND s.MerkleLeafHash IN (` + placeholderSQL + `) AND l.TreeId = ? AND s.TreeId = l.TreeId`
	// TODO(#1548): rework the code so the dummy hash isn't needed (e.g. this assumes hash size is 32)
	dummyMerkleLeafHash = "00000000000000000000000000000000"
	// This statement returns a dummy Merkle leaf hash value (which must be
	// of the right size) so that its signature matches that of the other
	// leaf-selection statements.
	// Note that this uses the MySQL-specific marker syntax here, but is eventually replaced with
	// the postgres syntax in getStmt.
	selectLeavesByLeafIdentityHashSQL = `SELECT '` + dummyMerkleLeafHash + `',l.LeafIdentityHash,l.LeafValue,-1,l.ExtraData,l.QueueTimestampNanos,s.IntegrateTimestampNanos
			FROM LeafData l LEFT JOIN SequencedLeafData s ON (l.LeafIdentityHash = s.LeafIdentityHash AND l.TreeID = s.TreeID)
			WHERE l.LeafIdentityHash IN (` + placeholderSQL + `) AND l.TreeId = ?`

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

	queueLatency            monitoring.Histogram
	queueInsertLatency      monitoring.Histogram
	queueReadLatency        monitoring.Histogram
	queueInsertLeafLatency  monitoring.Histogram
	queueInsertEntryLatency monitoring.Histogram
	dequeueLatency          monitoring.Histogram
	dequeueSelectLatency    monitoring.Histogram
	dequeueRemoveLatency    monitoring.Histogram
)

func createMetrics(mf monitoring.MetricFactory) { _ = "STUB: not implemented"; return }

func labelForTX(t *logTreeTX) string { _ = "STUB: not implemented"; return "" }

func observe(hist monitoring.Histogram, duration time.Duration, label string) {
	_ = "STUB: not implemented"
	return
}

type crdbLogStorage struct {
	*crdbTreeStorage
	admin         storage.AdminStorage
	metricFactory monitoring.MetricFactory
}

// NewLogStorage creates a storage.LogStorage instance for the specified CockroachDB URL.
// It assumes storage.AdminStorage is backed by the same CockroachDB database as well.
func NewLogStorage(db *sql.DB, mf monitoring.MetricFactory) storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

func (m *crdbLogStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *crdbLogStorage) getLeavesByMerkleHashStmt(ctx context.Context, num int, orderBySequence bool) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *crdbLogStorage) getLeavesByLeafIdentityHashStmt(ctx context.Context, num int) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *crdbLogStorage) GetActiveLogIDs(ctx context.Context) ([]int64, error) {
	_ = "STUB: not implemented"
	// Include logs that are DRAINING in the active list as we're still
	// integrating leaves into them.
	return nil, nil
}

func (m *crdbLogStorage) beginInternal(ctx context.Context, tree *trillian.Tree) (*logTreeTX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(pavelkalinnikov): This and many other methods of this storage
// implementation can leak a specific sql.ErrTxDone all the way to the client,
// if the transaction is rolled back as a result of a canceled context. It must
// return "generic" errors, and only log the specific ones for debugging.
func (m *crdbLogStorage) ReadWriteTransaction(ctx context.Context, tree *trillian.Tree, f storage.LogTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *crdbLogStorage) AddSequencedLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure we don't leak the transaction. For example if we get an
// ErrTreeNeedsInit from beginInternal() or if AddSequencedLeaves fails
// below.

func (m *crdbLogStorage) SnapshotForTree(ctx context.Context, tree *trillian.Tree) (storage.ReadOnlyLogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyLogTreeTX), nil
}

func (m *crdbLogStorage) QueueLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure we don't leak the transaction. For example if we get an
// ErrTreeNeedsInit from beginInternal() or if QueueLeaves fails
// below.

type logTreeTX struct {
	treeTX
	ls       *crdbLogStorage
	root     types.LogRootV1
	readRev  int64
	slr      *trillian.SignedLogRoot
	dequeued map[string]dequeuedLeaf
}

// GetMerkleNodes returns the requested nodes at the read revision.
func (t *logTreeTX) GetMerkleNodes(ctx context.Context, ids []compact.NodeID) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) DequeueLeaves(ctx context.Context, limit int, cutoffTime time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(pavelkalinnikov): Optimize this by fetching only the required
// fields of LogLeaf. We can avoid joining with LeafData table here.

// dupe, user probably called DequeueLeaves more than once.

// sortLeavesForInsert returns a slice containing the passed in leaves sorted
// by LeafIdentityHash, and paired with their original positions.
// QueueLeaves and AddSequencedLeaves use this to make the order that LeafData
// row locks are acquired deterministic and reduce the chance of deadlocks.
func sortLeavesForInsert(leaves []*trillian.LogLeaf) []leafAndPosition {
	_ = "STUB: not implemented"
	return nil
}

func (t *logTreeTX) QueueLeaves(ctx context.Context, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't accept batches if any of the leaves are invalid.

// CockroachDB/Postgres will cancel a transaction if an insert
// statement is run with a duplicate key. This is not ideal for
// QueueLeaves, as we want to detect these errors in-code
// and return them as AlreadyExists errors and add metrics.
// Thus, we use a SAVEPOINT and rollback on duplicates.

// Remember the duplicate leaf, using the requested leaf for now.

// Note: one must roll back since there are side-effects in the transaction
// in crdb/postgres

// Create the work queue entry

// For existing leaves, we need to retrieve the contents.  First collate the desired LeafIdentityHash values.

// Replace the requested leaves with the actual leaves.

func (t *logTreeTX) AddSequencedLeaves(ctx context.Context, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Leaves in this transaction are inserted in two tables. For each leaf, if
// one of the two inserts fails, we remove the side effect by rolling back to
// a savepoint installed before the first insert of the two.

// TODO(pavelkalinnikov): Consider performance implication of executing this
// extra SAVEPOINT, especially for 1-entry batches. Optimize if necessary.

// Note: LeafData inserts are presumably protected from deadlocks due to
// sorting, but the order of the corresponding SequencedLeafData inserts
// becomes indeterministic. However, in a typical case when leaves are
// supplied in contiguous non-intersecting batches, the chance of having
// circular dependencies between transactions is significantly lower.

// This should fail on insert, but catch it early.

// TODO(pavelkalinnikov): Measure latencies.

// TODO(pavelkalinnikov): Detach PREORDERED_LOG integration latency metric.

// TODO(pavelkalinnikov): Support opting out from duplicates detection.

// Note: one must roll back since there are side-effects in the transaction
// in crdb/postgres

// TODO(pavelkalinnikov): Update IntegrateTimestamp on integrating the leaf.

// TODO(pavelkalinnikov): Load LeafData for conflicting entries.

func (t *logTreeTX) GetLeavesByRange(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *logTreeTX) getLeavesByRangeInternal(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure no entries queried/returned beyond the tree.

// TODO(pavelkalinnikov): Further clip `count` to a safe upper bound like 64k.

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

// fetchLatestRoot reads the latest root and the revision from the DB.
func (t *logTreeTX) fetchLatestRoot(ctx context.Context) (*trillian.SignedLogRoot, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// It's possible there are no roots for this tree yet

// Put logRoot back together. Fortunately LogRoot has a deterministic serialization.

func (t *logTreeTX) StoreSignedLogRoot(ctx context.Context, root *trillian.SignedLogRoot) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *logTreeTX) getLeavesByHashInternal(ctx context.Context, leafHashes [][]byte, tmpl *sql.Stmt, desc string) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The tree could include duplicates so we don't know how many results will be returned

// We might be using a LEFT JOIN in our statement, so leaves which are
// queued but not yet integrated will have a NULL IntegrateTimestamp
// when there's no corresponding entry in SequencedLeafData, even though
// the table definition forbids that, so we use a nullable type here and
// check its validity below.

// leafAndPosition records original position before sort.
type leafAndPosition struct {
	leaf *trillian.LogLeaf
	idx  int
}

// byLeafIdentityHashWithPosition allows sorting (as above), but where we need
// to remember the original position
type byLeafIdentityHashWithPosition []leafAndPosition

func (l byLeafIdentityHashWithPosition) Len() int { _ = "STUB: not implemented"; return 0 }

func (l byLeafIdentityHashWithPosition) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (l byLeafIdentityHashWithPosition) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}
