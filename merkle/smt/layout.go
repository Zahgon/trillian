// Copyright 2021 Google LLC. All Rights Reserved.
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

// Layout defines the mapping between tree nodes and tiles.
//
// The tree is divided into horizontal tiers of the heights specified in the
// constructor. The tiers are counted down from the tree root. Nodes located at
// the tier borders are roots of the corresponding tiles. A tile includes all
// the nodes strictly below the tile root, within a single tier.
type Layout struct {
	depths []int
}

// NewLayout creates a tree layout based on the passed-in tier heights.
func NewLayout(heights []uint) Layout { _ = "STUB: not implemented"; return *new(Layout) }

// Locate returns the tier info that nodes at the given depth belong to. The
// first returned value is the depth of the root node, for any tile in this
// tier. The second value is the height of the tiles.
//
// If depth is 0, or beyond the tree height, then the result is undefined.
func (l Layout) Locate(depth uint) (uint, uint) {
	_ = "STUB: not implemented"
	// The root belongs to its own "pseudo" tile.
	return 0, 0
}

// 1 <= i <= len(depths)

// Anything below the leaves does not exist.
