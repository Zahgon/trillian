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

// Package flagsaver provides a simple way to save and restore flag values.
// TODO(RJPercival): Move this to its own GitHub project.
//
// Example:
//
//	func TestFoo(t *testing.T) {
//	  defer flagsaver.Save().Restore()
//	  // Test code that changes flags
//	} // flags are reset to their original values here.
package flagsaver

// Stash holds flag values so that they can be restored at the end of a test.
type Stash struct {
	flags map[string]string
}

// Restore sets all non-hidden flags to the values they had when the Stash was created.
func (s *Stash) Restore() error { _ = "STUB: not implemented"; return nil }

// Save returns a Stash that captures the current value of all non-hidden flags.
func Save() *Stash { _ = "STUB: not implemented"; return nil }

// Exclude the go test related flags. Also exclude log_backtrace_at because
// while it may have an empty value it can't be set to one without an
// error.

// MustRestore calls Restore and exits on failure. It can be used in a defer for
// tests. If Restore fails then the flags may be in an arbitrary
// state that could cause subsequent tests to misbehave.
func (s *Stash) MustRestore() { _ = "STUB: not implemented"; return }
