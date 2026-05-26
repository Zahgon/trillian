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

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultSequenceIntervalSeconds = 60

	selectTrees = "SELECT TreeId,TreeState,TreeType,DisplayName,Description,CreateTimeMillis,UpdateTimeMillis,MaxRootDurationMillis,Deleted,DeleteTimeMillis " +
		"FROM Trees"
	selectNonDeletedTrees = selectTrees + " WHERE (Deleted IS NULL OR Deleted='false')"
	selectTreeByID        = selectTrees + " WHERE TreeId=$1"

	updateTreeSQL = "UPDATE Trees " +
		"SET TreeState=$1,TreeType=$2,DisplayName=$3,Description=$4,UpdateTimeMillis=$5,MaxRootDurationMillis=$6 " +
		"WHERE TreeId=$7"
)

// NewAdminStorage returns a PostgreSQL storage.AdminStorage implementation backed by DB.
func NewAdminStorage(db *pgxpool.Pool) *postgresqlAdminStorage {
	_ = "STUB: not implemented"
	return nil
}

// postgresqlAdminStorage implements storage.AdminStorage
type postgresqlAdminStorage struct {
	db *pgxpool.Pool
}

func (s *postgresqlAdminStorage) Snapshot(ctx context.Context) (storage.ReadOnlyAdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyAdminTX), nil
}

func (s *postgresqlAdminStorage) beginInternal(ctx context.Context) (storage.AdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.AdminTX), nil
}

func (s *postgresqlAdminStorage) ReadWriteTransaction(ctx context.Context, f storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *postgresqlAdminStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type adminTX struct {
	tx pgx.Tx

	// mu guards reads/writes on closed, which happen on Commit/Close methods.
	//
	// We don't check closed on methods apart from the ones above, as we trust tx
	// to keep tabs on its state, and hence fail to do queries after closed.
	mu     sync.RWMutex
	closed bool
}

func (t *adminTX) Commit() error { _ = "STUB: not implemented"; return nil }

func (t *adminTX) Close() error { _ = "STUB: not implemented"; return nil }

func (t *adminTX) GetTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	// GetTree is an entry point for most RPCs, let's provide somewhat nicer error messages.
	return nil, nil
}

// ErrNoRows doesn't provide useful information, so we don't forward it.

func (t *adminTX) ListTrees(ctx context.Context, includeDeleted bool) ([]*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) CreateTree(ctx context.Context, tree *trillian.Tree) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the time truncated-to-millis throughout, as that's what's stored.

/* SigningEnabled */
/* SequencingEnabled */

func (t *adminTX) UpdateTree(ctx context.Context, treeID int64, updateFunc func(*trillian.Tree)) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(robstradling): When switching TreeType from PREORDERED_LOG to LOG,
// ensure all entries in SequencedLeafData are integrated.

// Use the time truncated-to-millis throughout, as that's what's stored.

func (t *adminTX) SoftDeleteTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* deleted */ /* deleteTimeMillis */

func (t *adminTX) UndeleteTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* deleted */ /* deleteTimeMillis */

// updateDeleted updates the Deleted and DeleteTimeMillis fields of the specified tree.
// deleteTimeMillis must be either an int64 (in millis since epoch) or nil.
func (t *adminTX) updateDeleted(ctx context.Context, treeID int64, deleted bool, deleteTimeMillis interface{}) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) HardDeleteTree(ctx context.Context, treeID int64) error {
	_ = "STUB: not implemented"
	return nil
}

/* wantDeleted */

func validateDeleted(ctx context.Context, tx pgx.Tx, treeID int64, wantDeleted bool) error {
	_ = "STUB: not implemented"
	return nil
}
