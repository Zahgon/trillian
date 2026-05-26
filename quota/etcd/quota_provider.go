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

// Package etcd provides the configuration and initialization of the etcd
// quota manager.
package etcd

import (
	"flag"

	"github.com/google/trillian/quota"
	"github.com/google/trillian/quota/cacheqm"
	"k8s.io/klog/v2"
)

// QuotaManagerName identifies the etcd quota implementation.
const QuotaManagerName = "etcd"

var (
	// Servers is a flag containing the address(es) of etcd servers
	Servers = flag.String("etcd_servers", "", "A comma-separated list of etcd servers; no etcd registration if empty")
	// TODO(Martin2112): suggested renaming these to etc_... to avoid clashes, but will it break existing deploys?
	quotaMinBatchSize = flag.Int("quota_min_batch_size", cacheqm.DefaultMinBatchSize, "Minimum number of tokens to request from the quota system. "+
		"Zero or lower means batching is disabled. Applicable for etcd quotas.")
	quotaMaxCacheEntries = flag.Int("quota_max_cache_entries", cacheqm.DefaultMaxCacheEntries, "Max number of quota specs in the quota cache. "+
		"Zero or lower means batching is disabled. Applicable for etcd quotas.")
)

func init() {
	if err := quota.RegisterProvider(QuotaManagerName, newEtcdQuotaManager); err != nil {
		klog.Fatalf("Failed to register quota manager %v: %v", QuotaManagerName, err)
	}
}

func newEtcdQuotaManager() (quota.Manager, error) {
	_ = "STUB: not implemented"
	return *new(quota.Manager), nil
}
