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
	"github.com/google/trillian/types"
)

// runLogTX is a helps avoid copying out "if err != nil { blah }" all over the place
func runLogTX(s storage.LogStorage, tree *trillian.Tree, t *testing.T, f storage.LogTXFunc) {
	_ = "STUB: not implemented"
	return
}

// createTestLeaves creates some test leaves with predictable data
func createTestLeaves(n, startSeq int64) []*trillian.LogLeaf { _ = "STUB: not implemented"; return nil }

func mustSignAndStoreLogRoot(ctx context.Context, t *testing.T, l storage.LogStorage, tree *trillian.Tree, r *types.LogRootV1) {
	_ = "STUB: not implemented"
	return
}

func dequeueLeavesInTx(ctx context.Context, ls storage.LogStorage, tree *trillian.Tree, t time.Time, limit int) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
