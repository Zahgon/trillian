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

// Package postgresql provides a PostgreSQL-based storage layer implementation.
package postgresql

import (
	"context"
	"sync"

	"github.com/google/trillian"
	"github.com/google/trillian/storage/cache"
	"github.com/google/trillian/storage/storagepb"
	"github.com/google/trillian/storage/tree"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// These statements are fixed
const (
	createTempSubtreeTable = "CREATE TEMP TABLE TempSubtree (" +
		" TreeId BIGINT," +
		" SubtreeId BYTEA," +
		" Nodes BYTEA," +
		" CONSTRAINT TempSubtree_pk PRIMARY KEY (TreeId,SubtreeId)" +
		") ON COMMIT DROP"
	insertSubtreeMultiSQL = "INSERT INTO Subtree(TreeId,SubtreeId,Nodes) " +
		"SELECT TreeId,SubtreeId,Nodes " +
		"FROM TempSubtree " +
		"ON CONFLICT ON CONSTRAINT Subtree_pk DO UPDATE SET Nodes=EXCLUDED.Nodes"
	insertTreeHeadSQL = "INSERT INTO TreeHead(TreeId,TreeHeadTimestamp,TreeSize,RootHash,RootSignature) " +
		"VALUES($1,$2,$3,$4,$5) " +
		"ON CONFLICT DO NOTHING"

	selectSubtreeSQL = "SELECT SubtreeId,Nodes " +
		"FROM Subtree " +
		"WHERE TreeId=$1" +
		" AND SubtreeId=ANY($2)"
)

// postgreSQLTreeStorage is shared between the postgreSQLLog- and (forthcoming) postgreSQLMap-
// Storage implementations, and contains functionality which is common to both,
type postgreSQLTreeStorage struct {
	db *pgxpool.Pool

	// pgx automatically prepares and caches statements, so there is no need for
	// a statement map in this struct.
	// (See https://github.com/jackc/pgx/wiki/Automatic-Prepared-Statement-Caching)
}

// OpenDB opens a database connection pool for all PostgreSQL-based storage implementations.
func OpenDB(dbURL string) (*pgxpool.Pool, error) { _ = "STUB: not implemented"; return nil, nil }

// Don't log uri as it could contain credentials

func newTreeStorage(db *pgxpool.Pool) *postgreSQLTreeStorage { _ = "STUB: not implemented"; return nil }

func (m *postgreSQLTreeStorage) beginTreeTx(ctx context.Context, tree *trillian.Tree, hashSizeBytes int, subtreeCache *cache.SubtreeCache) (treeTX, error) {
	_ = "STUB: not implemented"
	return *new(treeTX), nil
}

type treeTX struct {
	// mu ensures that tx can only be used for one query/exec at a time.
	mu            *sync.Mutex
	closed        bool
	tx            pgx.Tx
	ts            *postgreSQLTreeStorage
	treeID        int64
	treeType      trillian.TreeType
	hashSizeBytes int
	subtreeCache  *cache.SubtreeCache
}

func (t *treeTX) getSubtrees(ctx context.Context, ids [][]byte) ([]*storagepb.SubtreeProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The InternalNodes cache is possibly nil here, but the SubtreeCache (which called
// this method) will re-populate it.

func (t *treeTX) storeSubtrees(ctx context.Context, subtrees []*storagepb.SubtreeProto) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(robstradling): probably need to be able to batch this in the case where we have
// a really large number of subtrees to store.

// Create temporary subtree table.

// Copy subtrees to temporary table.

// Upsert the subtrees.

func checkResultOkAndRowCountIs(res pgconn.CommandTag, err error, count int64) error {
	_ = "STUB: not implemented"
	// The Exec() might have just failed
	return nil
}

// Otherwise we have to look at the result of the operation

func checkResultOkAndCopyCountIs(rowsAffected int64, err error, count int64) error {
	_ = "STUB: not implemented"
	// The Exec() might have just failed
	return nil
}

// Otherwise we have to look at the result of the operation

// getSubtreesFunc returns a GetSubtreesFunc which reads at the passed in rev.
func (t *treeTX) getSubtreesFunc(ctx context.Context) cache.GetSubtreesFunc {
	_ = "STUB: not implemented"
	return *new(cache.GetSubtreesFunc)
}

func (t *treeTX) SetMerkleNodes(ctx context.Context, nodes []tree.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *treeTX) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *treeTX) rollbackInternal() error { _ = "STUB: not implemented"; return nil }

func (t *treeTX) Close() error { _ = "STUB: not implemented"; return nil }
