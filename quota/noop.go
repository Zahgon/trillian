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

package quota

import (
	"context"

	"k8s.io/klog/v2"
)

// noopManagerName represents the noop quota implementation.
const noopManagerName = "noop"

type noopManager struct{}

func init() {
	if err := RegisterProvider(noopManagerName, func() (Manager, error) {
		return Noop(), nil
	}); err != nil {
		klog.Fatalf("Failed to register %q: %v", noopManagerName, err)
	}
}

// Noop returns a noop implementation of Manager. It allows all requests without restriction.
func Noop() Manager { _ = "STUB: not implemented"; return *new(Manager) }

func (n noopManager) GetTokens(ctx context.Context, numTokens int, specs []Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (n noopManager) PutTokens(ctx context.Context, numTokens int, specs []Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (n noopManager) ResetQuota(ctx context.Context, specs []Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (n noopManager) SetupInitialQuota(ctx context.Context, treeID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNumTokens(numTokens int) error { _ = "STUB: not implemented"; return nil }

func validateSpecs(specs []Spec) error { _ = "STUB: not implemented"; return nil }
