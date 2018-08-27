package rsd_csi

import (
    "testing"
//    "github.com/container-storage-interface/spec/lib/go/csi/v0"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
)

var fakeNs *nodeServer

func init() {
    if fakeNs == nil {
        d := NewDriver("CSINodeID", "tcp://127.0.0.1:10000")
        fakeNs = NewNodeServer(d)
    }
}

func TestInit(t *testing.T) {

}
