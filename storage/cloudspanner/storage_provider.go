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

package cloudspanner

import (
	"flag"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	"google.golang.org/api/option"
)

var (
	csURI                                = flag.String("cloudspanner_uri", "", "Connection URI for CloudSpanner database")
	csNumChannels                        = flag.Int("cloudspanner_num_channels", 0, "Number of gRPC channels to use to talk to CloudSpanner.")
	_                                    = flag.Uint64("cloudspanner_max_open_sessions", 0, "DEPRECATED. Spanner no longer supports this feature.")
	_                                    = flag.Uint64("cloudspanner_min_open_sessions", 0, "DEPRECATED. Spanner no longer supports this feature.")
	_                                    = flag.Uint64("cloudspanner_max_idle_sessions", 0, "DEPRECATED. Spanner no longer supports this feature.")
	_                                    = flag.Float64("cloudspanner_write_sessions", 0, "DEPRECATED. This flag is unused and will be removed in the future. Fraction of write capable sessions to maintain.")
	_                                    = flag.Int("cloudspanner_num_healthcheckers", 0, "DEPRECATED. Spanner no longer supports this feature.")
	_                                    = flag.Duration("cloudspanner_healthcheck_interval", 0, "DEPRECATED. Spanner no longer supports this feature.")
	_                                    = flag.Bool("cloudspanner_track_session_handles", false, "DEPRECATED. Spanner no longer supports this feature.")
	csDequeueAcrossMerkleBucketsFraction = flag.Float64("cloudspanner_dequeue_bucket_fraction", 0.75, "Fraction of merkle keyspace to dequeue from, set to zero to disable.")
	csReadOnlyStaleness                  = flag.Duration("cloudspanner_readonly_staleness", time.Minute, "How far in the past to perform readonly operations. Within limits, raising this should help to increase performance/reduce latency.")
	_                                    = flag.Uint64("cloudspanner_max_burst_sessions", 0, "No longer used")

	csMu              sync.RWMutex
	csStorageInstance *cloudSpannerProvider
	warnOnce          sync.Once
)

func init() {
	if err := storage.RegisterProvider("cloud_spanner", newCloudSpannerStorageProvider); err != nil {
		panic(err)
	}
}

func warn() { _ = "STUB: not implemented"; return }

// No need to exit, it's an unlikely error and doesn't affect operation.

type cloudSpannerProvider struct {
	client *spanner.Client
}

func configFromFlags() spanner.ClientConfig {
	_ = "STUB: not implemented"
	return *new(spanner.ClientConfig)
}

func optionsFromFlags() []option.ClientOption { _ = "STUB: not implemented"; return nil }

func newCloudSpannerStorageProvider(_ monitoring.MetricFactory) (storage.Provider, error) {
	_ = "STUB: not implemented"
	return *new(storage.Provider), nil
}

// LogStorage builds and returns a new storage.LogStorage using CloudSpanner.
func (s *cloudSpannerProvider) LogStorage() storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

// AdminStorage builds and returns a new storage.AdminStorage using CloudSpanner.
func (s *cloudSpannerProvider) AdminStorage() storage.AdminStorage {
	_ = "STUB: not implemented"
	return *new(storage.AdminStorage)
}

// Close shuts down this provider. Calls to the other methods will fail
// after this.
func (s *cloudSpannerProvider) Close() error { _ = "STUB: not implemented"; return nil }
