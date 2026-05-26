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

// A binary to produce LaTeX documents representing Merkle trees.
// The generated document should be fed into xelatex, and the Forest package
// must be available.
//
// Usage: go run main.go | xelatex
// This should generate a PDF file called treetek.pdf containing a drawing of
// the tree.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/bits"
	"strings"

	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/proof"
)

const (
	preamble = `
% Hash-tree
% Author: treetex
\documentclass[convert]{standalone}
\usepackage[dvipsnames]{xcolor}
\usepackage{forest}


\begin{document}

% Change colours here:
\definecolor{proof}{rgb}{1,0.5,0.5}
\definecolor{proof_ephemeral}{rgb}{1,0.7,0.7}
\definecolor{perfect}{rgb}{1,0.9,0.5}
\definecolor{target}{rgb}{0.5,0.5,0.9}
\definecolor{target_path}{rgb}{0.7,0.7,0.9}
\definecolor{mega}{rgb}{0.9,0.9,0.9}
\definecolor{target0}{rgb}{0.1,0.9,0.1}
\definecolor{target1}{rgb}{0.1,0.1,0.9}
\definecolor{target2}{rgb}{0.9,0.1,0.9}
\definecolor{range0}{rgb}{0.3,0.9,0.3}
\definecolor{range1}{rgb}{0.3,0.3,0.9}
\definecolor{range2}{rgb}{0.9,0.3,0.9}

\forestset{
	% This defines a new "edge" style for drawing the perfect subtrees.
	% Rather than simply drawing a line representing an edge, this draws a
	% triangle between the labelled anchors on the given nodes.
	% See "Anchors" section in the Forest manual for more details:
	%  http://mirrors.ibiblio.org/CTAN/graphics/pgf/contrib/forest/forest-doc.pdf
	perfect/.style={edge path={%
		\noexpand\path[fill=mega, \forestoption{edge}]
				(.parent first)--(!u.children)--(.parent last)--cycle
				\forestoption{edge label};
		}
	},
}
\begin{forest}
`

	postfix = `\end{forest}
\end{document}
`

	// Maximum number of ranges to allow.
	maxRanges = 3
)

var (
	treeSize   = flag.Uint64("tree_size", 23, "Size of tree to produce")
	leafData   = flag.String("leaf_data", "", "Comma separated list of leaf data text (setting this overrides --tree_size")
	nodeFormat = flag.String("node_format", "address", "Format for internal node text, one of: address, hash")
	inclusion  = flag.Int64("inclusion", -1, "Leaf index to show inclusion proof")
	megaMode   = flag.Uint("megamode_threshold", 4, "Treat perfect trees larger than this many layers as a single entity")
	ranges     = flag.String("ranges", "", "Comma-separated Open-Closed ranges of the form L:R")

	attrPerfectRoot   = flag.String("attr_perfect_root", "", "Latex treatment for perfect root nodes (e.g. 'line width=3pt')")
	attrEphemeralNode = flag.String("attr_ephemeral_node", "draw, dotted", "Latex treatment for ephemeral nodes")

	// nInfo holds nodeInfo data for the tree.
	nInfo = make(map[compact.NodeID]nodeInfo)
)

// nodeInfo represents the style to be applied to a tree node.
// TODO(al): separate out leafdata bits from here.
type nodeInfo struct {
	proof            bool
	incPath          bool
	target           bool
	perfectRoot      bool
	ephemeral        bool
	leaf             bool
	dataRangeIndices []int
	rangeIndices     []int
}

type nodeTextFunc func(id compact.NodeID) string

// String returns a string containing Forest attributes suitable for
// rendering the node, given its type.
func (n nodeInfo) String() string { _ = "STUB: not implemented"; return "" }

// Figure out which colour to fill with:

// Otherwise, we need to be a bit cleverer, and use the shading feature.

// modifyNodeInfo applies f to the nodeInfo associated with node id.
func modifyNodeInfo(id compact.NodeID, f func(*nodeInfo)) {
	_ = "STUB: not implemented"
	// Note: Returns an empty nodeInfo if id is not found.
	return
}

// perfectMega renders a large perfect subtree as a single entity.
func perfectMega(prefix string, id compact.NodeID) { _ = "STUB: not implemented"; return }

// Create some hidden nodes to preseve the tier spacings:

// perfect renders a perfect subtree.
func perfect(prefix string, id compact.NodeID, nodeText, dataText nodeTextFunc) {
	_ = "STUB: not implemented"
	return
}

// drawLeaf emits TeX code to render a leaf.
func drawLeaf(prefix string, index uint64, leafText, dataText nodeTextFunc) {
	_ = "STUB: not implemented"
	return
}

