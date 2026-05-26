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

package testonly

import (
	"context"
	"testing"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	// LogTree is a valid, LOG-type trillian.Tree for tests.
	LogTree = &trillian.Tree{
		TreeState:       trillian.TreeState_ACTIVE,
		TreeType:        trillian.TreeType_LOG,
		DisplayName:     "Llamas Log",
		Description:     "Registry of publicly-owned llamas",
		MaxRootDuration: durationpb.New(0 * time.Millisecond),
	}

	// PreorderedLogTree is a valid, PREORDERED_LOG-type trillian.Tree for tests.
	PreorderedLogTree = &trillian.Tree{
		TreeState:       trillian.TreeState_ACTIVE,
		TreeType:        trillian.TreeType_PREORDERED_LOG,
		DisplayName:     "Pre-ordered Log",
		Description:     "Mirror registry of publicly-owned llamas",
		MaxRootDuration: durationpb.New(0 * time.Millisecond),
	}
)

// AdminStorageTester runs a suite of tests against AdminStorage implementations.
type AdminStorageTester struct {
	// NewAdminStorage returns an AdminStorage instance pointing to a clean
	// test database.
	NewAdminStorage func() storage.AdminStorage
}

// RunAllTests runs all AdminStorage tests.
func (tester *AdminStorageTester) RunAllTests(t *testing.T) { _ = "STUB: not implemented"; return }

// TestCreateTree tests AdminStorage Tree creation.
func (tester *AdminStorageTester) TestCreateTree(t *testing.T) {
	_ = "STUB: not implemented"
	// Check that validation runs, but leave details to the validation
	// tests.
	return
}

// Test CreateTree up to the tx commit

// Tested above

// Ignore storage_settings changes (OK to vary between implementations)

// TestUpdateTree tests AdminStorage Tree updates.
func (tester *AdminStorageTester) TestUpdateTree(t *testing.T) { _ = "STUB: not implemented"; return }

// Test for an unknown tree outside the loop: it makes the test logic simpler

// Copy storage-generated values to want before comparing

// Ignore storage_settings changes (OK to vary between implementations)

// TestListTrees tests ListTrees.
func (tester *AdminStorageTester) TestListTrees(t *testing.T) { _ = "STUB: not implemented"; return }

// Always return nil, as we're reporting errors independently above.

// Capture Begin() / Commit() errors

// Do a first pass with an empty DB
/* includeDeleted */ /* wantTrees */
/* includeDeleted */ /* wantTrees */

// Add some trees and do another pass

/* includeDeleted */
/* includeDeleted */

func runListTreesTest(ctx context.Context, tx storage.ReadOnlyAdminTX, includeDeleted bool, wantTrees []*trillian.Tree) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore storage_settings changes (OK to vary between implementations)

// TestSoftDeleteTree tests success scenarios of SoftDeleteTree.
func (tester *AdminStorageTester) TestSoftDeleteTree(t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Ignore storage_settings changes (OK to vary between implementations)

// TestSoftDeleteTreeErrors tests error scenarios of SoftDeleteTree.
func (tester *AdminStorageTester) TestSoftDeleteTreeErrors(t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// TestHardDeleteTree tests success scenarios of HardDeleteTree.
func (tester *AdminStorageTester) TestHardDeleteTree(t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// TestHardDeleteTreeErrors tests error scenarios of HardDeleteTree.
func (tester *AdminStorageTester) TestHardDeleteTreeErrors(t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// TestUndeleteTree tests success scenarios of UndeleteTree.
func (tester *AdminStorageTester) TestUndeleteTree(t *testing.T) { _ = "STUB: not implemented"; return }

// TestUndeleteTreeErrors tests error scenarios of UndeleteTree.
func (tester *AdminStorageTester) TestUndeleteTreeErrors(t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// TestAdminTXReadWriteTransaction tests the ReadWriteTransaction method on AdminStorage.
func (tester *AdminStorageTester) TestAdminTXReadWriteTransaction(t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Multiple Close() calls are fine too

// assertStoredTree verifies that "want" is equal to the tree stored under its ID.
func assertStoredTree(ctx context.Context, s storage.AdminStorage, want *trillian.Tree) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore storage_settings changes (OK to vary between implementations)

type spec struct {
	Tree            *trillian.Tree
	Frozen, Deleted bool
}

// makeTreeOrFail delegates to makeTree. If makeTree returns a non-nil error, failFn is called.
func makeTreeOrFail(ctx context.Context, s storage.AdminStorage, spec spec, failFn func(string, ...interface{})) *trillian.Tree {
	_ = "STUB: not implemented"
	return nil
}

// makeTree creates a tree and updates it to Frozen and/or Deleted, according to "spec".
func makeTree(ctx context.Context, s storage.AdminStorage, spec spec) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sanity checks

func tweakedCopy(tree *trillian.Tree, modFn func(t *trillian.Tree)) *trillian.Tree {
	_ = "STUB: not implemented"
	return nil
}
