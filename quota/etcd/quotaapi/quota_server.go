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

// Package quotaapi provides a Quota admin server implementation.
package quotaapi

import (
	"context"

	"github.com/google/trillian/quota/etcd/quotapb"
	"github.com/google/trillian/quota/etcd/storage"
	"github.com/google/trillian/quota/etcd/storagepb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server is a quotapb.QuotaServer implementation backed by etcd.
type Server struct {
	qs *storage.QuotaStorage
}

// NewServer returns a new Server instance backed by client.
func NewServer(client *clientv3.Client) *Server { _ = "STUB: not implemented"; return nil }

// CreateConfig implements quotapb.QuotaServer.CreateConfig.
func (s *Server) CreateConfig(ctx context.Context, req *quotapb.CreateConfigRequest) (*quotapb.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* reset */

// DeleteConfig implements quotapb.QuotaServer.DeleteConfig.
func (s *Server) DeleteConfig(ctx context.Context, req *quotapb.DeleteConfigRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* reset */

// GetConfig implements quotapb.QuotaServer.GetConfig.
func (s *Server) GetConfig(ctx context.Context, req *quotapb.GetConfigRequest) (*quotapb.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getConfig finds the Config named "name" on "cfgs", converts it to API and returns it.
// If the config cannot be found an error with code "code" is returned.
func (s *Server) getConfig(ctx context.Context, name string, cfgs *storagepb.Configs, code codes.Code) (*quotapb.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListConfigs implements quotapb.QuotaServer.ListConfigs.
func (s *Server) ListConfigs(ctx context.Context, req *quotapb.ListConfigsRequest) (*quotapb.ListConfigsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Peek token counts for FULL view

func listMatches(nfs []nameFilter, cfg *storagepb.Config) bool {
	_ = "STUB: not implemented"
	return false
}

// Match all

func listView(view quotapb.ListConfigsRequest_ListView, src *storagepb.Config) *quotapb.Config {
	_ = "STUB: not implemented"
	return nil
}

// UpdateConfig implements quotapb.QuotaServer.UpdateConfig.
func (s *Server) UpdateConfig(ctx context.Context, req *quotapb.UpdateConfigRequest) (*quotapb.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For convenience, reset-only requests are allowed.

func findByName(name string, cfgs *storagepb.Configs) (*storagepb.Config, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func validateName(name string) error { _ = "STUB: not implemented"; return nil }
