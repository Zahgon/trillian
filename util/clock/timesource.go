// Copyright 2016 Google LLC. All Rights Reserved.
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

// Package clock contains time utilities, and types that allow mocking system
// time in tests.
package clock

import (
	"sync"
	"time"
)

// System is a default TimeSource that provides system time.
var System TimeSource = systemTimeSource{}

// TimeSource can provide the current time, or be replaced by a mock in tests
// to return specific values.
type TimeSource interface {
	// Now returns the current time as seen by this TimeSource.
	Now() time.Time
	// NewTimer creates a timer that fires after the specified duration.
	NewTimer(d time.Duration) Timer
}

// SecondsSince returns the time in seconds elapsed since t until now, as
// measured by the TimeSource.
func SecondsSince(ts TimeSource, t time.Time) float64 { _ = "STUB: not implemented"; return 0 }

// systemTimeSource provides the current system local time.
type systemTimeSource struct{}

// Now returns the true current local time.
func (s systemTimeSource) Now() time.Time {
	_ = "STUB: not implemented"

	// NewTimer returns a real timer.
	return *new(time.Time)
}

func (s systemTimeSource) NewTimer(d time.Duration) Timer {
	_ = "STUB: not implemented"
	return *new(Timer)
}

// FakeTimeSource provides time that can be arbitrarily set. For tests only.
type FakeTimeSource struct {
	mu     sync.RWMutex
	now    time.Time
	timers map[int]*fakeTimer
	nextID int
}

// NewFake creates a FakeTimeSource instance.
func NewFake(t time.Time) *FakeTimeSource { _ = "STUB: not implemented"; return nil }

// Now returns the time value this instance contains.
func (f *FakeTimeSource) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NewTimer returns a fake Timer.
func (f *FakeTimeSource) NewTimer(d time.Duration) Timer {
	_ = "STUB: not implemented"
	return *new(Timer)
}

// unsubscribe removes the Timer with the specified ID if it exists, and
// returns the existence bit.
func (f *FakeTimeSource) unsubscribe(id int) bool { _ = "STUB: not implemented"; return false }

// Set updates the time that this instance will report.
func (f *FakeTimeSource) Set(t time.Time) { _ = "STUB: not implemented"; return }

// PredefinedFake is a TimeSource that returns a predefined set of times
// computed as base time + delays[i]. Delays don't have to be monotonic.
type PredefinedFake struct {
	Base   time.Time
	Delays []time.Duration
	Next   int
}

// Now returns the current time, which depends on how many times this method
// has already been invoked. Must not be called more than len(delays) times.
func (p *PredefinedFake) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NewTimer creates a timer with the specified delay. Not implemented.
func (p *PredefinedFake) NewTimer(d time.Duration) Timer {
	_ = "STUB: not implemented"
	return *new(Timer)
}
