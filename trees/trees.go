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

// Package trees contains utility method for retrieving trees and acquiring objects (hashers,
// signers) associated with them.
package trees

import (
	"context"

	"github.com/google/trillian"
	"github.com/google/trillian/storage"
	"google.golang.org/grpc/codes"
)

const traceSpanRoot = "/trillian/trees"

type treeKey struct{}

type accessRule struct {
	// Tree states are accepted if there is a 'true' value for them in this map.
	okStates map[trillian.TreeState]bool
	// Allows the error code to be specified for specific rejected states.
	rejectCodes map[trillian.TreeState]codes.Code
	// Tree types are accepted if there is a 'true' value for them in this map.
	okTypes map[trillian.TreeType]bool
}

// These rules define the permissible combinations of tree state and type
// for each operation type.
var rules = map[OpType]accessRule{
	Unknown: {},
	Admin: {
		okStates: map[trillian.TreeState]bool{
			trillian.TreeState_UNKNOWN_TREE_STATE: true,
			trillian.TreeState_ACTIVE:             true,
			trillian.TreeState_DRAINING:           true,
			trillian.TreeState_FROZEN:             true,
		},
		okTypes: map[trillian.TreeType]bool{
			trillian.TreeType_LOG:            true,
			trillian.TreeType_PREORDERED_LOG: true,
		},
	},
	Query: {
		okStates: map[trillian.TreeState]bool{
			// Have to allow queries on unknown state so storage can get a chance
			// to return ErrTreeNeedsInit.
			trillian.TreeState_UNKNOWN_TREE_STATE: true,
			trillian.TreeState_ACTIVE:             true,
			trillian.TreeState_DRAINING:           true,
			trillian.TreeState_FROZEN:             true,
		},
		okTypes: map[trillian.TreeType]bool{
			trillian.TreeType_LOG:            true,
			trillian.TreeType_PREORDERED_LOG: true,
		},
	},
	QueueLog: {
		okStates: map[trillian.TreeState]bool{
			trillian.TreeState_ACTIVE: true,
		},
		rejectCodes: map[trillian.TreeState]codes.Code{
			trillian.TreeState_DRAINING: codes.PermissionDenied,
			trillian.TreeState_FROZEN:   codes.PermissionDenied,
		},
		okTypes: map[trillian.TreeType]bool{
			trillian.TreeType_LOG:            true,
			trillian.TreeType_PREORDERED_LOG: true,
		},
	},
	SequenceLog: {
		okStates: map[trillian.TreeState]bool{
			trillian.TreeState_ACTIVE:   true,
			trillian.TreeState_DRAINING: true,
		},
		okTypes: map[trillian.TreeType]bool{
			trillian.TreeType_LOG:            true,
			trillian.TreeType_PREORDERED_LOG: true,
		},
		rejectCodes: map[trillian.TreeState]codes.Code{
			trillian.TreeState_FROZEN: codes.PermissionDenied,
		},
	},
}

// NewContext returns a ctx with the given tree.
func NewContext(ctx context.Context, tree *trillian.Tree) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext returns the tree within ctx if present, together with an indication of whether a
// tree was present.
func FromContext(ctx context.Context) (*trillian.Tree, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func validate(o GetOpts, tree *trillian.Tree) error {
	_ = "STUB: not implemented"
	// Do the special case checks first
	return nil
}

// Reject any operation types we don't know about.

// Apply the rule, ensure it allows the tree type and state that we have.

// If we have a status code to use it takes precedence, otherwise it's
// a generic InvalidArgument code.

// GetTree returns the specified tree, either from the ctx (if present) or read from storage.
// The tree will be validated according to GetOpts before returned. Tree state is also considered
// (for example, deleted tree will return NotFound errors).
func GetTree(ctx context.Context, s storage.AdminStorage, treeID int64, opts GetOpts) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No operations should span multiple trees. If a tree is already in the context
// it had better be the one that we want. If the tree comes back from the DB with
// the wrong ID then this checks that too.

func spanFor(ctx context.Context, name string) (context.Context, func()) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
