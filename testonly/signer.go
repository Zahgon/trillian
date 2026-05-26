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

package testonly

import (
	"crypto"
	"io"
)

// signerStub returns a fixed signature and error, no matter the input.
// It implements crypto.Signer.
type signerStub struct {
	publicKey crypto.PublicKey
	signature []byte
	err       error
}

// Public returns the public key associated with the signer that this stub is based on.
func (s *signerStub) Public() crypto.PublicKey {
	_ = "STUB: not implemented"

	// Sign will return the signature or error that the signerStub was created to provide.
	return *new(crypto.PublicKey)
}

func (s *signerStub) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// NewSignerWithErr creates a signer that always returns err when Sign() is called.
		nil
}

func NewSignerWithErr(pubKey crypto.PublicKey, err error) crypto.Signer {
	_ = "STUB: not implemented"
	return *new(crypto.Signer)
}

// NewSignerWithFixedSig creates a signer that always return sig when Sign() is called.
func NewSignerWithFixedSig(pubKey crypto.PublicKey, sig []byte) crypto.Signer {
	_ = "STUB: not implemented"
	return *new(crypto.Signer)
}
