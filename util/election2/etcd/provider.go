package etcd

import (
	"flag"

	"github.com/google/trillian/util/election2"
	"k8s.io/klog/v2"
)

// ElectionName identifies the etcd election implementation.
const ElectionName = "etcd"

var (
	lockDir = flag.String("lock_file_path", "/test/multimaster", "etcd lock file directory path")
)

func init() {
	if err := election2.RegisterProvider(ElectionName, newFactory); err != nil {
		klog.Fatalf("Failed to register election implementation %v: %v", ElectionName, err)
	}
}

// NewFactory builds an election factory that uses the given parameters.
func newFactory() (election2.Factory, error) {
	_ = "STUB: not implemented"
	return *new(election2.Factory), nil
}

// The passed in etcd client should remain valid for the lifetime of the object.
