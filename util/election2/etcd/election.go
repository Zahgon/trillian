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

// Package etcd provides an implementation of master election based on etcd.
package etcd

import (
	"context"

	"github.com/google/trillian/util/election2"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const resignID = "<resign>"

// Election is an implementation of election2.Election based on etcd.
type Election struct {
	resourceID string
	instanceID string
	lockFile   string

	client   *clientv3.Client
	session  *concurrency.Session
	election *concurrency.Election
}

// Await blocks until the instance captures mastership.
func (e *Election) Await(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// WithMastership returns a "mastership context" which remains active until the
// instance stops being the master, or the passed in context is canceled.
func (e *Election) WithMastership(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	// Get a channel for notifications of election status (using the cancelable
	// context so that the monitoring goroutine below and the goroutine started
	// by WithMastership will reliably terminate).
	return *new(context.Context), nil
}

// The revision at which e became the master.

// Not even tried to become the master. Return a canceled context.

// Was the master once, so watch for latest mastership updates.

// Verify that we are still the master, before returning context.

// Mastership has been overtaken, released, or not capturead at all.

// At this point we have observed confirmation that we are the master; start
// a goroutine to monitor for anyone else overtaking us.

// TODO(pavelkalinnikov): conquerorID can be resignID too. Serialize a
// protobuf with all mastership details instead of ID string.

// Resign releases mastership for this instance. The instance can be elected
// again using Await. Idempotent, might be useful to retry if fails.
func (e *Election) Resign(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Trigger Observe callers to see the update, and cancel mastership contexts.
	return nil
}

// Resigning if not master is a no-op.

// Close resigns and permanently stops participating in election. No other
// method should be called after Close.
func (e *Election) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Session's Close revokes the underlying lease, which results in removing
// the election-related keys. This achieves the effect of resignation even if
// the above Resign call failed (e.g. due to ctx cancelation).

// Factory creates Election instances.
type Factory struct {
	client     *clientv3.Client
	instanceID string
	lockDir    string
}

// NewElection creates a specific Election instance.
func (f *Factory) NewElection(ctx context.Context, resourceID string) (election2.Election, error) {
	_ = "STUB: not implemented"
	// TODO(pavelkalinnikov): Re-create the session if it expires.
	// TODO(pavelkalinnikov): Share the same session between Election instances.
	return *new(election2.Election), nil
}
