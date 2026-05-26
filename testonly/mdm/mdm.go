// Copyright 2019 Google LLC. All Rights Reserved.
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

// Package mdm provides test-only code for checking the merge delay of a
// Trillian log.
package mdm

import (
	"context"
	"sync"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/client"
	"github.com/google/trillian/monitoring"
)

// MergeDelayMonitor submits leaves to a Log and measures merge delay.
type MergeDelayMonitor struct {
	client []*client.LogClient
	logID  int64
	opts   MergeDelayOptions
}

// MergeDelayOptions holds the parameters for a MergeDelayMonitor.
type MergeDelayOptions struct {
	ParallelAdds  int
	LeafSize      int
	NewLeafChance int // percentage
	EmitInterval  time.Duration
	Deadline      time.Duration
	MinMergeDelay time.Duration
	MetricFactory monitoring.MetricFactory
}

// NewMonitor creates a MergeDelayMonitor instance for the given log ID, accessed
// via the cl client.
func NewMonitor(ctx context.Context, logID int64, cl trillian.TrillianLogClient, adminCl trillian.TrillianAdminClient, opts MergeDelayOptions) (*MergeDelayMonitor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Monitor runs merge delay monitoring until its context is cancelled or an error occurs.
func (m *MergeDelayMonitor) Monitor(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MergeDelayMonitor) monitor(ctx context.Context, idx int) error {
	_ = "STUB: not implemented"
	return nil
}

// Always need a new leaf to start with

// Add the leaf data and wait for its inclusion.

// Stats returns the total count of requests and the total elapsed time across
// all invocations.
func (m *MergeDelayMonitor) Stats(newLeaf bool) (uint64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

var (
	metricsOnce    sync.Once
	mergeDelayDist monitoring.Histogram
	newLeafLabel   = map[bool]string{true: "true", false: "false"}
)

func initMetrics(mf monitoring.MetricFactory) { _ = "STUB: not implemented"; return }
