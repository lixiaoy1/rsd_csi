package rsd

import (
    "fmt"
    "testing"
)

func TestCreateVolume(t *testing.T) {
    resp := CreateVolume("lisatest", 1)
    fmt.Printf("error: %#v\n", resp)
}
