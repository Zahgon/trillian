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

// Package pem contains functions for marshaling and unmarshaling keys in PEM format.
package pem

import (
	"context"
	"crypto"

	"google.golang.org/protobuf/proto"
)

// ReadPrivateKeyFile reads a PEM-encoded private key from a file.
// The key may be protected by a password. If password is empty, the key is
// assumed to be unencrypted.
func ReadPrivateKeyFile(file, password string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// UnmarshalPrivateKey reads a PEM-encoded private key from a string.
// The key may be protected by a password.
func UnmarshalPrivateKey(keyPEM, password string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

//nolint:staticcheck

// FromProto builds a crypto.Signer from a proto.Message, which must be of type PEMKeyFile.
func FromProto(_ context.Context, pb proto.Message) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
