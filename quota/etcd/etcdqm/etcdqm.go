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

// Package etcdqm contains an etcd-based quota.Manager implementation.
package etcdqm

import (
	"context"

	"github.com/google/trillian/quota"
	"github.com/google/trillian/quota/etcd/storage"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Manager implements a quota manager based on etcd.
type Manager struct {
	qs *storage.QuotaStorage
}

// New returns a new etcd-based quota.Manager.
func New(client *clientv3.Client) *Manager { _ = "STUB: not implemented"; return nil }

// GetTokens implements the quota.Manager API.
func (m *Manager) GetTokens(ctx context.Context, numTokens int, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) peekTokens(ctx context.Context, specs []quota.Spec) (map[quota.Spec]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PutTokens implements the quota.Manager API.
func (m *Manager) PutTokens(ctx context.Context, numTokens int, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ResetQuota implements the quota.Manager API.
func (m *Manager) ResetQuota(ctx context.Context, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func configNames(specs []quota.Spec) []string { _ = "STUB: not implemented"; return nil }

func configName(spec quota.Spec) string { _ = "STUB: not implemented"; return "" }
