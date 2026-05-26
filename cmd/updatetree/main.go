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

// Package main contains the implementation and entry point for the updatetree
// command.
//
// Example usage:
// $ ./updatetree --admin_server=host:port --tree_id=123456789 --tree_state=FROZEN
//
// The output is minimal to allow for easy usage in automated scripts.
package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/google/trillian"
	"google.golang.org/protobuf/encoding/prototext"
	"k8s.io/klog/v2"
)

var (
	adminServerAddr = flag.String("admin_server", "", "Address of the gRPC Trillian Admin Server (host:port)")
	rpcDeadline     = flag.Duration("rpc_deadline", time.Second*10, "Deadline for RPC requests")
	treeID          = flag.Int64("tree_id", 0, "The ID of the tree to be set updated")
	treeState       = flag.String("tree_state", "", "If set the tree state will be updated")
	treeType        = flag.String("tree_type", "", "If set the tree type will be updated")
	printTree       = flag.Bool("print", false, "Print the resulting tree")
)

// TODO(Martin2112): Pass everything needed into this and don't refer to flags.
func updateTree(ctx context.Context) (*trillian.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We only want to update certain fields of the tree, which means we
// need a field mask on the request.

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	defer klog.Flush()

	ctx, cancel := context.WithTimeout(context.Background(), *rpcDeadline)
	defer cancel()
	tree, err := updateTree(ctx)
	if err != nil {
		klog.Exitf("Failed to update tree: %v", err)
	}

	if *printTree {
		fmt.Println(prototext.Format(tree))
	} else {
		// DO NOT change the default output format, some scripts depend on it. If
		// you really want to change it, hide the new format behind a flag.
		fmt.Println(tree.TreeState)
	}
}
