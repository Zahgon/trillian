// Copyright 2018 Google LLC. All Rights Reserved.
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

package client

import (
	"context"

	"github.com/google/trillian"
)

// CreateAndInitTree uses the adminClient and logClient to create the tree
// described by req.
// If req describes a LOG tree, then this function will also call the InitLog
// function using logClient.
// Internally, the function will continue to retry failed requests until either
// the tree is created (and if necessary, initialised) successfully, or ctx is
// cancelled.
func CreateAndInitTree(
	ctx context.Context,
	req *trillian.CreateTreeRequest,
	adminClient trillian.TrillianAdminClient,
	logClient trillian.TrillianLogClient) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InitLog initialises a freshly created Log tree.
func InitLog(ctx context.Context, tree *trillian.Tree, logClient trillian.TrillianLogClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for log root to become available.
