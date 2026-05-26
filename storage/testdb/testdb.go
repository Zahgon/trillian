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

// Package testdb creates new databases for tests.
package testdb

import (
	"context"
	"database/sql"
	"net/url"
	"testing"

	"github.com/google/trillian/testonly"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

const (
	// MySQLURIEnv is the name of the ENV variable checked for the test MySQL
	// instance URI to use. The value must have a trailing slash.
	MySQLURIEnv = "TEST_MYSQL_URI"

	// Note: sql.Open requires the URI to end with a slash.
	defaultTestMySQLURI = "root@tcp(127.0.0.1)/"

	// CockroachDBURIEnv is the name of the ENV variable checked for the test CockroachDB
	// instance URI to use. The value must have a trailing slash.
	CockroachDBURIEnv = "TEST_COCKROACHDB_URI"

	defaultTestCockroachDBURI = "postgres://root@localhost:26257/?sslmode=disable"
)

type storageDriverInfo struct {
	sqlDriverName string
	schema        string
	uriFunc       func(paths ...string) string
}

var (
	trillianMySQLSchema = testonly.RelativeToPackage("../mysql/schema/storage.sql")
	trillianCRDBSchema  = testonly.RelativeToPackage("../crdb/schema/storage.sql")
)

// DriverName is the name of a database driver.
type DriverName string

const (
	// DriverMySQL is the identifier for the MySQL storage driver.
	DriverMySQL DriverName = "mysql"
	// DriverCockroachDB is the identifier for the CockroachDB storage driver.
	DriverCockroachDB DriverName = "cockroachdb"
)

var driverMapping = map[DriverName]storageDriverInfo{
	DriverMySQL: {
		sqlDriverName: "mysql",
		schema:        trillianMySQLSchema,
		uriFunc:       mysqlURI,
	},
	DriverCockroachDB: {
		sqlDriverName: "postgres",
		schema:        trillianCRDBSchema,
		uriFunc:       crdbURI,
	},
}

// mysqlURI returns the MySQL connection URI to use for tests. It returns the
// value in the ENV variable defined by MySQLURIEnv. If the value is empty,
// returns defaultTestMySQLURI.
//
// We use an ENV variable, rather than a flag, for flexibility. Only a subset
// of the tests in this repo require a database and import this package. With a
// flag, it would be necessary to distinguish "go test" invocations that need a
// database, and those that don't. ENV allows to "blanket apply" this setting.
func mysqlURI(dbRef ...string) string { _ = "STUB: not implemented"; return "" }

// crdbURI returns the CockroachDB connection URI to use for tests. It returns the
// value in the ENV variable defined by CockroachDBURIEnv. If the value is empty,
// returns defaultTestCockroachDBURI.
func crdbURI(dbRef ...string) string { _ = "STUB: not implemented"; return "" }

func addPathToURI(uri *url.URL, paths ...string) string { _ = "STUB: not implemented"; return "" }

// If the path is the root path, we don't want to append a slash.

func getURL(unparsedurl string) *url.URL {
	_ = "STUB: not implemented"
	//nolint:errcheck // We're not expecting an error here.
	return nil
}

// MySQLAvailable indicates whether the configured MySQL database is available.
func MySQLAvailable() bool { _ = "STUB: not implemented"; return false }

// CockroachDBAvailable indicates whether the configured CockroachDB database is available.
func CockroachDBAvailable() bool { _ = "STUB: not implemented"; return false }

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
func newEmptyDB(ctx context.Context, driver DriverName) (*sql.DB, func(context.Context), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a randomly-named database and then connect using the new name.

// NewTrillianDB creates an empty database with the Trillian schema. The database name is randomly
// generated.
// NewTrillianDB is equivalent to Default().NewTrillianDB(ctx).
func NewTrillianDB(ctx context.Context, driver DriverName) (*sql.DB, func(context.Context), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func sanitize(script string) string { _ = "STUB: not implemented"; return "" }

// skip empty lines and comments

// SkipIfNoMySQL is a test helper that skips tests that require a local MySQL.
func SkipIfNoMySQL(t *testing.T) { _ = "STUB: not implemented"; return }

// SkipIfNoCockroachDB is a test helper that skips tests that require a local CockroachDB.
func SkipIfNoCockroachDB(t *testing.T) { _ = "STUB: not implemented"; return }
