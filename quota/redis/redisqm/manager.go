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

// Package redisqm defines a Redis-based quota.Manager implementation.
package redisqm

import (
	"context"

	"github.com/google/trillian/quota"
	"github.com/google/trillian/quota/redis/redistb"
)

// ParameterFunc is a function that should return a token bucket's parameters
// for a given quota specification.
type ParameterFunc func(spec quota.Spec) (capacity int, rate float64)

// ManagerOptions holds the parameters for a Manager.
type ManagerOptions struct {
	// Parameters should return the parameters for a given quota.Spec. This
	// value must not be nil.
	Parameters ParameterFunc

	// Prefix is a static prefix to apply to all Redis keys; this is useful
	// if running on a multi-tenant Redis cluster.
	Prefix string
}

// Manager implements the quota.Manager interface backed by a Redis-based token
// bucket implementation.
type Manager struct {
	tb   *redistb.TokenBucket
	opts ManagerOptions
}

var _ quota.Manager = &Manager{}

// RedisClient is an interface that encompasses the various methods used by
// this quota.Manager, and allows selecting among different Redis client
// implementations (e.g. regular Redis, Redis Cluster, sharded, etc.)
type RedisClient interface {
	// Everything required by the redistb.RedisClient interface
	redistb.RedisClient
}

// New returns a new Redis-based quota.Manager.
func New(client RedisClient, opts ManagerOptions) *Manager { _ = "STUB: not implemented"; return nil }

// GetTokens implements the quota.Manager API.
func (m *Manager) GetTokens(ctx context.Context, numTokens int, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) getTokensSingle(ctx context.Context, numTokens int, spec quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// If we get back `MaxTokens` from our parameters call, this indicates
// that there's no actual limit. We don't need to do anything to "get"
// them; just ignore.

// PutTokens implements the quota.Manager API.
func (m *Manager) PutTokens(ctx context.Context, numTokens int, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	// Putting tokens into a time-based quota doesn't mean anything (since
	// tokens are replenished at the moment they're requested) and since
	// that's the only supported mechanism for this package currently, do
	// nothing.
	return nil
}

// ResetQuota implements the quota.Manager API.
//
// This function will reset every quota and return the first error encountered,
// if any, but will continue trying to reset every quota even if an error is
// encountered.
func (m *Manager) ResetQuota(ctx context.Context, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// Load attempts to load Redis scripts used by the Manager into the Redis
// cluster.
//
// A Manager will operate successfully if this method is not called or fails,
// but a successful Load will reduce bandwidth to/from the Redis cluster
// substantially.
func (m *Manager) Load(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func specNames(prefix string, specs []quota.Spec) []string { _ = "STUB: not implemented"; return nil }

func specName(prefix string, spec quota.Spec) string { _ = "STUB: not implemented"; return "" }
