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

package testonly

import (
	"github.com/golang/mock/gomock"
	"github.com/google/trillian"
	"github.com/google/trillian/testonly/tmock"
)

// MockServer implements the TrillianAdminServer, and TrillianLogServer.
type MockServer struct {
	Admin       *tmock.MockTrillianAdminServer
	Log         *tmock.MockTrillianLogServer
	AdminClient trillian.TrillianAdminClient
	LogClient   trillian.TrillianLogClient
	Addr        string
}

// NewMockServer starts a server on a random port.
// Returns the started server, the listener it's using for connection and a
// close function that must be defer-called on the scope the server is meant to
// stop.
func NewMockServer(ctrl *gomock.Controller) (*MockServer, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
