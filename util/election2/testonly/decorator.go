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
	"sync"

	"github.com/google/trillian/util/election2"
)

// Errs contains errors to be returned by each of Election methods.
type Errs struct {
	Await          error
	WithMastership error
	Resign         error
	Close          error
}

// Decorator is an election2.Election decorator injecting errors, for testing.
type Decorator struct {
	e     election2.Election
	errs  Errs
	block bool
	mu    sync.Mutex
	cond  *sync.Cond
}

// NewDecorator returns a Decorator wrapping the passed in Election object.
func NewDecorator(e election2.Election) *Decorator { _ = "STUB: not implemented"; return nil }

// Update updates errors returned by interface methods.
func (d *Decorator) Update(errs Errs) { _ = "STUB: not implemented"; return }

// BlockAwait enables or disables Await method blocking.
func (d *Decorator) BlockAwait(block bool) { _ = "STUB: not implemented"; return }

// Await blocks until the instance captures mastership.
func (d *Decorator) Await(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// WithMastership returns a mastership context.
func (d *Decorator) WithMastership(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Resign releases mastership for this instance.
func (d *Decorator) Resign(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close permanently stops participating in election.
func (d *Decorator) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
