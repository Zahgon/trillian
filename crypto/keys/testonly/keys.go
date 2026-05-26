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

// Package testonly contains code and data that should only be used by tests.
// Production code MUST NOT depend on anything in this package. This will be
// enforced by tools where possible.
package testonly

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
)

// MustMarshalPublicPEMToDER reads a PEM-encoded public key and returns it in DER encoding.
// If an error occurs, it panics.
func MustMarshalPublicPEMToDER(keyPEM string) []byte { _ = "STUB: not implemented"; return nil }

// SignAndVerify exercises a signer by using it to generate a signature, and
// then verifies that this signature is correct.
func SignAndVerify(signer crypto.Signer, pubKey crypto.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Ed25519 performs two passes over the data and so takes the whole message not just the digest.

func verifyECDSA(pubKey *ecdsa.PublicKey, digest, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyRSA(pubKey *rsa.PublicKey, digest, sig []byte, hasher crypto.Hash, opts crypto.SignerOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyEd25519(pubKey ed25519.PublicKey, digest, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}