// First render the leaf node of the Merkle tree.

// and then a child-node representing the leaf data itself:

// proofs don't include leafdata (just the leaf hash above)
// draw the target leaf darker if necessary.

// openInnerNode renders TeX code to open an internal node.
// The caller may emit any number of child nodes before calling the returned
// func to close the node.
// Returns a func to be called to close the node.
func openInnerNode(prefix string, id compact.NodeID, nodeText nodeTextFunc) func() {
	_ = "STUB: not implemented"
	return nil
}

// perfectInner renders the nodes of a perfect internal subtree.
func perfectInner(prefix string, id compact.NodeID, top bool, nodeText nodeTextFunc, dataText nodeTextFunc) {
	_ = "STUB: not implemented"
	return
}

// renderTree renders a tree node and recurses if necessary.
func renderTree(prefix string, size uint64, nodeText, dataText nodeTextFunc) {
	_ = "STUB: not implemented"
	// Get root IDs of all perfect subtrees.
	return
}

// parseRanges parses and validates a string of comma-separates open-closed
// ranges of the form L:R.
// Returns the parsed ranges, or an error if there's a problem.
func parseRanges(ranges string, treeSize uint64) ([][2]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// modifyRangeNodeInfo sets style info for nodes affected by ranges.
// This includes leaves and perfect subtree roots.
// TODO(al): Figure out what, if anything, to do to make this show ranges
// which are inside the perfect meganodes.
func modifyRangeNodeInfo() error { _ = "STUB: not implemented"; return nil }

// Set leaves:

var dataFormat = func(id compact.NodeID) string {
	return fmt.Sprintf("{$leaf_{%d}$}", id.Index)
}

var nodeFormats = map[string]nodeTextFunc{
	"address": func(id compact.NodeID) string {
		return fmt.Sprintf("%d.%d", id.Level, id.Index)
	},
	"hash": func(id compact.NodeID) string {
		// For "hash" format node text, levels >=1 need a different format
		// [H=H(childL|childR)]from the base level (H=H(leafN)].
		if id.Level >= 1 {
			childLevel := id.Level - 1
			leftChild := id.Index * 2
			return fmt.Sprintf("{$H_{%d.%d} =$ \\\\ $H(H_{%d.%d} || H_{%d.%d})$}", id.Level, id.Index, childLevel, leftChild, childLevel, leftChild+1)
		}
		return fmt.Sprintf("{$H_{%d.%d} =$ \\\\ $H(leaf_{%[2]d})$}", id.Level, id.Index)
	},
}

// Whee - here we go!
func main() {
	// TODO(al): check flag validity.
	flag.Parse()
	height := uint(bits.Len64(*treeSize-1)) + 1

	innerNodeText := nodeFormats[*nodeFormat]
	if innerNodeText == nil {
		log.Fatalf("unknown --node_format %s", *nodeFormat)
	}

	nodeText := innerNodeText

	if len(*leafData) > 0 {
		leaves := strings.Split(*leafData, ",")
		*treeSize = uint64(len(leaves))
		log.Printf("Overriding treeSize to %d since --leaf_data was set", *treeSize)
		dataFormat = func(id compact.NodeID) string {
			return leaves[id.Index]
		}
	}

	if *inclusion > 0 {
		leafID := compact.NewNodeID(0, uint64(*inclusion))
		modifyNodeInfo(leafID, func(n *nodeInfo) { n.incPath = true })
		nodes, err := proof.Inclusion(uint64(*inclusion), *treeSize)
		if err != nil {
			log.Fatalf("Failed to calculate inclusion proof addresses: %s", err)
		}
		_, begin, end := nodes.Ephem()
		for i, id := range nodes.IDs {
			// Skip children of the ephemeral node.
			if i >= begin && i < end && begin+1 < end {
				continue
			}
			modifyNodeInfo(id, func(n *nodeInfo) { n.proof = true })
		}
		// If the ephemeral node exists in the proof, make it a parent of the biggest subtree.
		if begin+1 < end {
			modifyNodeInfo(nodes.IDs[end-1].Parent(), func(n *nodeInfo) { n.proof = true })
		}

		for id := leafID; id.Level < height; id = id.Parent() {
			modifyNodeInfo(id, func(n *nodeInfo) { n.incPath = true })
		}
	}

	if len(*ranges) > 0 {
		if err := modifyRangeNodeInfo(); err != nil {
			log.Fatalf("Failed to modify range node styles: %s", err)
		}
	}

	// TODO(al): structify this into a util, and add ability to output to an
	// arbitrary stream.
	fmt.Print(preamble)
	renderTree("", *treeSize, nodeText, dataFormat)
	fmt.Print(postfix)
}
