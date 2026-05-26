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

package integration

import (
	"context"
	"time"

	"github.com/google/trillian"
	inmemory "github.com/transparency-dev/merkle/testonly"
)

// TestParameters bundles up all the settings for a test run
type TestParameters struct {
	TreeID              int64
	CheckLogEmpty       bool
	QueueLeaves         bool
	AwaitSequencing     bool
	StartLeaf           int64
	LeafCount           int64
	UniqueLeaves        int64
	QueueBatchSize      int
	SequencerBatchSize  int
	ReadBatchSize       int64
	SequencingWaitTotal time.Duration
	SequencingPollWait  time.Duration
	RPCRequestDeadline  time.Duration
	CustomLeafPrefix    string
}

// DefaultTestParameters builds a TestParameters object for a normal
// test of the given log.
func DefaultTestParameters(treeID int64) TestParameters {
	_ = "STUB: not implemented"
	return *new(TestParameters)
}

type consistencyProofParams struct {
	size1 int64
	size2 int64
}

// inclusionProofTestIndices are the 0 based leaf indices to probe inclusion proofs at.
var inclusionProofTestIndices = []int64{5, 27, 31, 80, 91}

// consistencyProofTestParams are the intervals
// to test proofs at
var consistencyProofTestParams = []consistencyProofParams{{1, 2}, {2, 3}, {1, 3}, {2, 4}}

// consistencyProofBadTestParams are the intervals to probe for consistency proofs, none of
// these should succeed. Zero is not a valid tree size, nor is -1. 10000000 is outside the
// range we'll reasonably queue (multiple of batch size).
var consistencyProofBadTestParams = []consistencyProofParams{{0, 0}, {-1, 0}, {10000000, 10000000}}

// RunLogIntegration runs a log integration test using the given client and test
// parameters.
func RunLogIntegration(client trillian.TrillianLogClient, params TestParameters) error {
	_ = "STUB: not implemented"
	// Step 1 - Optionally check log starts empty then optionally queue leaves on server
	return nil
}

// Step 2 - Wait for queue to drain when server sequences, give up if it doesn't happen (optional)

// Step 3 - Use get entries to read back what was written, check leaves are correct

// Step 4 - Cross validation between log and memory tree root hashes

// Now that the basic tree has passed validation we can start testing proofs

// Step 5 - Test some inclusion proofs

// Ensure log doesn't serve a proof for a leaf index outside the tree size

// Ensure that log doesn't serve a proof for a valid index at a size outside the tree

// Probe the log at several leaf indices each with a range of tree sizes

// TODO(al): test some inclusion proofs by Merkle hash too.

// Step 6 - Test some consistency proofs

// Make some consistency proof requests that we know should not succeed

// Probe the log between some tree sizes we know are included and check the results against
// the in memory tree. Request proofs at both STH and non STH sizes unless batch size is one,
// when these would be equivalent requests.

// Only do this if the batch size changes when halved

func genEntries(params TestParameters) []*trillian.LogLeaf { _ = "STUB: not implemented"; return nil }

// Shuffle the leaves to see if that breaks things, but record the rand seed
// so we can reproduce failures.

func queueLeaves(client trillian.TrillianLogClient, params TestParameters, entries []*trillian.LogLeaf) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForSequencing(treeID int64, client trillian.TrillianLogClient, params TestParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func readEntries(logID int64, client trillian.TrillianLogClient, params TestParameters) ([]*trillian.LogLeaf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check we got the right number of leaves.

func verifyEntries(written, read []*trillian.LogLeaf) error { _ = "STUB: not implemented"; return nil }

// Check that the MerkleLeafHash field is computed correctly.

// Ensure that the ExtraData in the leaf made it through the roundtrip.
// This was set up when we queued the leaves.

func checkLogRootHashMatches(tree *inmemory.Tree, client trillian.TrillianLogClient, params TestParameters) error {
	_ = "STUB: not implemented"
	// Check the STH against the hash we got from our tree
	return nil
}

// Hash must not be empty and must match the one we built ourselves

// checkInclusionProofLeafOutOfRange requests an inclusion proof beyond the current tree size. This
// should fail
func checkInclusionProofLeafOutOfRange(logID int64, client trillian.TrillianLogClient, params TestParameters) error {
	_ = "STUB: not implemented"
	// Test is a leaf index bigger than the current tree size
	return nil
}

// checkInclusionProofTreeSizeOutOfRange requests an inclusion proof for a leaf within the tree size at
// a tree size larger than the current tree size. This should succeed but with an STH for the current
// tree and an empty proof, because it is a result of skew.
func checkInclusionProofTreeSizeOutOfRange(logID int64, client trillian.TrillianLogClient, params TestParameters) error {
	_ = "STUB: not implemented"
	// Test is an in range leaf index for a tree size that doesn't exist
	return nil
}

// checkInclusionProofsAtIndex obtains and checks proofs at tree sizes from zero up to 2 x the sequencing
// batch size (or number of leaves queued if less). The log should only serve proofs for indices in a tree
// at least as big as the index where STHs where the index is a multiple of the sequencer batch size. All
// proofs returned should match ones computed by the alternate Merkle Tree implementation, which differs
// from what the log uses.
func checkInclusionProofsAtIndex(index int64, logID int64, tree *inmemory.Tree, client trillian.TrillianLogClient, params TestParameters) error {
	_ = "STUB: not implemented"
	return nil
}

// If the index is larger than the tree size we cannot have a valid proof

// Verify inclusion proof.

func checkConsistencyProof(consistParams consistencyProofParams, treeID int64, tree *inmemory.Tree, client trillian.TrillianLogClient, params TestParameters, batchSize int64) error {
	_ = "STUB: not implemented"
	// We expect the proof request to succeed
	return nil
}

// buildMerkleTree returns an in-memory Merkle tree built on the given leaves.
func buildMerkleTree(leaves []*trillian.LogLeaf, params TestParameters) *inmemory.Tree {
	_ = "STUB: not implemented"
	return nil
}

func getLatestSignedLogRoot(client trillian.TrillianLogClient, params TestParameters) (*trillian.GetLatestSignedLogRootResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getRPCDeadlineTime calculates the future time an RPC should expire based on our config
func getRPCDeadlineContext(params TestParameters) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func min(a, b int64) int64 { _ = "STUB: not implemented"; return 0 }
