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

package testonly

import (
	"context"
	"testing"
	"time"

	"github.com/google/trillian/util/election2"
)

// Tests is the full list of available Election tests.
// TODO(pavelkalinnikov): Add tests for unexpected mastership loss.
var Tests = []NamedTest{
	{Name: "RunElectionAwait", Run: runElectionAwait},
	{Name: "RunElectionWithMastership", Run: runElectionWithMastership},
	{Name: "RunElectionResign", Run: runElectionResign},
	{Name: "RunElectionClose", Run: runElectionClose},
	{Name: "RunElectionLoop", Run: runElectionLoop},
}

// NamedTest is a test function paired with its string name.
type NamedTest struct {
	Name string
	Run  func(t *testing.T, f election2.Factory)
}

// CheckNotDone ensures that the context is not done for some time.
func CheckNotDone(ctx context.Context, t *testing.T) { _ = "STUB: not implemented"; return }

// CheckDone ensures that the context is done within the specified duration.
func CheckDone(ctx context.Context, t *testing.T, wait time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Ok.

// runElectionAwait tests the Await call with different pre-conditions.
func runElectionAwait(t *testing.T, f election2.Factory) { _ = "STUB: not implemented"; return }

// runElectionWithMastership tests the WithMastership call.
func runElectionWithMastership(t *testing.T, f election2.Factory) {
	_ = "STUB: not implemented"
	return
}

// runElectionResign tests the Resign call.
func runElectionResign(t *testing.T, f election2.Factory) { _ = "STUB: not implemented"; return }

// runElectionClose tests the Close call.
func runElectionClose(t *testing.T, f election2.Factory) { _ = "STUB: not implemented"; return }

// TODO(pavelkalinnikov): reinstate when not flaky.
// {desc: "cancel", cancel: true, wantErr: nil},

// runElectionLoop runs a typical mastership loop.
func runElectionLoop(t *testing.T, f election2.Factory) { _ = "STUB: not implemented"; return }

// Do some work as master.

// The mastership context should close.
