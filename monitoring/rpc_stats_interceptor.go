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

// Package monitoring provides monitoring functionality.
package monitoring

import (
	"time"

	"github.com/google/trillian/util/clock"
	"google.golang.org/grpc"
)

const traceSpanRoot = "/trillian/mon/"

// RPCStatsInterceptor provides a gRPC interceptor that records statistics about the RPCs passing through it.
type RPCStatsInterceptor struct {
	prefix            string
	timeSource        clock.TimeSource
	ReqCount          Counter
	ReqSuccessCount   Counter
	ReqSuccessLatency Histogram
	ReqErrorCount     Counter
	ReqErrorLatency   Histogram
}

// NewRPCStatsInterceptor creates a new RPCStatsInterceptor for the given application/component, with
// a specified time source.
func NewRPCStatsInterceptor(timeSource clock.TimeSource, prefix string, mf MetricFactory) *RPCStatsInterceptor {
	_ = "STUB: not implemented"
	return nil
}

func prefixedName(prefix, name string) string { _ = "STUB: not implemented"; return "" }

func (r *RPCStatsInterceptor) recordFailureLatency(labels []string, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// Interceptor returns a UnaryServerInterceptor that can be registered with an RPC server and
// will record request counts / errors and latencies for that servers handlers
func (r *RPCStatsInterceptor) Interceptor() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// This interceptor wraps the request handler so we should track the
// additional latency it imposes.

// Increase the request count for the method and start the clock

// If we reach here then the handler exited via panic, count it as a server failure

// Invoke the actual operation

// Record success / failure and latency

// Pass the result of the handler invocation back
