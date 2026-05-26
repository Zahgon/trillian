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

// Package mysqlqm defines a MySQL-based quota.Manager implementation.
package mysqlqm

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/trillian/quota"
)

const (
	// DefaultMaxUnsequenced is a suggested value for MaxUnsequencedRows.
	// Note that this is a Global/Write quota suggestion, so it applies across trees.
	DefaultMaxUnsequenced = 500000 // About 2h of non-stop signing at 70QPS.

	countFromInformationSchemaQuery = `
		SELECT table_rows
		FROM information_schema.tables
		WHERE table_schema = schema()
			AND table_name = ?
			AND table_type = ?`
	countFromUnsequencedQuery = "SELECT COUNT(*) FROM Unsequenced"
)

// ErrTooManyUnsequencedRows is returned when tokens are requested but Unsequenced has grown
// beyond the configured limit.
var ErrTooManyUnsequencedRows = errors.New("too many unsequenced rows")

// QuotaManager is a MySQL-based quota.Manager implementation.
//
// It has two working modes: one queries the information schema for the number of Unsequenced rows,
// the other does a select count(*) on the Unsequenced table. Information schema queries are
// default, even though they are approximate, as they're constant time (select count(*) on InnoDB
// based MySQL needs to traverse the index and may take quite a while to complete).
//
// QuotaManager only implements Global/Write quotas, which is based on the number of Unsequenced
// rows (to be exact, tokens = MaxUnsequencedRows - actualUnsequencedRows).
// Other quotas are considered infinite.
type QuotaManager struct {
	DB                 *sql.DB
	MaxUnsequencedRows int
	UseSelectCount     bool
}

// GetTokens implements quota.Manager.GetTokens.
// It doesn't actually reserve or retrieve tokens, instead it allows access based on the number of
// rows in the Unsequenced table.
func (m *QuotaManager) GetTokens(ctx context.Context, numTokens int, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// Only allow global writes if Unsequenced is under the expected limit

// PutTokens implements quota.Manager.PutTokens.
// It's a noop for QuotaManager.
func (m *QuotaManager) PutTokens(ctx context.Context, numTokens int, specs []quota.Spec) error {
	_ = "STUB: not implemented"

	// ResetQuota implements quota.Manager.ResetQuota.
	// It's a noop for QuotaManager.
	return nil
}

func (m *QuotaManager) ResetQuota(ctx context.Context, specs []quota.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *QuotaManager) countUnsequenced(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func countFromInformationSchema(ctx context.Context, db *sql.DB) (int, error) {
	_ = "STUB: not implemented"
	// turn off statistics caching for MySQL 8
	return 0, nil
}

// information_schema.tables doesn't have an explicit PK, so let's play it safe and ensure
// the cursor returns a single row.

func countFromTable(ctx context.Context, db *sql.DB) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// turnOffInformationSchemaCache turn off statistics caching for MySQL 8
// To always retrieve the latest statistics directly from the storage engine and bypass cached values, set information_schema_stats_expiry to 0.
// See https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_information_schema_stats_expiry
// MySQL versions prior to 8 will fail safely.
func turnOffInformationSchemaCache(ctx context.Context, db *sql.DB) error {
	_ = "STUB: not implemented"
	return nil
}

// fail safely for all versions of MySQL prior to 8
