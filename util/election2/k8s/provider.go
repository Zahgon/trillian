package k8s

import (
	"flag"

	"github.com/google/trillian/util/election2"
	"k8s.io/klog/v2"
)

// ElectionName identifies the kubernetes election implementation.
const ElectionName = "k8s"

var (
	kubeconfig = flag.String("kubeconfig", "", "Paths to a kubeconfig. Only required if out-of-cluster.")
	namespace  = flag.String("lock_namespace", "", "The lease lock resource namespace. Only effective for election_system=k8s.")
	instanceID = flag.String("lock_holder_identity", "", "The identity of the holder of a current lease")
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
