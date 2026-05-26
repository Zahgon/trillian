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

package mysql

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
			PrivateKey, -- Unused
			PublicKey, -- Used to store StorageSettings
			MaxRootDurationMillis,
			Deleted,
			DeleteTimeMillis
		FROM Trees`
	selectNonDeletedTrees = selectTrees + nonDeletedWhere
	selectTreeByID        = selectTrees + " WHERE TreeId = ?"

	updateTreeSQL = `UPDATE Trees
		SET TreeState = ?, TreeType = ?, DisplayName = ?, Description = ?, UpdateTimeMillis = ?, MaxRootDurationMillis = ?, PrivateKey = ?
		WHERE TreeId = ?`
)

// NewAdminStorage returns a MySQL storage.AdminStorage implementation backed by DB.
func NewAdminStorage(db *sql.DB) *mysqlAdminStorage { _ = "STUB: not implemented"; return nil }

// mysqlAdminStorage implements storage.AdminStorage
type mysqlAdminStorage struct {
	db *sql.DB
}

func (s *mysqlAdminStorage) Snapshot(ctx context.Context) (storage.ReadOnlyAdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyAdminTX), nil
}

func (s *mysqlAdminStorage) beginInternal(ctx context.Context) (storage.AdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.AdminTX), nil
}

/* opts */

func (s *mysqlAdminStorage) ReadWriteTransaction(ctx context.Context, f storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mysqlAdminStorage) CheckDatabaseAccessible(ctx context.Context) error {
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

// When creating a new tree we automatically add StorageSettings to allow us to
// determine that this tree can support newer storage features. When reading
// trees that do not have this StorageSettings populated, it must be assumed that
// the tree was created with the oldest settings.
// The gist of this code is super simple: create a new StorageSettings with the most
// modern defaults if the created tree does not have one, and then create a struct that
// represents this to store in the DB. Unfortunately because this involves anypb, struct
// copies, marshalling, and proper error handling this turns into a scary amount of code.

// Default behaviour for new trees is to skip writing subtree revisions.

// Unused, filling in for backward compatibility.
// Unused, filling in for backward compatibility.
// Unused, filling in for backward compatibility.

// PrivateKey: Unused, filling in for backward compatibility.
// Using the otherwise unused PublicKey for storing StorageSettings.

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

// PrivateKey: Unused, filling in for backward compatibility.
// PublicKey should not be updated with any storageSettings here without
// a lot of thought put into it. At the moment storageSettings are inferred
// when reading the tree, even if no value is stored in the database.

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

// No storage settings is OK, we'll just use the defaults for new trees

// storageSettings allows us to persist storage settings to the DB.
// It is a tempting trap to use protos for this, but the way they encode
// makes it impossible to tell the difference between no value ever written
// and a value that was written with the default values for each field.
// Using an explicit struct and gob encoding allows us to tell the difference.
type storageSettings struct {
	Revisioned bool
}
