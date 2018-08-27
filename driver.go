/*
*/

package rsd_csi

import (
//	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/golang/glog"

	"github.com/kubernetes-csi/drivers/pkg/csi-common"
)



type driver struct {
	csiDriver *csicommon.CSIDriver
    endpoint    string

	ids *identityServer
	ns  *nodeServer
	cs  *controllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
    driverName = "csi-rsdplugin"
)

var (
	version  = "0.1.0"
)

func NewDriver(nodeID, endpoint string) *driver {
    glog.Infof("Driver: %v version: %v", driverName, version)

    d := &driver{}
    d.endpoint = endpoint

    csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
    csiDriver.AddControllerServiceCapabilities(
		[]csi.ControllerServiceCapability_RPC_Type{
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		})
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER})

	d.csiDriver = csiDriver

	return d
}

func NewIdentityServer(d *driver) *identityServer {
	return &identityServer{
		DefaultIdentityServer: csicommon.NewDefaultIdentityServer(d.csiDriver),
	}
}

func NewControllerServer(d *driver) *controllerServer {
	return &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d.csiDriver),
	}
}

func NewNodeServer(d *driver) *nodeServer {
	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.csiDriver),
	}
}

func (d *driver) Run() {
	glog.Infof("Driver: %v ", driverName)
//    csicommon.RunControllerandNodePublishServer(d.endpoint, d.csiDriver, NewControllerServer(d), NewNodeServer(d))
}

