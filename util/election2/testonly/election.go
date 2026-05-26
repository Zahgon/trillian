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

// Package testonly contains an Election implementation for testing.
package testonly

import (
	"context"
	"sync"

	"github.com/google/trillian/util/election2"
)

// Factory allows creating Election instances for testing.
var Factory factory

// Election implements election2.Election interface for testing.
type Election struct {
	isMaster bool
	revision int
	mu       sync.Mutex
	cond     *sync.Cond
}

// NewElection returns a new initialized Election for testing.
func NewElection() *Election { _ = "STUB: not implemented"; return nil }

// update updates this instance's mastership status. Must be called under lock.
func (e *Election) update(isMaster bool) { _ = "STUB: not implemented"; return }

// Await sets this instance to be the master. It always succeeds. To imitate
// errors and/or blocking behavior use the Decorator type.
func (e *Election) Await(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// WithMastership returns mastership context, which gets canceled if / when
// this instance is not / stops being the master.
func (e *Election) WithMastership(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Notify e.cond on ctx cancelation.

// Watch mastership and the context in the background.

// Resign resets mastership.
func (e *Election) Resign(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close resets mastership permanently.
func (e *Election) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func watchContext(ctx context.Context, l sync.Locker, cond *sync.Cond) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// Avoid racing with cond waiters on ctx status.

// factory allows creating Election instances.
type factory struct{}

// NewElection creates a new Election instance.
// TODO(pavelkalinnikov): Use resourceID in tests with multiple resources.
func (f factory) NewElection(ctx context.Context, resourceID string) (election2.Election, error) {
	_ = "STUB: not implemented"
	return *new(election2.Election), nil
}
