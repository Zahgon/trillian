// Copyright 2024 Trillian Authors. All Rights Reserved.
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

package postgresql

import (
	"context"
	"time"

	"github.com/google/trillian"
	"github.com/jackc/pgx/v5"
)

const (
	// If this statement ORDER BY clause is changed refer to the comment in removeSequencedLeaves
	selectQueuedLeavesSQL = "SELECT LeafIdentityHash,MerkleLeafHash,QueueTimestampNanos,QueueID " +
		"FROM Unsequenced " +
		"WHERE TreeId=$1" +
		" AND Bucket=0" +
		" AND QueueTimestampNanos<=$2 " +
		"ORDER BY QueueTimestampNanos,LeafIdentityHash " +
		"LIMIT $3"
	insertUnsequencedEntrySQL = "INSERT INTO Unsequenced(TreeId,Bucket,LeafIdentityHash,MerkleLeafHash,QueueTimestampNanos,QueueID) VALUES($1,0,$2,$3,$4,$5)"
	deleteUnsequencedSQL      = "DELETE FROM Unsequenced WHERE QueueID=ANY($1)"
)

type dequeuedLeaf []byte

func dequeueInfo(_ []byte, queueID []byte) dequeuedLeaf {
	_ = "STUB: not implemented"
	return *new(dequeuedLeaf)
}

func (t *logTreeTX) dequeueLeaf(rows pgx.Rows) (*trillian.LogLeaf, dequeuedLeaf, error) {
	_ = "STUB: not implemented"
	return nil, *new(dequeuedLeaf), nil
}

// Note: the LeafData and ExtraData being nil here is OK as this is only used by the
// sequencer. The sequencer only writes to the SequencedLeafData table and the client
// supplied data was already written to LeafData as part of queueing the leaf.

func generateQueueID(treeID int64, leafIdentityHash []byte, timestamp int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func queueArgs(treeID int64, identityHash []byte, queueTimestamp time.Time) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (t *logTreeTX) UpdateSequencedLeaves(ctx context.Context, leaves []*trillian.LogLeaf) error {
	_ = "STUB: not implemented"
	return nil
}

// This should fail on insert but catch it early

// Copy sequenced leaves to SequencedLeafData table.

// removeSequencedLeaves removes the passed in leaves slice (which may be
// modified as part of the operation).
func (t *logTreeTX) removeSequencedLeaves(ctx context.Context, queueIDs []dequeuedLeaf) error {
	_ = "STUB: not implemented"

	// Don't need to re-sort because the query ordered by leaf hash. If that changes because
	// the query is expensive then the sort will need to be done here. See comment in
	// QueueLeaves.
	return nil
}

// Error is handled by checkResultOkAndRowCountIs() below
