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
	"database/sql"
	"time"

	"github.com/google/trillian"
)

// toMillisSinceEpoch converts a timestamp into milliseconds since epoch
func toMillisSinceEpoch(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

// fromMillisSinceEpoch converts
func fromMillisSinceEpoch(ts int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// setNullStringIfValid assigns src to dest if src is Valid.
func setNullStringIfValid(src sql.NullString, dest *string) { _ = "STUB: not implemented"; return }

// row defines a common interface between sql.Row and sql.Rows(!)
type row interface {
	Scan(dest ...interface{}) error
}

// readTree takes a sql row and returns a tree
func readTree(r row) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil,

		// Enums and Datetimes need an extra conversion step
		nil
}

// Convert all things!

// Let's make sure we didn't mismatch any of the casts above
