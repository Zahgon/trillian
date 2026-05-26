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

// Package interceptor defines gRPC interceptors for Trillian.
package interceptor

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/quota"
	"github.com/google/trillian/storage"
	"google.golang.org/grpc"
)

const (
	badInfoReason            = "bad_info"
	badTreeReason            = "bad_tree"
	insufficientTokensReason = "insufficient_tokens"
	getTreeStage             = "get_tree"
	getTokensStage           = "get_tokens"
	traceSpanRoot            = "/trillian/server/int"
)

var (
	// PutTokensTimeout is the timeout used for PutTokens calls.
	// PutTokens happens in a separate goroutine and with an independent context, therefore it has
	// its own timeout, separate from the RPC that causes the calls.
	PutTokensTimeout = 5 * time.Second

	requestCounter       monitoring.Counter
	requestDeniedCounter monitoring.Counter
	contextErrCounter    monitoring.Counter
	metricsOnce          sync.Once
	enabledServices      = map[string]bool{
		"trillian.TrillianLog":   true,
		"trillian.TrillianAdmin": true,
		"TrillianLog":            true,
		"TrillianAdmin":          true,
	}
)

// RequestProcessor encapsulates the logic to intercept a request, split into separate stages:
// before and after the handler is invoked.
type RequestProcessor interface {

	// Before implements all interceptor logic that happens before the handler is called.
	// It returns a (potentially) modified context that's passed forward to the handler (and After),
	// plus an error, in case the request should be interrupted before the handler is invoked.
	Before(ctx context.Context, req interface{}, method string) (context.Context, error)

	// After implements all interceptor logic that happens after the handler is invoked.
	// Before must be invoked prior to After and the same RequestProcessor instance must to be used
	// to process a given request.
	After(ctx context.Context, resp interface{}, method string, handlerErr error)
}

// TrillianInterceptor checks that:
// * Requests addressing a tree have the correct tree type and tree state;
// * TODO(codingllama): Requests are properly authenticated / authorized ; and
// * Requests are rate limited appropriately.
type TrillianInterceptor struct {
	admin storage.AdminStorage
	qm    quota.Manager

	// quotaDryRun controls whether lack of tokens actually blocks requests (if set to true, no
	// requests are blocked by lack of tokens).
	quotaDryRun bool
}

// New returns a new TrillianInterceptor instance.
func New(admin storage.AdminStorage, qm quota.Manager, quotaDryRun bool, mf monitoring.MetricFactory) *TrillianInterceptor {
	_ = "STUB: not implemented"
	return nil
}

func initMetrics(mf monitoring.MetricFactory) { _ = "STUB: not implemented"; return }

func incRequestDeniedCounter(reason string, treeID int64, quotaUser string) {
	_ = "STUB: not implemented"
	return
}

// UnaryInterceptor executes the TrillianInterceptor logic for unary RPCs.
func (i *TrillianInterceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	_ = "STUB: not implemented"
	// Implement UnaryInterceptor using a RequestProcessor, so we
	// 1. exercise it
	// 2. make it easier to port this logic to non-gRPC implementations.
	return nil, nil
}

// NewProcessor returns a RequestProcessor for the TrillianInterceptor logic.
func (i *TrillianInterceptor) NewProcessor() RequestProcessor {
	_ = "STUB: not implemented"
	return *new(RequestProcessor)
}

type trillianProcessor struct {
	parent *TrillianInterceptor
	info   *rpcInfo
}

func (tp *trillianProcessor) Before(ctx context.Context, req interface{}, method string) (context.Context, error) {
	_ = "STUB: not implemented"
	// Skip if the interceptor is not enabled for this service.
	return *new(context.Context), nil
}

// Don't want the Before to contain the action, so don't overwrite the ctx.

// TODO(codingllama): Add auth interception

func (tp *trillianProcessor) After(ctx context.Context, resp interface{}, method string, handlerErr error) {
	_ = "STUB: not implemented"
	return
}

// After() currently only does quota processing

// Decide if we have to replenish tokens. There are a few situations that require tokens to
// be replenished:
// * Invalid requests (a bad request shouldn't spend sequencing-based tokens, as it won't
//   cause a corresponding sequencing to happen)
// * Requests that filter out duplicates (e.g., QueueLeaf, for the same reason as above:
//   duplicates aren't queued for sequencing)
// These are only applied for Refundable specs.

// Return the tokens spent by invalid requests

// Run PutTokens in a separate goroutine and with a separate context.
// It shouldn't block RPC completion, nor should it share the RPC's context deadline.

// TODO(codingllama): If PutTokens turns out to be unreliable we can still leak tokens. In
// this case, we may want to keep tabs on how many tokens we failed to replenish and bundle
// them up in the next PutTokens call (possibly as a QuotaManager decorator, or internally
// in its impl).

func isLeafOK(leaf *trillian.QueuedLogLeaf) bool {
	_ = "STUB: not implemented"
	// Be biased in favor of OK, as that matches TrillianLogRPCServer's behavior.
	return false
}

var (
	fullyQualifiedRE = regexp.MustCompile(`^/([\w.]+)/(\w+)$`)
	unqualifiedRE    = regexp.MustCompile(`^/(\w+)\.(\w+)$`)
)

// serviceName returns the fully qualified service name
// "some.package.service" for "/some.package.service/method".
// It returns the unqualified service name "service" for "/service.method".
func serviceName(fullMethod string) string { _ = "STUB: not implemented"; return "" }

type rpcInfo struct {
	// getTree indicates whether the interceptor should populate treeID.
	getTree bool

	readonly  bool
	treeID    int64
	treeTypes []trillian.TreeType

	specs  []quota.Spec
	tokens int
	// Single string describing all of the users against which quota is requested.
	quotaUsers string
}

// chargable is satisfied by request proto messages which contain a GetChargeTo
// accessor.
type chargable interface {
	GetChargeTo() *trillian.ChargeTo
}

// chargedUsers returns user identifiers for any chargable user quotas.
func chargedUsers(req interface{}) []string { _ = "STUB: not implemented"; return nil }

func newRPCInfoForRequest(req interface{}) (*rpcInfo, error) {
	_ = "STUB: not implemented"
	// Set "safe" defaults: enable all interception and assume requests are readonly.
	return nil, nil
}

// Not intercepted at all

// Quota configuration requests

// Doesn't really matter as all interceptors are turned off

// Admin create

// Tree doesn't exist

// Admin list

// Zero to many trees

// Admin / readonly

// Read done within RPC handler

// Admin / readwrite

// Read-modify-write done within RPC handler

// (Log + Pre-ordered Log) / readonly

// Log / readwrite

// Pre-ordered Log / readwrite

// (Log + Pre-ordered Log) / readwrite

func newRPCInfo(req interface{}) (*rpcInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// Only Global tokens are refunded.

type logIDRequest interface {
	GetLogId() int64
}

type treeIDRequest interface {
	GetTreeId() int64
}

type treeRequest interface {
	GetTree() *trillian.Tree
}

// ErrorWrapper is a grpc.UnaryServerInterceptor that wraps the errors emitted by the underlying handler.
func ErrorWrapper(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func spanFor(ctx context.Context, name string) (context.Context, func()) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
