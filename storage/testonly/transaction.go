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

package testonly

import (
	"context"
	"errors"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
)

// RunOnLogTX is a helper for mocking out the LogStorage.ReadWriteTransaction method.
func RunOnLogTX(tx storage.LogTreeTX) func(ctx context.Context, treeID int64, f storage.LogTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RunOnAdminTX is a helper for mocking out the AdminStorage.ReadWriteTransaction method.
func RunOnAdminTX(tx storage.AdminTX) func(ctx context.Context, f storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrNotImplemented is returned by unimplemented methods on the storage fakes.
var ErrNotImplemented = errors.New("not implemented")

// FakeLogStorage is a LogStorage implementation which is used for testing.
type FakeLogStorage struct {
	TX         storage.LogTreeTX
	ReadOnlyTX storage.ReadOnlyLogTreeTX

	TXErr                 error
	QueueLeavesErr        error
	AddSequencedLeavesErr error
}

// GetActiveLogIDs implements LogStorage.GetActiveLogIDs.
func (f *FakeLogStorage) GetActiveLogIDs(ctx context.Context) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// SnapshotForTree implements LogStorage.SnapshotForTree
}

func (f *FakeLogStorage) SnapshotForTree(ctx context.Context, _ *trillian.Tree) (storage.ReadOnlyLogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyLogTreeTX), nil
}

// ReadWriteTransaction implements LogStorage.ReadWriteTransaction
func (f *FakeLogStorage) ReadWriteTransaction(ctx context.Context, tree *trillian.Tree, fn storage.LogTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// QueueLeaves implements LogStorage.QueueLeaves.
func (f *FakeLogStorage) QueueLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, queueTimestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddSequencedLeaves implements LogStorage.AddSequencedLeaves.
func (f *FakeLogStorage) AddSequencedLeaves(ctx context.Context, tree *trillian.Tree, leaves []*trillian.LogLeaf, timestamp time.Time) ([]*trillian.QueuedLogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckDatabaseAccessible implements LogStorage.CheckDatabaseAccessible
func (f *FakeLogStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"

	// FakeAdminStorage is a AdminStorage implementation which is used for testing.
	return nil
}

type FakeAdminStorage struct {
	TX          []storage.AdminTX
	ReadOnlyTX  []storage.ReadOnlyAdminTX
	TXErr       []error
	SnapshotErr []error
}

// Begin implements AdminStorage.Begin
func (f *FakeAdminStorage) Begin(ctx context.Context) (storage.AdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.AdminTX), nil
}

// Snapshot implements AdminStorage.Snapshot
func (f *FakeAdminStorage) Snapshot(ctx context.Context) (storage.ReadOnlyAdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyAdminTX), nil
}

// ReadWriteTransaction implements AdminStorage.ReadWriteTransaction
func (f *FakeAdminStorage) ReadWriteTransaction(ctx context.Context, fn storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckDatabaseAccessible implements AdminStorage.CheckDatabaseAccessible
func (f *FakeAdminStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
