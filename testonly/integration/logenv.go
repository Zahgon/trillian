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

// Package integration provides test-only code for performing integrated
// tests of Trillian functionality.
package integration

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/google/trillian"
	"github.com/google/trillian/extension"
	"github.com/google/trillian/log"
	"github.com/google/trillian/server"
	"github.com/google/trillian/server/admin"
	"github.com/google/trillian/util/clock"

	_ "github.com/go-sql-driver/mysql" // Load MySQL driver
)

var (
	sequencerWindow = time.Duration(0)
	batchSize       = 50
	// SequencerInterval is the time between runs of the sequencer.
	SequencerInterval = 500 * time.Millisecond
	timeSource        = clock.System
)

// LogEnv is a test environment that contains both a log server and a connection to it.
type LogEnv struct {
	registry        extension.Registry
	pendingTasks    *sync.WaitGroup
	grpcServer      *grpc.Server
	adminServer     *admin.Server
	logServer       *server.TrillianLogRPCServer
	LogOperation    log.Operation
	Sequencer       *log.OperationManager
	sequencerCancel context.CancelFunc
	ClientConn      *grpc.ClientConn // TODO(gbelvin): Deprecate.

	Address string
	Log     trillian.TrillianLogClient
	Admin   trillian.TrillianAdminClient
	DB      *sql.DB
	dbDone  func(context.Context)
}

// NewLogEnv creates a fresh DB, log server, and client. The numSequencers parameter
// indicates how many sequencers to run in parallel; if numSequencers is zero a
// manually-controlled test sequencer is used.
//
// Deprecated: Use NewLogEnvWithGRPCOptions instead
//
// TODO(Martin2112): Remove this constructor, it is only used by tests and
// can be replaced by one of the others.
func NewLogEnv(ctx context.Context, numSequencers int, _ string) (*LogEnv, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLogEnvWithGRPCOptions creates a fresh DB, log server, and client. The
// numSequencers parameter indicates how many sequencers to run in parallel;
// if numSequencers is zero a manually-controlled test sequencer is used.
// Additional grpc.ServerOption and grpc.DialOption values can be provided.
func NewLogEnvWithGRPCOptions(ctx context.Context, numSequencers int, serverOpts []grpc.ServerOption, clientOpts []grpc.DialOption) (*LogEnv, error) {
	_ = "STUB: not implemented"
	// TODO(jaosorior): Make this configurable for Cockroach or MySQL
	return nil, nil
}

// NewLogEnvWithRegistry uses the passed in Registry to create a log server,
// and client. The numSequencers parameter indicates how many sequencers to
// run in parallel; if numSequencers is zero a manually-controlled test
// sequencer is used.
func NewLogEnvWithRegistry(ctx context.Context, numSequencers int, registry extension.Registry) (*LogEnv, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLogEnvWithRegistryAndGRPCOptions works the same way as NewLogEnv, but allows callers to also set additional grpc.ServerOption and grpc.DialOption values.
func NewLogEnvWithRegistryAndGRPCOptions(ctx context.Context, numSequencers int, registry extension.Registry, serverOpts []grpc.ServerOption, clientOpts []grpc.DialOption) (*LogEnv, error) {
	_ = "STUB: not implemented"
	// Create the GRPC Server.
	return nil, nil
}

// Setup the Admin Server.

// Setup the Log Server.

// Create Sequencer.

// Start a live sequencer in a goroutine.

// Listen and start server.

// Connect to the server.

// Close shuts down the server.
func (env *LogEnv) Close() { _ = "STUB: not implemented"; return }
