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
	"flag"
	"sync"

	"github.com/google/trillian/monitoring"
	"github.com/google/trillian/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"k8s.io/klog/v2"
)

var (
	postgreSQLURI        = flag.String("postgresql_uri", "postgresql:///defaultdb?host=localhost&user=test", "Connection URI for PostgreSQL database")
	postgresqlTLSCA      = flag.String("postgresql_tls_ca", "", "Path to the CA certificate file for PostgreSQL TLS connection ")
	postgresqlVerifyFull = flag.Bool("postgresql_verify_full", false, "Enable full TLS verification for PostgreSQL (sslmode=verify-full). If false, only sslmode=verify-ca is used.")

	postgresqlMu              sync.Mutex
	postgresqlErr             error
	postgresqlDB              *pgxpool.Pool
	postgresqlStorageInstance *postgresqlProvider
)

// GetDatabase returns an instance of PostgreSQL database, or creates one.
//
// TODO(robstradling): Make the dependency of PostgreSQL quota provider from
// PostgreSQL storage provider explicit.
func GetDatabase() (*pgxpool.Pool, error) { _ = "STUB: not implemented"; return nil, nil }

func init() {
	if err := storage.RegisterProvider("postgresql", newPostgreSQLStorageProvider); err != nil {
		klog.Fatalf("Failed to register storage provider postgresql: %v", err)
	}
}

type postgresqlProvider struct {
	db *pgxpool.Pool
	mf monitoring.MetricFactory
}

func newPostgreSQLStorageProvider(mf monitoring.MetricFactory) (storage.Provider, error) {
	_ = "STUB: not implemented"
	return *new(storage.Provider), nil
}

// getPostgreSQLDatabaseLocked returns an instance of PostgreSQL database, or creates
// one. Requires postgresqlMu to be locked.
func getPostgreSQLDatabaseLocked() (*pgxpool.Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *postgresqlProvider) LogStorage() storage.LogStorage {
	_ = "STUB: not implemented"
	return *new(storage.LogStorage)
}

func (s *postgresqlProvider) AdminStorage() storage.AdminStorage {
	_ = "STUB: not implemented"
	return *new(storage.AdminStorage)
}

func (s *postgresqlProvider) Close() error { _ = "STUB: not implemented"; return nil }

// BuildPostgresTLSURI modifies the given PostgreSQL URI to include TLS parameters based on flags.
// It returns the modified URI and any error encountered.
func BuildPostgresTLSURI(uri string) (string, error) { _ = "STUB: not implemented"; return "", nil }
