// Copyright 2020 Google LLC. All Rights Reserved.
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

package storagetest

import (
	"context"
	"testing"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
)

// LogStorageFactory creates LogStorage and AdminStorage for a test to use.
type LogStorageFactory = func(ctx context.Context, t *testing.T) (storage.LogStorage, storage.AdminStorage)

// LogStorageTest executes a test using the given storage implementations.
type LogStorageTest = func(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage)

// RunLogStorageTests runs all the log storage tests against the provided log storage implementation.
func RunLogStorageTests(t *testing.T, storageFactory LogStorageFactory) {
	_ = "STUB: not implemented"
	return
}

func logTestFunctions(t *testing.T, x interface{}) map[string]LogStorageTest {
	_ = "STUB: not implemented"
	return nil
}

// Method exists but has the wrong type signature.

// logTests is a suite of tests to run against the storage.LogTest interface.
type logTests struct{}

func (*logTests) TestCheckDatabaseAccessible(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

func (*logTests) TestSnapshot(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

func (*logTests) TestReadWriteTransaction(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

func logTree(logID int64) *trillian.Tree { _ = "STUB: not implemented"; return nil }

// AddSequencedLeaves tests. ---------------------------------------------------

type addSequencedLeavesTest struct {
	t    *testing.T
	s    storage.LogStorage
	tree *trillian.Tree
}

func initAddSequencedLeavesTest(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) addSequencedLeavesTest {
	_ = "STUB: not implemented"
	return *new(addSequencedLeavesTest)
}

func (t *addSequencedLeavesTest) addSequencedLeaves(leaves []*trillian.LogLeaf) {
	_ = "STUB: not implemented"
	return

	// Time we will queue all leaves at.
}

// TODO(pavelkalinnikov): Verify returned status for each leaf.

func (t *addSequencedLeavesTest) verifySequencedLeaves(start, count int64, exp []*trillian.LogLeaf) {
	_ = "STUB: not implemented"
	return
}

func (*logTests) TestAddSequencedLeavesUnordered(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

func (*logTests) TestAddSequencedLeavesWithDuplicates(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// Full dup.

// Hash dup.
// Index dup.

// TODO: Remove when spannertest has transaction support.

// Time we'll request for guard cutoff in tests that don't test this (should include all above)
var fakeDequeueCutoffTime = time.Date(2016, 11, 10, 15, 16, 30, 0, time.UTC)

func (*logTests) TestDequeueLeavesNoneQueued(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// GetLeavesByRange tests. -----------------------------------------------------

type getLeavesByRangeTest struct {
	start, count int64
	want         []int64
	wantErr      bool
}

func testGetLeavesByRangeImpl(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage, create *trillian.Tree, tests []getLeavesByRangeTest) {
	_ = "STUB: not implemented"
	return
}

// Note: GetLeavesByRange loads the root internally to get the tree size.

// Create leaves [0]..[19] but drop leaf [5] and set the tree size to 14.

func (*logTests) TestGetLeavesByRangeFromLog(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// Starts right after tree size.
// Starts further away.
// Hits non-contiguous leaves.
// Starts from a missing leaf.
// Empty range.
// Negative start.
// Negative count.
// Starts after all stored leaves.

func (*logTests) TestGetLeavesByRangeFromPreorderedLog(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// Starts right after tree size.

// Starts further away.
// Hits non-contiguous leaves.
// Starts from a missing leaf.
// Empty range.
// Negative start.
// Negative count.
// Starts after all stored leaves.

// Time we will queue all leaves at
var fakeQueueTime = time.Date(2016, 11, 10, 15, 16, 27, 0, time.UTC)

func createFakeLeaf(ctx context.Context, s storage.LogStorage, tree *trillian.Tree, rawHash, hash, data, extraData []byte, seq int64, t *testing.T) *trillian.LogLeaf {
	_ = "STUB: not implemented"
	return nil
}

func (*logTests) TestDequeueLeaves(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// Now try to dequeue them
// Some dequeue implementations probabalistically dequeue and require retrying until timeout.
// Retry until timeout

// If we dequeue again then we should now get nothing

// dequeueAndSequence repeatedly dequeues in a single transaction until limit is reached or a timeout occurs.
// Then, it sequences the leaves with UpdateSequencedLeaves.
func dequeueAndSequence(ctx context.Context, t *testing.T, ls storage.LogStorage, tree *trillian.Tree, ts time.Time, limit int, startIndex int64) []*trillian.LogLeaf {
	_ = "STUB: not implemented"

	// We'll retry a few times if we get nothing back since we're now dependent
	// on the underlying queue delivering unsequenced entries.
	return nil
}

func ensureAllLeavesDistinct(t *testing.T, leaves []*trillian.LogLeaf) {
	_ = "STUB: not implemented"
	return
}

func ensureLeavesHaveQueueTimestamp(t *testing.T, leaves []*trillian.LogLeaf, want time.Time) {
	_ = "STUB: not implemented"
	return
}

func (*logTests) TestDequeueLeavesTwoBatches(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// Retry until timeout

// Plus the union of the leaf batches should all have distinct hashes

// If we dequeue again then we should now get nothing

func (*logTests) TestAddSequencedLeavesAndDequeueLeaves(ctx context.Context, t *testing.T, s storage.LogStorage, as storage.AdminStorage) {
	_ = "STUB: not implemented"
	return
}

// Check that the first sequenced entry is returned.

// Fake a signing run that signs 1 entry.

// Check that the 2nd and 3rd sequenced entries are returned.
