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

// Package prometheus provides a Prometheus-based implementation of the
// MetricFactory abstraction.
package prometheus

import (
	"github.com/google/trillian/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

// MetricFactory allows the creation of Prometheus-based metrics.
type MetricFactory struct {
	// Prefix is an identifier that will be used before local metric names that
	// are reported. It is strongly recommended that this ends with a valid
	// separator (e.g. "_") in order to improve readability; no separator is
	// added by this library.
	Prefix string
}

// NewCounter creates a new Counter object backed by Prometheus.
func (pmf MetricFactory) NewCounter(name, help string, labelNames ...string) monitoring.Counter {
	_ = "STUB: not implemented"
	return *new(monitoring.Counter)
}

// NewGauge creates a new Gauge object backed by Prometheus.
func (pmf MetricFactory) NewGauge(name, help string, labelNames ...string) monitoring.Gauge {
	_ = "STUB: not implemented"
	return *new(monitoring.Gauge)
}

// NewHistogramWithBuckets creates a new Histogram object backed by
// Prometheus and using the supplied bucketing intervals. Note: the
// number of buckets should be kept within reasonable bounds.
func (pmf MetricFactory) NewHistogramWithBuckets(name, help string, buckets []float64, labelNames ...string) monitoring.Histogram {
	_ = "STUB: not implemented"
	return *new(monitoring.Histogram)
}

// NewHistogram creates a new Histogram object backed by Prometheus with
// the supplied buckets.
func (pmf MetricFactory) NewHistogram(name, help string, labelNames ...string) monitoring.Histogram {
	_ = "STUB: not implemented"
	return *new(monitoring.Histogram)
}

func (pmf MetricFactory) newHistogram(name, help string, buckets []float64, labelNames []string) monitoring.Histogram {
	_ = "STUB: not implemented"
	return *new(monitoring.Histogram)
}

// Counter is a wrapper around a Prometheus Counter or CounterVec object.
type Counter struct {
	labelNames []string
	single     prometheus.Counter
	vec        *prometheus.CounterVec
}

// Inc adds 1 to a counter.
func (m *Counter) Inc(labelVals ...string) { _ = "STUB: not implemented"; return }

// Add adds the given amount to a counter.
func (m *Counter) Add(val float64, labelVals ...string) { _ = "STUB: not implemented"; return }

// Value returns the current amount of a counter.
func (m *Counter) Value(labelVals ...string) float64 { _ = "STUB: not implemented"; return 0 }

// Gauge is a wrapper around a Prometheus Gauge or GaugeVec object.
type Gauge struct {
	labelNames []string
	single     prometheus.Gauge
	vec        *prometheus.GaugeVec
}

// Inc adds 1 to a gauge.
func (m *Gauge) Inc(labelVals ...string) { _ = "STUB: not implemented"; return }

// Dec subtracts 1 from a gauge.
func (m *Gauge) Dec(labelVals ...string) { _ = "STUB: not implemented"; return }

// Add adds given value to a gauge.
func (m *Gauge) Add(val float64, labelVals ...string) { _ = "STUB: not implemented"; return }

// Set sets the value of a gauge.
func (m *Gauge) Set(val float64, labelVals ...string) { _ = "STUB: not implemented"; return }

// Value returns the current amount of a gauge.
func (m *Gauge) Value(labelVals ...string) float64 { _ = "STUB: not implemented"; return 0 }

// Histogram is a wrapper around a Prometheus Histogram or HistogramVec object.
type Histogram struct {
	labelNames []string
	single     prometheus.Histogram
	vec        *prometheus.HistogramVec
}

// Observe adds a single observation to the histogram.
func (m *Histogram) Observe(val float64, labelVals ...string) { _ = "STUB: not implemented"; return }

// Info returns the count and sum of observations for the histogram.
func (m *Histogram) Info(labelVals ...string) (uint64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func labelsFor(names, values []string) (prometheus.Labels, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels), nil
}
