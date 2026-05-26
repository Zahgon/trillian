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
	"github.com/google/trillian/storage/tree"
	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/proof"
)

// nodeReader provides read-only access to the tree nodes.
type nodeReader interface {
	// GetMerkleNodes returns tree nodes by their IDs, in the requested order.
	GetMerkleNodes(ctx context.Context, ids []compact.NodeID) ([]tree.Node, error)
}

// fetchNodesAndBuildProof is used by both inclusion and consistency proofs. It fetches the nodes
// from storage and converts them into the proof proto that will be returned to the client.
// This includes rehashing where necessary to serve proofs for tree sizes between stored tree
// revisions. This code only relies on the nodeReader interface so can be tested without
// a complete storage implementation.
func fetchNodesAndBuildProof(ctx context.Context, nr nodeReader, hasher compact.HashFn, leafIndex uint64, pn proof.Nodes) (*trillian.Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetchNodes obtains the nodes denoted by the given NodeFetch structs, and
// returns them after some validation checks.
func fetchNodes(ctx context.Context, nr nodeReader, ids []compact.NodeID) ([]tree.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Additional check that the correct node was returned.
