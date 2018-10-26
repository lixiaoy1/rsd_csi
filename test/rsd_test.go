package rsd_csi

import (
    "fmt"
    "testing"
)

func TestAttachVolume(t *testing.T) {
    rsdClient, _ := GetRSDProvider()
    rsdClient.AttachVolume("1",
            "5-sv-1-vl-5")
}

func TestGetVolume(t *testing.T) {
    rsdClient, _ := GetRSDProvider()
    volume, _ := rsdClient.GetVolume("5-sv-1-vl-5")
    fmt.Printf("volume id %v\n", volume.ID)
}
