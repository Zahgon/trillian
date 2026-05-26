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

package memory

import (
	"context"
	"sync"

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
)

// NewAdminStorage returns a storage.AdminStorage implementation backed by
// TreeStorage.
func NewAdminStorage(ms *TreeStorage) storage.AdminStorage {
	_ = "STUB: not implemented"
	return *new(storage.AdminStorage)
}

// memoryAdminStorage implements storage.AdminStorage
type memoryAdminStorage struct {
	ms *TreeStorage
}

func (s *memoryAdminStorage) Snapshot(ctx context.Context) (storage.ReadOnlyAdminTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyAdminTX), nil
}

func (s *memoryAdminStorage) ReadWriteTransaction(ctx context.Context, f storage.AdminTXFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *memoryAdminStorage) CheckDatabaseAccessible(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type adminTX struct {
	ms *TreeStorage

	// mu guards reads/writes on closed, which happen on Commit/Close methods.
	//
	// We don't check closed on methods apart from the ones above, as we trust tx
	// to keep tabs on its state, and hence fail to do queries after closed.
	mu     sync.RWMutex
	closed bool
}

func (t *adminTX) Commit() error {
	_ = "STUB: not implemented"
	// TODO(al): The admin implementation isn't transactional.
	return nil
}

func (t *adminTX) Close() error {
	_ = "STUB: not implemented"
	// TODO(al): The admin implementation isn't transactional.
	return nil
}

func (t *adminTX) GetTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) ListTrees(ctx context.Context, includeDeleted bool) ([]*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) CreateTree(ctx context.Context, tr *trillian.Tree) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) UpdateTree(ctx context.Context, treeID int64, updateFunc func(*trillian.Tree)) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) SoftDeleteTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *adminTX) HardDeleteTree(ctx context.Context, treeID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *adminTX) UndeleteTree(ctx context.Context, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateStorageSettings(tree *trillian.Tree) error { _ = "STUB: not implemented"; return nil }
