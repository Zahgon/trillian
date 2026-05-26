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

// The mdmtest binary runs merge delay tests against a Trillian Log.
package main

import (
	"context"
	"flag"
	"time"

	"k8s.io/klog/v2"
)

var (
	logID           = flag.Int64("log_id", 0, "Log ID to test against; ephemeral tree used if zero")
	rpcServer       = flag.String("rpc_server", "", "Trillian log server address:port")
	adminServer     = flag.String("admin_server", "", "Trillian admin server address:port (defaults to --rpc_server value)")
	metricsEndpoint = flag.String("metrics_endpoint", "", "Endpoint for serving metrics; if left empty, metrics will not be exposed")
	leafSize        = flag.Uint("leaf_size", 500, "Size of leaf values")
	newLeafChance   = flag.Uint("new_leaf_chance", 50, "Percentage chance of using a new leaf for each submission")
	checkers        = flag.Int("checkers", 3, "Number of parallel checker goroutines to run")
	emitInterval    = flag.Duration("emit_interval", 10*time.Second, "How often to output the summary info")
	deadline        = flag.Duration("deadline", 60*time.Second, "Deadline for single add+get-proof operation")
	minMergeDelay   = flag.Duration("min_merge_delay", 3*time.Second, "Minimum merge delay; don't check for inclusion until this interval has passed")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	if err := innerMain(ctx); err != nil {
		klog.Exit(err)
	}
}

func innerMain(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// No logID provided, so create an ephemeral tree to test against.
