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

// Package serverutil holds code for running Trillian servers.
package serverutil

import (
	"context"
	"net/http"
	"time"

	"github.com/google/trillian"
	"github.com/google/trillian/extension"
	"google.golang.org/grpc"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// DefaultTreeDeleteThreshold is the suggested threshold for tree deletion.
	// It represents the minimum time a tree has to remain Deleted before being hard-deleted.
	DefaultTreeDeleteThreshold = 7 * 24 * time.Hour

	// DefaultTreeDeleteMinInterval is the suggested min interval between tree GC sweeps.
	// A tree GC sweep consists of listing deleted trees older than the deletion threshold and
	// hard-deleting them.
	// Actual runs happen randomly between [minInterval,2*minInterval).
	DefaultTreeDeleteMinInterval = 4 * time.Hour
)

// Main encapsulates the data and logic to start a Trillian server (Log or Map).
type Main struct {
	// Endpoints for RPC and HTTP servers.
	// HTTP is optional, if empty it'll not be bound.
	RPCEndpoint, HTTPEndpoint string

	// TLS Certificate and Key files for the server.
	TLSCertFile, TLSKeyFile string

	DBClose func() error

	Registry extension.Registry

	StatsPrefix string
	QuotaDryRun bool

	// RegisterServerFn is called to register RPC servers.
	RegisterServerFn func(*grpc.Server, extension.Registry) error

	// IsHealthy will be called whenever "/healthz" is called on the mux.
	// A nil return value from this function will result in a 200-OK response
	// on the /healthz endpoint.
	IsHealthy func(context.Context) error
	// HealthyDeadline is the maximum duration to wait wait for a successful
	// IsHealthy() call.
	HealthyDeadline time.Duration

	// AllowedTreeTypes determines which types of trees may be created through the Admin Server
	// bound by Main. nil means unrestricted.
	AllowedTreeTypes []trillian.TreeType

	TreeGCEnabled         bool
	TreeDeleteThreshold   time.Duration
	TreeDeleteMinInterval time.Duration

	// These will be added to the GRPC server options.
	ExtraOptions []grpc.ServerOption
}

func (m *Main) healthz(rw http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Run starts the configured server. Blocks until the server exits.
func (m *Main) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Let http.ListenAndServeTLS handle the error case when only one of the flags is set.

// 15 second exit time limit

// wait for all jobs to exit gracefully

// Give things a few seconds to tidy up

// newGRPCServer starts a new Trillian gRPC server.
func (m *Main) newGRPCServer() (*grpc.Server, error) { _ = "STUB: not implemented"; return nil, nil }

// Let credentials.NewServerTLSFromFile handle the error case when only one of the flags is set.

// AnnounceSelf announces this binary's presence to etcd. This calls the cancel
// function if the keepalive lease with etcd expires.  Returns a function that
// should be called on process exit.
// AnnounceSelf does nothing if client is nil.
func AnnounceSelf(ctx context.Context, client *clientv3.Client, etcdService, endpoint string, cancel func()) func() {
	_ = "STUB: not implemented"
	return nil
}

// Get a lease so our entry self-destructs.

// Use a background context because the original context may have been cancelled.

// listenKeepAliveRsp listens to `keepAliveRspCh` channel, and calls the cancel function
// to notify the lease expired.
func listenKeepAliveRsp(ctx context.Context, keepAliveRspCh <-chan *clientv3.LeaseKeepAliveResponse, cancel func()) {
	_ = "STUB: not implemented"
	return
}

// srvRun run the server and call `shutdown` when the context has been cancelled
func srvRun(ctx context.Context, run func() error, shutdown func()) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for run to return
