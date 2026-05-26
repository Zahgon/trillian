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
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/storage/cloudspanner/spannerpb"
	"google.golang.org/protobuf/proto"
)

var (
	// NumUnseqBuckets is the length of the unsequenced time ring buffer.
	NumUnseqBuckets = int64(4)
	// NumMerkleBuckets is the number of individual buckets below each unsequenced ring buffer.
	NumMerkleBuckets = int64(16)
	// TimeNow is the function used to get the current time. Exposed so it may be mocked by tests.
	TimeNow = time.Now

	treeStateMap = map[trillian.TreeState]spannerpb.TreeState{
		trillian.TreeState_ACTIVE: spannerpb.TreeState_ACTIVE,
		trillian.TreeState_FROZEN: spannerpb.TreeState_FROZEN,
	}
	treeTypeMap = map[trillian.TreeType]spannerpb.TreeType{
		trillian.TreeType_LOG:            spannerpb.TreeType_LOG,
		trillian.TreeType_PREORDERED_LOG: spannerpb.TreeType_PREORDERED_LOG,
	}

	treeStateReverseMap = reverseTreeStateMap(treeStateMap)
	treeTypeReverseMap  = reverseTreeTypeMap(treeTypeMap)
)

const nanosPerMilli = int64(time.Millisecond / time.Nanosecond)

func reverseTreeStateMap(m map[trillian.TreeState]spannerpb.TreeState) map[spannerpb.TreeState]trillian.TreeState {
	_ = "STUB: not implemented"
	return nil
}

func reverseTreeTypeMap(m map[trillian.TreeType]spannerpb.TreeType) map[spannerpb.TreeType]trillian.TreeType {
	_ = "STUB: not implemented"
	return nil
}

// adminTX implements both storage.ReadOnlyAdminTX and storage.AdminTX.
type adminTX struct {
	client *spanner.Client

	// mu guards tx, but it's only actively used for Commit/Close. In other
	// scenarios we trust Spanner to blow up if you try to use a closed tx.
	//
	// Note that, if tx is a spanner.SnapshotTransaction, it'll be set to nil
	// when adminTX is closed.
	mu sync.RWMutex

	// tx is either spanner.ReadOnlyTransaction or spanner.ReadWriteTransaction,
	// according to the role adminTX is meant to fill.
	//
	// If tx is a snapshot transaction it'll be set to nil when adminTX is closed
	// to avoid reuse.
	tx spanRead
}

// adminStorage implements storage.AdminStorage.
type adminStorage struct {
	client *spanner.Client
}

// NewAdminStorage returns a Spanner-based storage.AdminStorage implementation.
func NewAdminStorage(client *spanner.Client) storage.AdminStorage {
	_ = "STUB: not implemented"
	return *new(storage.AdminStorage)
}

// CheckDatabaseAccessible implements AdminStorage.CheckDatabaseAccessible.
func (s *adminStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Snapshot implements AdminStorage.Snapshot.
func (s *adminStorage) Snapshot(ctx context.Context) (storage.ReadOnlyAdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyAdminTX), nil
}

// Begin implements AdminStorage.Begin.
func (s *adminStorage) Begin(ctx context.Context) (storage.AdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.AdminTX), nil
}

// ReadWriteTransaction implements AdminStorage.ReadWriteTransaction.
func (s *adminStorage) ReadWriteTransaction(ctx context.Context, f storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// Commit implements ReadOnlyAdminTX.Commit.
func (t *adminTX) Commit() error {
	_ = "STUB: not implemented"

	// Close implements ReadOnlyAdminTX.Close.
	return nil
}

func (t *adminTX) Close() error { _ = "STUB: not implemented"; return nil }

// tx will be committed by ReadWriteTransaction(), so only close readonly tx here

// GetTree implements ReadOnlyAdminTX.GetTree.
func (t *adminTX) GetTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) getTreeInfo(ctx context.Context, treeID int64) (*spannerpb.TreeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Improve on the error message

// info.TreeState,
// info.TreeType,

// TODO(al): check other denormalisations are consistent too.

// Sanity checks

// ListTrees implements ReadOnlyAdminTX.ListTrees.
func (t *adminTX) ListTrees(ctx context.Context, includeDeleted bool) ([]*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* idOnly */

func (t *adminTX) readTrees(ctx context.Context, includeDeleted, idOnly bool, f func(*spanner.Row) error) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTree implements AdminWriter.CreateTree.
func (t *adminTX) CreateTree(ctx context.Context, tree *trillian.Tree) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newTreeInfo creates a new TreeInfo from a Tree. Meant to be used for new trees.
func newTreeInfo(tree *trillian.Tree, treeID int64, now time.Time) (*spannerpb.TreeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logConfigOrDefault(tree *trillian.Tree) (*spannerpb.LogStorageConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateTree implements AdminWriter.UpdateTree.
func (t *adminTX) UpdateTree(ctx context.Context, treeID int64, updateFunc func(*trillian.Tree)) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update (just) the mutable fields in treeInfo.

func (t *adminTX) updateTreeInfo(ctx context.Context, info *spannerpb.TreeInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// SoftDeleteTree implements AdminWriter.SoftDeleteTree.
func (t *adminTX) SoftDeleteTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HardDeleteTree implements AdminWriter.HardDeleteTree.
func (t *adminTX) HardDeleteTree(ctx context.Context, treeID int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Due to cloud spanner sizing recommendations, we don't interleave our tables
// which means no ON DELETE CASCADE goodies for us, so we have to
// transactionally delete related data from all tables.

// UndeleteTree implements AdminWriter.UndeleteTree.
func (t *adminTX) UndeleteTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toTrillianTree(info *spannerpb.TreeInfo) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unmarshalSettings returns the message obtained from tree.StorageSettings.
// If tree.StorageSettings is nil no unmarshaling will be attempted; instead the method will return
// (nil, nil).
func unmarshalSettings(tree *trillian.Tree) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}

func validateLogStorageConfig(config *spannerpb.LogStorageConfig) error {
	_ = "STUB: not implemented"
	return nil
}
