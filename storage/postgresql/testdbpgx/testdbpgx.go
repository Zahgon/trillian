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

// Package testdbpgx creates new PostgreSQL databases for tests.
package testdbpgx

import (
	"context"
	"testing"

	"github.com/google/trillian/testonly"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// PostgreSQLURIEnv is the name of the ENV variable checked for the test PostgreSQL
	// instance URI to use. The value must have a trailing slash.
	PostgreSQLURIEnv = "TEST_POSTGRESQL_URI"

	defaultTestPostgreSQLURI = "postgresql:///defaultdb?host=localhost&user=postgres&password=postgres"
)

type storageDriverInfo struct {
	schema  string
	uriFunc func(paths ...string) string
}

var (
	trillianPostgreSQLSchema = testonly.RelativeToPackage("../schema/storage.sql")
)

// DriverName is the name of a database driver.
type DriverName string

const (
	// DriverPostgreSQL is the identifier for the PostgreSQL storage driver.
	DriverPostgreSQL DriverName = "postgresql"
)

var driverMapping = map[DriverName]storageDriverInfo{
	DriverPostgreSQL: {
		schema:  trillianPostgreSQLSchema,
		uriFunc: postgresqlURI,
	},
}

// postgresqlURI returns the PostgreSQL connection URI to use for tests. It returns the
// value in the ENV variable defined by PostgreSQLURIEnv. If the value is empty,
// returns defaultTestPostgreSQLURI.
//
// We use an ENV variable, rather than a flag, for flexibility. Only a subset
// of the tests in this repo require a database and import this package. With a
// flag, it would be necessary to distinguish "go test" invocations that need a
// database, and those that don't. ENV allows to "blanket apply" this setting.
func postgresqlURI(dbRef ...string) string { _ = "STUB: not implemented"; return "" }

// No equals character, so use this string as the database name.

// PostgreSQLAvailable indicates whether the configured PostgreSQL database is available.
func PostgreSQLAvailable() bool { _ = "STUB: not implemented"; return false }

func dbAvailable(driver DriverName) bool { _ = "STUB: not implemented"; return false }

// SetFDLimit sets the soft limit on the maximum number of open file descriptors.
// See http://man7.org/linux/man-pages/man2/setrlimit.2.html
func SetFDLimit(uLimit uint64) error { _ = "STUB: not implemented"; return nil }

// newEmptyDB creates a new, empty database.
// It returns the database handle and a clean-up function, or an error.
// The returned clean-up function should be called once the caller is finished
// using the DB, the caller should not continue to use the returned DB after
// calling this function as it may, for example, delete the underlying
// instance.
func newEmptyDB(ctx context.Context, driver DriverName) (*pgxpool.Pool, func(context.Context), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a randomly-named database and then connect using the new name.

// NewTrillianDB creates an empty database with the Trillian schema. The database name is randomly
// generated.
// NewTrillianDB is equivalent to Default().NewTrillianDB(ctx).
func NewTrillianDB(ctx context.Context, driver DriverName) (*pgxpool.Pool, func(context.Context), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Execute each statement in the schema file.  Each statement must end with a semicolon, and there must be a blank line before the next statement.

func sanitize(script string) string { _ = "STUB: not implemented"; return "" }

// skip empty lines and comments

// SkipIfNoPostgreSQL is a test helper that skips tests that require a local PostgreSQL.
func SkipIfNoPostgreSQL(t *testing.T) { _ = "STUB: not implemented"; return }
