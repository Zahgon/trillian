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

// Package signer is a sample implementation of a commit-log based signer.
package signer

import (
	"flag"
	"sync"
	"time"

	"github.com/google/trillian/docs/storage/commit_log/simelection"
)

var (
	batchSize         = flag.Int("batch_size", 5, "Maximum leaves to sign in one run")
	pessimizeInterval = flag.Duration("signer_pessimize", 10*time.Millisecond, "Pause interval in signing to induce inter-signer problems")
)

// Signer is a simulated signer instance.
type Signer struct {
	mu        sync.RWMutex
	Name      string
	election  *simelection.Election
	epoch     int64
	dbSTHInfo STHInfo
	db        FakeDatabase
}

// New creates a simulated signer that uses the provided election.
func New(name string, election *simelection.Election, epoch int64) *Signer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Signer) String() string { _ = "STUB: not implemented"; return "" }

// LatestSTHInfo returns the most recent STHInfo known about by the signer
func (s *Signer) LatestSTHInfo() STHInfo {
	_ = "STUB: not implemented"

	// StoreSTHInfo updates the STHInfo known about by the signer.
	return *new(STHInfo)
}

func (s *Signer) StoreSTHInfo(info STHInfo) {
	_ = "STUB: not implemented"

	// IsMaster indicates if this signer is master.
	return
}

func (s *Signer) IsMaster() bool { _ = "STUB: not implemented"; return false }

// Run performs a single signing run.
func (s *Signer) Run() { _ = "STUB: not implemented"; return }

// Read from local DB to see what STH we know about locally.

// Sanity check that the STH table has what we already know.

// Look to see if anyone else has already stored data just ahead of our STH.  This will
// normally be the next entry, but we need to ignore any entries that have inconsistent
// offsets.

// Found an entry in the STHs topic that didn't get stored at the offset its writer
// expected it to be stored at, probably due to another master signer nipping in an
// entry ahead of it (due to a bug in mastership election).
// Kafka adjudicates the clash: whichever entry got the correct offset wins.

// We're up-to-date with the STHs topic (as of a moment ago) ...

// ... and we're the master. Move the STHs topic along to encompass any unincorporated leaves.

// The offset we expect this STH to end up at in STH topic

// The STH didn't get stored at the offset we expected, presumably because someone else got there first

// Now the STH topic is updated (correctly), do our local DB

// There is an STH one ahead of us that we're not caught up with yet.
// Read the leaves between what we have in our DB, and that STH...

// ... and store it in our local DB

// We may still not be caught up, but that's for the next time around.
