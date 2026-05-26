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

package storage

import (
	"sync"

	"github.com/google/trillian/monitoring"
)

// NewProviderFunc is the signature of a function which can be registered to
// provide instances of storage providers.
type NewProviderFunc func(monitoring.MetricFactory) (Provider, error)

var (
	spMu     sync.RWMutex
	spByName = make(map[string]NewProviderFunc)
)

// RegisterProvider registers the given storage Provider.
func RegisterProvider(name string, sp NewProviderFunc) error { _ = "STUB: not implemented"; return nil }

// NewProvider returns a new Provider instance of the type specified by name.
func NewProvider(name string, mf monitoring.MetricFactory) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

// Providers returns a slice of all registered storage provider names.
func Providers() []string { _ = "STUB: not implemented"; return nil }

// Provider is an interface which allows Trillian binaries to use different
// storage implementations.
type Provider interface {
	// LogStorage creates and returns a LogStorage implementation.
	LogStorage() LogStorage
	// AdminStorage creates and returns a AdminStorage implementation.
	AdminStorage() AdminStorage

	// Close closes the underlying storage.
	Close() error
}
