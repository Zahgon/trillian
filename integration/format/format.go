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

// Package format contains an integration test which builds a log using an
// in-memory storage end-to-end, and makes sure the SubtreeProto storage format
// has no regressions.
package format

import (
	"context"

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/storage/storagepb"
	"github.com/transparency-dev/merkle"
)

func run(treeSize, batchSize int, leafFormat string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Read the latest LogRoot back.

func createTree(ctx context.Context, as storage.AdminStorage, ls storage.LogStorage) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateLeaves(count int, format string) []*trillian.LogLeaf {
	_ = "STUB: not implemented"
	return nil
}

func sequenceLeaves(ctx context.Context, ls storage.LogStorage, tree *trillian.Tree, leaves []*trillian.LogLeaf, batchSize int) error {
	_ = "STUB: not implemented"
	return nil
}

type treeAndRev struct {
	subtree  *storagepb.SubtreeProto
	revision int
}

func latestRevisions(ls storage.LogStorage, treeID int64, hasher merkle.LogHasher) (string, error) {
	_ = "STUB: not implemented"
	// vMap maps subtree prefixes (as strings) to the corresponding subtree proto and its revision
	return "", nil
}

// Relies on the btree key space for subtrees being /tree_id/subtree/id/revision.

// Store the keys in sorted order.

// The map should now contain the latest revisions per subtree.

// TODO(mhutchinson): This error should be propagated.
