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

package storage

import (
	"context"

	"github.com/google/trillian"
)

// ValidateTreeForCreation returns nil if tree is valid for insertion, error
// otherwise.
// See the documentation on trillian.Tree for reference on which values are
// valid.
func ValidateTreeForCreation(ctx context.Context, tree *trillian.Tree) error {
	_ = "STUB: not implemented"
	return nil
}

// validateTreeTypeUpdate returns nil iff oldTree.TreeType can be updated to
// newTree.TreeType. The tree type is changeable only if the Tree is and
// remains in the FROZEN state.
// At the moment only PREORDERED_LOG->LOG type transition is permitted.
func validateTreeTypeUpdate(oldTree, newTree *trillian.Tree) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateTreeForUpdate returns nil if newTree is valid for update, error
// otherwise.
// The newTree is compared to the storedTree to determine if readonly fields
// have been changed. It's assumed that storage-generated fields, such as
// update_time, have not yet changed when this method is called.
// See the documentation on trillian.Tree for reference on which fields may be
// changed and what is considered valid for each of them.
func ValidateTreeForUpdate(ctx context.Context, storedTree, newTree *trillian.Tree) error {
	_ = "STUB: not implemented"
	// Check that readonly fields didn't change
	return nil
}

func validateMutableTreeFields(ctx context.Context, tree *trillian.Tree) error {
	_ = "STUB: not implemented"
	return nil
}

// Implementations may vary, so let's assume storage_settings is mutable.
// Other than checking that it's a valid Any there isn't much to do at this layer, though.
