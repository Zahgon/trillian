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

// Package client verifies responses from the Trillian log.
package client

import (
	"context"
	"sync"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/types"
	"github.com/transparency-dev/merkle"
)

// LogClient represents a client for a given Trillian log instance.
type LogClient struct {
	*LogVerifier
	LogID         int64
	MinMergeDelay time.Duration
	client        trillian.TrillianLogClient
	root          types.LogRootV1
	rootLock      sync.Mutex
	updateLock    sync.Mutex
}

// New returns a new LogClient.
func New(logID int64, client trillian.TrillianLogClient, verifier *LogVerifier, root types.LogRootV1) *LogClient {
	_ = "STUB: not implemented"
	return nil
}

// NewFromTree creates a new LogClient given a tree config.
func NewFromTree(client trillian.TrillianLogClient, config *trillian.Tree, root types.LogRootV1) (*LogClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddLeaf adds leaf to the append only log.
// Blocks and continuously updates the trusted root until a successful inclusion proof
// can be retrieved.
func (c *LogClient) AddLeaf(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ListByIndex returns the requested leaves by index.
func (c *LogClient) ListByIndex(ctx context.Context, start, count int64) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify that we got back the requested leaves.

// WaitForRootUpdate repeatedly fetches the latest root until there is an
// update, which it then applies, or until ctx times out.
func (c *LogClient) WaitForRootUpdate(ctx context.Context) (*types.LogRootV1, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retry.

// getAndVerifyLatestRoot fetches and verifies the latest root against a trusted root, seen in the past.
// Pass nil for trusted if this is the first time querying this log.
func (c *LogClient) getAndVerifyLatestRoot(ctx context.Context, trusted *types.LogRootV1) (*types.LogRootV1, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(gbelvin): Turn on root verification.
/*
	logRoot, err := c.VerifyRoot(&types.LogRootV1{}, resp.GetSignedLogRoot(), nil)
	if err != nil {
		return nil, err
	}
*/
// TODO(gbelvin): Remove this hack when all implementations store digital signatures.

// Tree has not been updated.

// Verify root update if the tree / the latest signed log root isn't empty.

// GetRoot returns a copy of the latest trusted root.
func (c *LogClient) GetRoot() *types.LogRootV1 { _ = "STUB: not implemented"; return nil }

// Copy the internal trusted root in order to prevent clients from modifying it.

// UpdateRoot retrieves the current SignedLogRoot, verifying it against roots this client has
// seen in the past, and updating the currently trusted root if the new root verifies, and is
// newer than the currently trusted root.
func (c *LogClient) UpdateRoot(ctx context.Context) (*types.LogRootV1, error) {
	_ = "STUB: not implemented"
	// Only one root update should be running at any point in time, because
	// the update involves a consistency proof from the old value, and if the
	// old value could change along the way (in another goroutine) then the
	// result could be inconsistent.
	//
	// For example, if the current root is A and two root updates A->B and A->C
	// happen in parallel, then we might end up with the transitions A->B->C:
	//     cur := A            cur := A
	//    getRoot() => B      getRoot() => C
	//    proof(A->B) ok      proof(A->C) ok
	//    c.root = B
	//                        c.root = C
	// and the last step (B->C) has no proof and so could hide a forked tree.
	return nil, nil
}

// Lock "rootLock" for the "root" update.

// Take a copy of the new trusted root in order to prevent clients from modifying it.

// WaitForInclusion blocks until the requested data has been verified with an
// inclusion proof.
//
// It will continuously update the root to the latest one available until the
// data is found, or an error is returned.
//
// It is best to call this method with a context that will timeout to avoid
// waiting forever.
func (c *LogClient) WaitForInclusion(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// If a minimum merge delay has been configured, wait at least that long before
// starting to poll

// It is illegal to ask for an inclusion proof with TreeSize = 0.

// If not found or tree is empty, wait for a root update before retrying again.

// Retry

func (c *LogClient) getAndVerifyInclusionProof(ctx context.Context, leafHash []byte, sth *types.LogRootV1) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddSequencedLeaves adds any number of pre-sequenced leaves to the log.
// Indexes must be contiguous.
func (c *LogClient) AddSequencedLeaves(ctx context.Context, dataByIndex map[int64][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Check index continuity.

// QueueLeaf adds a leaf to a Trillian log without blocking.
// AlreadyExists is considered a success case by this function.
func (c *LogClient) QueueLeaf(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// prepareLeaf returns a trillian.LogLeaf prepopulated with leaf data and hash.
func prepareLeaf(hasher merkle.LogHasher, data []byte) *trillian.LogLeaf {
	_ = "STUB: not implemented"
	return nil
}
