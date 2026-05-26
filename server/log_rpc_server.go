// Copyright 2016 Google LLC. All Rights Reserved.
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

package server

import (
	"context"

	"github.com/google/trillian"
	"github.com/google/trillian/extension"
	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	"github.com/google/trillian/trees"
	"github.com/google/trillian/util/clock"
	"github.com/transparency-dev/merkle"
)

// TODO: There is no access control in the server yet and clients could easily modify
// any tree.

const traceSpanRoot = "/trillian"

var (
	optsLogInit            = trees.NewGetOpts(trees.Admin, trillian.TreeType_LOG, trillian.TreeType_PREORDERED_LOG)
	optsLogRead            = trees.NewGetOpts(trees.Query, trillian.TreeType_LOG, trillian.TreeType_PREORDERED_LOG)
	optsLogWrite           = trees.NewGetOpts(trees.QueueLog, trillian.TreeType_LOG)
	optsPreorderedLogWrite = trees.NewGetOpts(trees.SequenceLog, trillian.TreeType_PREORDERED_LOG)
)

// TrillianLogRPCServer implements the RPC API defined in the proto
type TrillianLogRPCServer struct {
	registry              extension.Registry
	timeSource            clock.TimeSource
	leafCounter           monitoring.Counter
	proofIndexPercentiles monitoring.Histogram
	fetchedLeaves         monitoring.Counter
}

// NewTrillianLogRPCServer creates a new RPC server backed by a LogStorageProvider.
func NewTrillianLogRPCServer(registry extension.Registry, timeSource clock.TimeSource) *TrillianLogRPCServer {
	_ = "STUB: not implemented"
	return nil
}

// IsHealthy returns nil if the server is healthy, error otherwise.
func (t *TrillianLogRPCServer) IsHealthy() error { _ = "STUB: not implemented"; return nil }

// QueueLeaf submits one leaf to the queue.
func (t *TrillianLogRPCServer) QueueLeaf(ctx context.Context, req *trillian.QueueLeafRequest) (*trillian.QueueLeafResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mirror the use of this counter in AddSequencedLeaves below.

func hashLeaves(leaves []*trillian.LogLeaf, hasher merkle.LogHasher) {
	_ = "STUB: not implemented"
	return
}

// AddSequencedLeaves submits a batch of sequenced leaves to a pre-ordered log
// for later integration into its underlying tree.
func (t *TrillianLogRPCServer) AddSequencedLeaves(ctx context.Context, req *trillian.AddSequencedLeavesRequest) (*trillian.AddSequencedLeavesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInclusionProof obtains the proof of inclusion in the tree for a leaf that has been sequenced.
// Similar to the get proof by hash handler but one less step as we don't need to look up the index
func (t *TrillianLogRPCServer) GetInclusionProof(ctx context.Context, req *trillian.GetInclusionProofRequest) (*trillian.GetInclusionProofResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next we need to make sure the requested tree size corresponds to an STH, so that we
// have a usable tree revision

// GetInclusionProofByHash obtains proofs of inclusion by leaf hash. Because some logs can
// contain duplicate hashes it is possible for multiple proofs to be returned.
func (t *TrillianLogRPCServer) GetInclusionProofByHash(ctx context.Context, req *trillian.GetInclusionProofByHashRequest) (*trillian.GetInclusionProofByHashResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next we need to make sure the requested tree size corresponds to an STH, so that we
// have a usable tree revision

// Find the leaf index of the supplied hash

// TODO(Martin2112): Need to define a limit on number of results or some form of paging etc.

// Don't include leaves that aren't in the requested TreeSize.

// TODO(gbelvin): Rename "Proof" -> "Proofs"

// GetConsistencyProof obtains a proof that two versions of the tree are consistent with each
// other and that the later tree includes all the entries of the prior one. For more details
// see the example trees in RFC 6962.
func (t *TrillianLogRPCServer) GetConsistencyProof(ctx context.Context, req *trillian.GetConsistencyProofRequest) (*trillian.GetConsistencyProofResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to get consistency proof

// We have everything we need. Return the proof

// GetLatestSignedLogRoot obtains the latest published tree root for the Merkle Tree that
// underlies the log.
func (t *TrillianLogRPCServer) GetLatestSignedLogRoot(ctx context.Context, req *trillian.GetLatestSignedLogRootRequest) (*trillian.GetLatestSignedLogRootResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no need to get consistency proof in this case

// Try to get consistency proof

// We have everything we need. Return the response

func tryGetConsistencyProof(ctx context.Context, firstTreeSize, secondTreeSize uint64, tx storage.ReadOnlyLogTreeTX, hasher merkle.LogHasher) (*trillian.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLeavesByRange obtains leaves based on a range of sequence numbers within the tree.
// This only fetches sequenced leaves; leaves that have been queued but not yet integrated
// are not visible.
func (t *TrillianLogRPCServer) GetLeavesByRange(ctx context.Context, req *trillian.GetLeavesByRangeRequest) (*trillian.GetLeavesByRangeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEntryAndProof returns both a Merkle Leaf entry and an inclusion proof for a given index
// and tree size.
func (t *TrillianLogRPCServer) GetEntryAndProof(ctx context.Context, req *trillian.GetEntryAndProofRequest) (*trillian.GetEntryAndProofResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next we need to make sure the requested tree size corresponds to an STH, so that we
// have a usable tree revision

// return latest proof we can manage

// We also need the leaf entry

// Work is complete, we have everything we need for the response

func (t *TrillianLogRPCServer) commitAndLog(ctx context.Context, logID int64, tx storage.ReadOnlyLogTreeTX, op string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrillianLogRPCServer) closeAndLog(ctx context.Context, logID int64, tx storage.ReadOnlyLogTreeTX, op string) {
	_ = "STUB: not implemented"
	return
}

// getInclusionProofForLeafIndex is used by multiple handlers. It does the storage fetching
// and makes additional checks on the returned proof. Returns a Proof suitable for inclusion in
// an RPC response
func getInclusionProofForLeafIndex(ctx context.Context, tx storage.ReadOnlyLogTreeTX, hasher merkle.LogHasher, size, leafIndex uint64) (*trillian.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrillianLogRPCServer) getTreeAndHasher(ctx context.Context, treeID int64, opts trees.GetOpts) (*trillian.Tree, merkle.LogHasher, error) {
	_ = "STUB: not implemented"
	return nil, *new(merkle.LogHasher), nil
}

func (t *TrillianLogRPCServer) getTreeAndContext(ctx context.Context, treeID int64, opts trees.GetOpts) (*trillian.Tree, context.Context, error) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context), nil
}

// InitLog initialises a freshly created Log by creating the first STH with
// size 0.
func (t *TrillianLogRPCServer) InitLog(ctx context.Context, req *trillian.InitLogRequest) (*trillian.InitLogResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Belt and braces check.

func (t *TrillianLogRPCServer) recordIndexPercent(leafIndex int64, treeSize uint64) {
	_ = "STUB: not implemented"

	// Work out what percentage of the current log size this index corresponds to.
	return
}

func spanFor(ctx context.Context, name string) (context.Context, func()) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *TrillianLogRPCServer) snapshotForTree(ctx context.Context, tree *trillian.Tree, method string) (storage.ReadOnlyLogTreeTX, error) {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyLogTreeTX), nil
}

// Special case to handle ErrTreeNeedsInit, which leaves the TX open.
// To avoid leaking it make sure it's closed.
