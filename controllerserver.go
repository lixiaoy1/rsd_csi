/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rsd_csi

import (
	//"fmt"
	//"os"

	"github.com/golang/glog"
	//"github.com/pborman/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
)

const (
	deviceID           = "deviceID"
	provisionRoot      = "/tmp/"
	maxStorageCapacity = 1024 * 1024 * 1024 * 1024
)

type controllerServer struct {
	*csicommon.DefaultControllerServer
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
    glog.V(1).Infof("Start CreateVolume")
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		glog.V(3).Infof("invalid create volume req: %v", req)
		return nil, err
	}

    volSizeBytes := int64(1 * 1024 * 1024 * 1024)
    if req.GetCapacityRange() != nil {
         volSizeBytes = int64(req.GetCapacityRange().GetRequiredBytes())
    }

    rsdClient, err := GetRSDProvider()
    if err != nil {
        glog.V(3).Infof("Failed to GetRSDProvider: %v", err)
        return nil, err
    }
    volumeID := ""
    volumeID, err = rsdClient.CreateVolume("", volSizeBytes)
    if err != nil {
        glog.V(3).Infof("Failed to CreateVolume: %v", err)
        return nil, err
    }

    glog.V(4).Infof("create volume %s", volumeID)
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			Id:            volumeID,
			CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
			Attributes:    req.GetParameters(),
		},
	}, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	glog.V(1).Infof("Start DeleteVolume")
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		glog.V(3).Infof("invalid delete volume req: %v", req)
		return nil, err
	}

	rsdClient, err := GetRSDProvider()
    if err != nil {
        glog.V(3).Infof("Failed to GetRSDProvider: %v", err)
        return nil, err
	}

	volumeID := req.VolumeId
    err = rsdClient.DeleteVolume(volumeID)
    if err != nil {
        glog.V(3).Infof("Failed to DeleteVolume: %v", err)
        return nil, err
	}

	glog.V(4).Infof("delete volume %s", volumeID)
	return &csi.DeleteVolumeResponse{}, nil
}


func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
    // Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if req.GetVolumeCapabilities() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities missing in request")
	}
//	if _, ok := hostPathVolumes[req.GetVolumeId()]; !ok {
//		return nil, status.Error(codes.NotFound, "Volume does not exist")
//	}

	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Supported: false, Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{Supported: true, Message: ""}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
    return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
    // Start to detach
    return &csi.ControllerUnpublishVolumeResponse{}, nil
}

