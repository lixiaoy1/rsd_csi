package rsd_csi

import (
    "fmt"
    "testing"
)

func TestGetVolume(t *testing.T) {
    rsdClient, err := GetRSDProvider()

    _, err = rsdClient.GetVolume("1-sv-1-vl-5")
    fmt.Printf("error: %#v\n", err)
}

func TestAttachVolume(t *testing.T) {
    rsdClient, err := GetRSDProvider()
    err =
    rsdClient.AttachVolume("https://podm-otc.jf.intel.com:8443/redfish/v1/Nodes/1/Actions/ComposedNode.AttachResource",
            "1-sv-1-vl-5")
    fmt.Printf("error: %#v\n", err)
}
