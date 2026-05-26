// Copyright 2019 Google LLC. All Rights Reserved.
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

package smt

import (
	"github.com/google/trillian/merkle/smt/node"
)

// Node represents a sparse Merkle tree node.
type Node struct {
	ID   node.ID
	Hash []byte
}

// NodesRow contains nodes at the same tree level sorted by ID from left to
// right. The IDs of the nodes are unique.
//
// TODO(pavelkalinnikov): Hide nodes so that only this package can modify them.
type NodesRow []Node

// NewNodesRow creates a NodesRow from the given list of nodes. The nodes are
// reordered in-place if not already sorted.
func NewNodesRow(nodes []Node) (NodesRow, error) {
	_ = "STUB: not implemented"
	return *new(NodesRow), nil
}

// inSubtree returns whether all the nodes in this row are strictly under the
// node with the given ID. Panics if the row is empty.
func (n NodesRow) inSubtree(root node.ID) bool { _ = "STUB: not implemented"; return false }

// Note: It is enough to check only the first and the last node ID because
// the list is sorted.

// Prepare sorts the nodes slice for it to be usable by HStar3 algorithm and
// the sparse Merkle tree Writer. It also verifies that the nodes are placed at
// the required depth, and there are no duplicate IDs.
//
// TODO(pavelkalinnikov): Make this algorithm independent of Node type.
func Prepare(nodes []Node, depth uint) error { _ = "STUB: not implemented"; return nil }

// compareHorizontal compares relative position of two node IDs at the same
// tree level. Returns -1 if the first node is to the left from the second one,
// 1 if the first node is to the right, and 0 if IDs are the same. The result
// is undefined if nodes are not at the same level.
func compareHorizontal(a, b node.ID) int { _ = "STUB: not implemented"; return 0 }
