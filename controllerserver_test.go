package rsd_csi

import (
    "fmt"
    "testing"
    "golang.org/x/net/context"
    "github.com/container-storage-interface/spec/lib/go/csi/v0"
)

var fakeCs *controllerServer


func init() {
    if fakeCs == nil {
        d := NewDriver("CSINodeID", "tcp://127.0.0.1:10000")
        fakeCs = NewControllerServer(d)
    }
}

func TestCreateVolume(t *testing.T) {
    fakeReq := &csi.CreateVolumeRequest{
		VolumeCapabilities: nil,
	}
    fakeCtx := context.Background()
    actualRes, err := fakeCs.CreateVolume(fakeCtx, fakeReq)
    if err != nil {
        t.Errorf("failed to CreateVolume: %v", err)
    }
    fmt.Printf("%q", actualRes)
}

func TestDeleteVolume(t *testing.T) {
    fakeReq := &csi.DeleteVolumeRequest{
		VolumeId: "5-sv-1-vl-8",
	}
    fakeCtx := context.Background()
    actualRes, err := fakeCs.DeleteVolume(fakeCtx, fakeReq)
    if err != nil {
        t.Errorf("failed to Delete: %v", err)
    }
    fmt.Printf("%q", actualRes)
}

