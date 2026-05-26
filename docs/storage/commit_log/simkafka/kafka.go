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

// Package simkafka is a toy simulation of a Kafka commit log.
package simkafka

import (
	"sync"
)

type commitLog []string

const showCount = 10

func (c commitLog) String() string { _ = "STUB: not implemented"; return "" }

var (
	mu     sync.RWMutex
	topics = make(map[string]commitLog)
)

// Status reports the current status of the simulated Kafka instance
func Status() string { _ = "STUB: not implemented"; return "" }

// Read returns a value for a topic at a specific offset.
func Read(which string, offset int) string { _ = "STUB: not implemented"; return "" }

// ReadLast returns the latest value for a topic, and its offset
func ReadLast(which string) (string, int) { _ = "STUB: not implemented"; return "", 0 }

// ReadMultiple reads values for a topic, starting at the given offset, up to the
// given maximum number of results.
func ReadMultiple(which string, offset, max int) []string { _ = "STUB: not implemented"; return nil }

// Append adds a value to the end of a topic, and returns the offset of the added value.
func Append(which string, what string) int { _ = "STUB: not implemented"; return 0 }
