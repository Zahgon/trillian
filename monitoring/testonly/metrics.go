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

package testonly

import (
	"testing"

	"github.com/google/trillian/monitoring"
)

// TestCounter runs a test on a Counter produced from the provided MetricFactory.
func TestCounter(t *testing.T, factory monitoring.MetricFactory) { _ = "STUB: not implemented"; return }

// Use an invalid number of labels.

// Check that the value hasn't changed.

// Use a different set of label values
// Metrics with different valued label values, are distinct
// This test is only applicable when a Metric has labels

// Increment counter using this different set of label values

// Counter with original set of label values should be unchanged

// TestGauge runs a test on a Gauge produced from the provided MetricFactory.
func TestGauge(t *testing.T, factory monitoring.MetricFactory) { _ = "STUB: not implemented"; return }

// Use an invalid number of labels.

// Ask for an invalid number of labels.

// Check that the value hasn't changed.

// Use a different set of label values
// Metrics with different valued label values, are distinct
// This test is only applicable when a Metric has labels

// Set gauge using this different set of label values

// Gauge with original set of label values should be unchanged

// TestHistogram runs a test on a Histogram produced from the provided MetricFactory.
func TestHistogram(t *testing.T, factory monitoring.MetricFactory) {
	_ = "STUB: not implemented"
	return
}

// Use an invalid number of labels.

// Check that the histogram hasn't changed.

// Use a different set of label values
// Metrics with different valued label values, are distinct
// This test is only applicable when a Metric has labels

// Observe histogram using this different set of label values

// Histogram with original set of label values should be unchanged
