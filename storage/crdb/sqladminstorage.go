// Copyright 2017 Trillian Authors
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

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
)

const (
	defaultSequenceIntervalSeconds = 60

	nonDeletedWhere = " WHERE (Deleted IS NULL OR Deleted = 'false')"

	selectTrees = `
		SELECT
			TreeId,
			TreeState,
			TreeType,
			HashStrategy,
			HashAlgorithm,
			SignatureAlgorithm,
			DisplayName,
			Description,
			CreateTimeMillis,
			UpdateTimeMillis,
			PrivateKey,
			PublicKey,
			MaxRootDurationMillis,
			Deleted,
			DeleteTimeMillis
		FROM Trees`
	selectNonDeletedTrees = selectTrees + nonDeletedWhere
	selectTreeByID        = selectTrees + " WHERE TreeId = $1"

	updateTreeSQL = `UPDATE Trees
		SET TreeState = $1, TreeType = $2, DisplayName = $3, Description = $4, UpdateTimeMillis = $5, MaxRootDurationMillis = $6, PrivateKey = $7
		WHERE TreeId = $8`
)

// NewSQLAdminStorage returns a SQL storage.AdminStorage implementation backed by DB.
// Should work for MySQL and CockroachDB
func NewSQLAdminStorage(db *sql.DB) storage.AdminStorage {
	_ = "STUB: not implemented"
	return *new(storage.AdminStorage)
}

// sqlAdminStorage implements storage.AdminStorage
type sqlAdminStorage struct {
	db *sql.DB
}

func (s *sqlAdminStorage) Snapshot(ctx context.Context) (storage.ReadOnlyAdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyAdminTX), nil
}

func (s *sqlAdminStorage) beginInternal(ctx context.Context) (storage.AdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.AdminTX), nil
}

/* opts */

func (s *sqlAdminStorage) ReadWriteTransaction(ctx context.Context, f storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlAdminStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type adminTX struct {
	tx *sql.Tx

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
	return nil, nil
}

// GetTree is an entry point for most RPCs, let's provide somewhat nicer error messages.

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

// Unused, filling in for backward compatibility.
// Unused, filling in for backward compatibility.
// Unused, filling in for backward compatibility.

// Unused, filling in for backward compatibility.
// Unused, filling in for backward compatibility.

// MySQL silently truncates data when running in non-strict mode.
// We shouldn't be using non-strict modes, but let's guard against it
// anyway.

// GetTree will fail for truncated enums (they get recorded as
// empty strings, which will not match any known value).

/* SigningEnabled */
/* SequencingEnabled */

func (t *adminTX) UpdateTree(ctx context.Context, treeID int64, updateFunc func(*trillian.Tree)) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(pavelkalinnikov): When switching TreeType from PREORDERED_LOG to LOG,
// ensure all entries in SequencedLeafData are integrated.

// Use the time truncated-to-millis throughout, as that's what's stored.

// Unused, filling in for backward compatibility.

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

// TreeControl didn't have "ON DELETE CASCADE" on previous versions, so let's hit it explicitly

func validateDeleted(ctx context.Context, tx *sql.Tx, treeID int64, wantDeleted bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStorageSettings(tree *trillian.Tree) error { _ = "STUB: not implemented"; return nil }
