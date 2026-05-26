// Copyright 2018 Google LLC. All Rights Reserved.
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

package mysql

import (
	"database/sql"
	"flag"
	"sync"

	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	"k8s.io/klog/v2"
	// Load MySQL driver
)

var (
	mySQLURI        = flag.String("mysql_uri", "test:zaphod@tcp(127.0.0.1:3306)/test", "Connection URI for MySQL database")
	maxConns        = flag.Int("mysql_max_conns", 0, "Maximum connections to the database")
	maxIdle         = flag.Int("mysql_max_idle_conns", -1, "Maximum idle database connections in the connection pool")
	mySQLTLSCA      = flag.String("mysql_tls_ca", "", "Path to the CA certificate file for MySQL TLS connection ")
	mySQLServerName = flag.String("mysql_server_name", "", "Name of the MySQL server to be used as the Server Name in the TLS configuration")

	mysqlMu              sync.Mutex
	mysqlErr             error
	mysqlDB              *sql.DB
	mysqlStorageInstance *mysqlProvider
)

// GetDatabase returns an instance of MySQL database, or creates one.
//
// TODO(pavelkalinnikov): Make the dependency of MySQL quota provider from
// MySQL storage provider explicit.
func GetDatabase() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func init() {
	if err := storage.RegisterProvider("mysql", newMySQLStorageProvider); err != nil {
		klog.Fatalf("Failed to register storage provider mysql: %v", err)
	}
}

type mysqlProvider struct {
	db *sql.DB
	mf monitoring.MetricFactory
}

func newMySQLStorageProvider(mf monitoring.MetricFactory) (storage.Provider, error) {
	_ = "STUB: not implemented"
	return *new(storage.Provider), nil
}

// getMySQLDatabaseLocked returns an instance of MySQL database, or creates
// one. Requires mysqlMu to be locked.
func getMySQLDatabaseLocked() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *mysqlProvider) LogStorage() storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

func (s *mysqlProvider) AdminStorage() storage.AdminStorage {
	_ = "STUB: not implemented"
	return *new(storage.AdminStorage)
}

func (s *mysqlProvider) Close() error { _ = "STUB: not implemented"; return nil }

// registerMySQLTLSConfig registers a custom TLS config for MySQL using a provided CA certificate and optional server name.
// Returns an error if the CA certificate can't be read or added to the root cert pool, or when the registration of the TLS config fails.
func registerMySQLTLSConfig() error { _ = "STUB: not implemented"; return nil }
