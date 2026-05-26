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

package election

import (
	"sync"
)

// MasterTracker tracks the current mastership state across multiple IDs.
type MasterTracker struct {
	mu          sync.RWMutex
	masterFor   map[string]bool
	masterCount int
	notify      func(id string, isMaster bool)
}

// NewMasterTracker creates a new MasterTracker instance to track the
// mastership status for the given set of IDs.
func NewMasterTracker(ids []string, notify func(id string, isMaster bool)) *MasterTracker {
	_ = "STUB: not implemented"
	return nil
}

// Set changes the tracked mastership status for the given ID. This method
// should be called exactly once for each state transition.
func (mt *MasterTracker) Set(id string, isMaster bool) { _ = "STUB: not implemented"; return }

// Count returns the number of IDs for which we are currently master.
func (mt *MasterTracker) Count() int { _ = "STUB: not implemented"; return 0 }

// Held returns a (sorted) list of the IDs for which we are currently master.
func (mt *MasterTracker) Held() []string { _ = "STUB: not implemented"; return nil }

// IDs returns a (sorted) list of the IDs that we are currently tracking.
func (mt *MasterTracker) IDs() []string { _ = "STUB: not implemented"; return nil }

// String returns a textual decription of the current mastership status.
func (mt *MasterTracker) String() string { _ = "STUB: not implemented"; return "" }

// HeldInfo produces a textual description of the set of held IDs, compared to
// a complete set of IDs.
func HeldInfo(held []string, ids []string) string { _ = "STUB: not implemented"; return "" }
