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

// Package der contains functions for marshaling and unmarshaling keys in DER format.
package der

import (
	"context"
	"crypto"

	"google.golang.org/protobuf/proto"
)

// FromProto builds a crypto.Signer from a proto.Message, which must be of type PrivateKey.
func FromProto(_ context.Context, pb proto.Message) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// UnmarshalPrivateKey reads a DER-encoded private key.
func UnmarshalPrivateKey(keyDER []byte) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// UnmarshalPublicKey reads a DER-encoded public key.
func UnmarshalPublicKey(keyDER []byte) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

// MarshalPublicKey serializes an RSA or ECDSA public key as DER.
func MarshalPublicKey(pubKey crypto.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalPrivateKey serializes an RSA or ECDSA private key as DER.
func MarshalPrivateKey(key crypto.Signer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
