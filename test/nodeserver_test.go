package rsd_csi

import (
    "testing"
    "github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
    "golang.org/x/net/context"
)

var fakeNs *nodeServer
var fakeDevicePath = "/dev/xxx"
var fakeTargetPath = "/mnt/5-sv-1-vl-5"
var fakeVolID = "5-sv-1-vl-5"
var fakeCtx = context.Background()


func init() {
    if fakeNs == nil {
        d := NewDriver("CSINodeID", "tcp://127.0.0.1:10000")
        fakeNs = NewNodeServer(d)
    }
}

func TestNodePublishVolume(t *testing.T) {
   assert := assert.New(t)

    expectedRes := &csi.NodePublishVolumeResponse{}

    fakeReq := &csi.NodePublishVolumeRequest{
		VolumeId:         fakeVolID,
		PublishInfo:      map[string]string{"DevicePath": fakeDevicePath},
		TargetPath:       fakeTargetPath,
		VolumeCapability: nil,
		Readonly:         false,
	}
    actualRes, err := fakeNs.NodePublishVolume(fakeCtx, fakeReq)
	if err != nil {
		t.Errorf("failed to NodePublishVolume: %v", err)
	}

	// Assert
	assert.Equal(expectedRes, actualRes)
}
