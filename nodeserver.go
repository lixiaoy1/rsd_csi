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
    "os"
//    "fmt"
//    "strings"
    "time"

	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"

	"github.com/kubernetes-csi/drivers/pkg/csi-common"
    "k8s.io/kubernetes/pkg/util/mount"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
    //node_uri string

}


var attachRequested = false
var nvmeConnected = false

var initiator = ""
var target = ""
var target_ip = ""

func (ns *nodeServer) NodeGetId(ctx context.Context, req *csi.NodeGetIdRequest) (*csi.NodeGetIdResponse, error) {
    glog.Errorf("NodeGetId enter\n [Context]: %v,\n [Request]: %v.\n", ctx, *req)
    return ns.DefaultNodeServer.NodeGetId(ctx, req)
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
    glog.Errorf("NodeGetInfo enter\n [Context]: %v,\n [Request]: %v.\n", ctx, *req)
    return ns.DefaultNodeServer.NodeGetInfo(ctx, req)
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	glog.Errorf("NodePublishVolume enter\n [Context]: %v,\n [Request]: %v.\n", ctx, *req)
	//panic("NodePublishVolume Stack Trace:")
        targetPath := req.GetTargetPath()
	volName := req.GetVolumeId()

	/*if !strings.HasPrefix(targetPath, "/mnt") {
		return nil, fmt.Errorf("rsd: malformed the value of target path: %s", targetPath)
	}
	s := strings.Split(strings.TrimSuffix(targetPath, "/mount"), "/")
	volName := s[len(s)-1]*/
	notMnt, err := mount.New("").IsLikelyNotMountPoint(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if !notMnt {
		return &csi.NodePublishVolumeResponse{}, nil
	}

    // Start to attach
    rsdClient, err := GetRSDProvider()
    if err != nil {
        glog.V(3).Infof("Failed to GetRSDProvider: %v", err)
        return nil, err
    }

    // Attach in storage
    if attachRequested == false {
        glog.Infof("__tingjie AttachVolume RESTful request ...")
        initiator, target, target_ip, err = rsdClient.AttachVolume("1", volName)
        if err != nil {
            glog.V(3).Infof("Failed to attach: %v", err)
            return nil, err
        }

        glog.Infof("__tingjie AttachVolume RESTful request success! initiator:%s, target:%s, target_ip:%s", initiator, target, target_ip)
        attachRequested = true
    }

    // nvme connect
    if nvmeConnected == false {
        glog.Infof("__tingjie connectRSDVolume CLI request ...")
        err = connectRSDVolume(initiator, target, target_ip)
        if err != nil {
            glog.Errorf("Failed to find the device: %v", err)
            return nil, err
        }

        glog.Infof("__tingjie connectRSDVolume CLI request success!")
        nvmeConnected = true
    }

    time.Sleep(5 * time.Second)
    devicePath := getDevicePath2(target)
    if devicePath == "" {
        glog.Errorf("Failed to getDevicePath")
        return nil, err
    }
    glog.Errorf("Success to get devicePath: %v\n", devicePath)

    // Mount
    //devicePath := req.GetPublishInfo()["DevicePath"]
    fsType := req.GetVolumeCapability().GetMount().GetFsType()
	readOnly := req.GetReadonly()
	attrib := req.GetVolumeAttributes()
	mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()

	glog.Errorf("target %v\nfstype %v\ndevice %v\nreadonly %v\nattributes %v\n mountflags %v\n",
		targetPath, fsType, devicePath, readOnly, attrib, mountFlags)

	options := []string{}
	if readOnly {
		options = append(options, "ro")
	}

    diskMounter := &mount.SafeFormatAndMount{Interface: mount.New(""), Exec: mount.NewOsExec()}
	if err := diskMounter.FormatAndMount(devicePath, targetPath, fsType, options); err != nil {
		return nil, err
	}
    return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	glog.Errorf("NodeUnpublishVolume enter\n [Context]: %v,\n [Request]: %v.\n", ctx, *req)
	targetPath := req.GetTargetPath()

	/*if !strings.HasPrefix(targetPath, "/mnt") {
		return nil, fmt.Errorf("rsd: malformed the value of target path: %s", targetPath)
	}
	s := strings.Split(strings.TrimSuffix(targetPath, "/mount"), "/")
	volName := s[len(s)-1]*/

    volName := req.GetVolumeId()
    glog.V(4).Infof("unmount targetPath %v\n", targetPath)
    err := mount.New("").Unmount(req.GetTargetPath())
    if err != nil {
        return nil, err
    }

    // Start to detach
    rsdClient, err := GetRSDProvider()
    if err != nil {
        glog.V(3).Infof("Failed to GetRSDProvider: %v", err)
        return nil, err
    }

    // Detach in storage
    err = rsdClient.DetachVolume("1", volName)
    if err != nil {
        glog.V(3).Infof("Failed to detach: %v", err)
        return nil, err
    }

    return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {

	return &csi.NodeUnstageVolumeResponse{}, nil
}


