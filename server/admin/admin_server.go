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

package admin

import (
	"context"

	"github.com/google/trillian"
	"github.com/google/trillian/extension"
	"google.golang.org/genproto/protobuf/field_mask"
)

// Server is an implementation of trillian.TrillianAdminServer.
type Server struct {
	registry         extension.Registry
	allowedTreeTypes []trillian.TreeType
}

// New returns a trillian.TrillianAdminServer implementation.
// registry is the extension.Registry used by the Server.
// allowedTreeTypes defines which tree types may be created through this server,
// with nil meaning unrestricted.
func New(registry extension.Registry, allowedTreeTypes []trillian.TreeType) *Server {
	_ = "STUB: not implemented"
	return nil
}

// IsHealthy returns nil if the server is healthy, error otherwise.
// TODO(Martin2112): This method (and the one in the log server) should probably have ctx as a param
func (s *Server) IsHealthy() error { _ = "STUB: not implemented"; return nil }

// ListTrees implements trillian.TrillianAdminServer.ListTrees.
func (s *Server) ListTrees(ctx context.Context, req *trillian.ListTreesRequest) (*trillian.ListTreesResponse, error) {
	_ = "STUB: not implemented"
	// TODO(codingllama): This needs access control
	return nil, nil
}

// GetTree implements trillian.TrillianAdminServer.GetTree.
func (s *Server) GetTree(ctx context.Context, req *trillian.GetTreeRequest) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateTree implements trillian.TrillianAdminServer.CreateTree.
func (s *Server) CreateTree(ctx context.Context, req *trillian.CreateTreeRequest) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clear generated fields, storage must set those

func (s *Server) validateAllowedTreeType(tt trillian.TreeType) error {
	_ = "STUB: not implemented"
	return nil
}

// All types OK

// UpdateTree implements trillian.TrillianAdminServer.UpdateTree.
func (s *Server) UpdateTree(ctx context.Context, req *trillian.UpdateTreeRequest) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply the mask to a couple of empty trees just to check that the paths are correct.

// Should never happen (famous last words).

func applyUpdateMask(from, to *trillian.Tree, mask *field_mask.FieldMask) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteTree implements trillian.TrillianAdminServer.DeleteTree.
func (s *Server) DeleteTree(ctx context.Context, req *trillian.DeleteTreeRequest) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UndeleteTree implements trillian.TrillianAdminServer.UndeleteTree.
func (s *Server) UndeleteTree(ctx context.Context, req *trillian.UndeleteTreeRequest) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
