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

package monitoring

import (
	"sync"
)

// InertMetricFactory creates inert metrics for testing.
type InertMetricFactory struct{}

// NewCounter creates a new inert Counter.
func (imf InertMetricFactory) NewCounter(name, help string, labelNames ...string) Counter {
	_ = "STUB: not implemented"
	return *new(Counter)
}

// NewGauge creates a new inert Gauge.
func (imf InertMetricFactory) NewGauge(name, help string, labelNames ...string) Gauge {
	_ = "STUB: not implemented"
	return *new(Gauge)
}

// NewHistogram creates a new inert Histogram.
func (imf InertMetricFactory) NewHistogram(name, help string, labelNames ...string) Histogram {
	_ = "STUB: not implemented"
	return *new(Histogram)
}

// NewHistogramWithBuckets creates a new inert Histogram with supplied buckets.
// The buckets are not actually used.
func (imf InertMetricFactory) NewHistogramWithBuckets(name, help string, _ []float64, labelNames ...string) Histogram {
	_ = "STUB: not implemented"
	return *new(Histogram)
}

// InertFloat is an internal-only implementation of both the Counter and Gauge interfaces.
type InertFloat struct {
	labelCount int
	mu         sync.Mutex
	vals       map[string]float64
}

// Inc adds 1 to the value.
func (m *InertFloat) Inc(labelVals ...string) { _ = "STUB: not implemented"; return }

// Dec subtracts 1 from the value.
func (m *InertFloat) Dec(labelVals ...string) { _ = "STUB: not implemented"; return }

// Add adds the given amount to the value.
func (m *InertFloat) Add(val float64, labelVals ...string) { _ = "STUB: not implemented"; return }

// Set sets the value.
func (m *InertFloat) Set(val float64, labelVals ...string) { _ = "STUB: not implemented"; return }

// Value returns the current value.
func (m *InertFloat) Value(labelVals ...string) float64 { _ = "STUB: not implemented"; return 0 }

// InertDistribution is an internal-only implementation of the Distribution interface.
type InertDistribution struct {
	labelCount int
	mu         sync.Mutex
	counts     map[string]uint64
	sums       map[string]float64
}

// Observe adds a single observation to the distribution.
func (m *InertDistribution) Observe(val float64, labelVals ...string) {
	_ = "STUB: not implemented"
	return
}

// Info returns count, sum for the distribution.
func (m *InertDistribution) Info(labelVals ...string) (uint64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func keyForLabels(labelVals []string, count int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
