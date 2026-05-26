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
	"github.com/google/trillian"
	"github.com/transparency-dev/merkle"
)

func validateGetInclusionProofRequest(req *trillian.GetInclusionProofRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGetInclusionProofByHashRequest(req *trillian.GetInclusionProofByHashRequest, hasher merkle.LogHasher) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGetLeavesByRangeRequest(req *trillian.GetLeavesByRangeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGetConsistencyProofRequest(req *trillian.GetConsistencyProofRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGetEntryAndProofRequest(req *trillian.GetEntryAndProofRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAddSequencedLeavesRequest(req *trillian.AddSequencedLeavesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: Not empty, as verified by validateLogLeaves.

func validateLogLeaves(leaves []*trillian.LogLeaf, errPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLogLeaf(leaf *trillian.LogLeaf, errPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLeafHash(hash []byte, hasher merkle.LogHasher) error {
	_ = "STUB: not implemented"
	return nil
}
