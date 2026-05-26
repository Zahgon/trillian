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

package storage

import (
	"context"

	"github.com/google/trillian"
)

const traceSpanRoot = "/trillian/storage"

// GetTree reads a tree from storage using a snapshot transaction.
// It's a convenience wrapper around RunInAdminSnapshot and ReadOnlyAdminTX's GetTree.
// See RunInAdminSnapshot if you need to perform more than one action per transaction.
func GetTree(ctx context.Context, admin AdminStorage, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListTrees reads trees from storage using a snapshot transaction.
// It's a convenience wrapper around RunInAdminSnapshot and ReadOnlyAdminTX's ListTrees.
// See RunInAdminSnapshot if you need to perform more than one action per transaction.
func ListTrees(ctx context.Context, admin AdminStorage, includeDeleted bool) ([]*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateTree creates a tree in storage.
// It's a convenience wrapper around ReadWriteTransaction and AdminWriter's CreateTree.
// See ReadWriteTransaction if you need to perform more than one action per transaction.
func CreateTree(ctx context.Context, admin AdminStorage, tree *trillian.Tree) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateTree updates a tree in storage.
// It's a convenience wrapper around ReadWriteTransaction and AdminWriter's UpdateTree.
// See ReadWriteTransaction if you need to perform more than one action per transaction.
func UpdateTree(ctx context.Context, admin AdminStorage, treeID int64, fn func(*trillian.Tree)) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SoftDeleteTree soft-deletes a tree in storage.
// It's a convenience wrapper around ReadWriteTransaction and AdminWriter's SoftDeleteTree.
// See ReadWriteTransaction if you need to perform more than one action per transaction.
func SoftDeleteTree(ctx context.Context, admin AdminStorage, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HardDeleteTree hard-deletes a tree from storage.
// It's a convenience wrapper around ReadWriteTransaction and AdminWriter's HardDeleteTree.
// See ReadWriteTransaction if you need to perform more than one action per transaction.
func HardDeleteTree(ctx context.Context, admin AdminStorage, treeID int64) error {
	_ = "STUB: not implemented"
	return nil
}

// UndeleteTree undeletes a tree in storage.
// It's a convenience wrapper around ReadWriteTransaction and AdminWriter's UndeleteTree.
// See ReadWriteTransaction if you need to perform more than one action per transaction.
func UndeleteTree(ctx context.Context, admin AdminStorage, treeID int64) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunInAdminSnapshot runs fn against a ReadOnlyAdminTX and commits if no error is returned.
func RunInAdminSnapshot(ctx context.Context, admin AdminStorage, fn func(tx ReadOnlyAdminTX) error) error {
	_ = "STUB: not implemented"
	return nil
}

func spanFor(ctx context.Context, name string) (context.Context, func()) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
