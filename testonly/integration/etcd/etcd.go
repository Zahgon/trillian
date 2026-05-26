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

// Package etcd contains a helper to start an embedded etcd server.
package etcd

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
)

const (

	// MaxEtcdStartAttempts is the max number of start attempts made before it fails.
	MaxEtcdStartAttempts = 3

	defaultTimeout = 5 * time.Second
	tempDirPrefix  = "etcdquota-test-"
)

// StartEtcd returns a started, ready to use embedded etcd, along with a client and a cleanup
// function (that must be defer-called). There's no need to defer-close etcd of client, cleanup
// closes both.
//
// A temp directory and random ports are used to setup etcd.
func StartEtcd() (e *embed.Etcd, c *clientv3.Client, cleanup func(), err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// OK

func tryStartEtcd(dir string) (*embed.Etcd, error) { _ = "STUB: not implemented"; return nil, nil }

// OK to ignore err, it'll error below if parsing fails
