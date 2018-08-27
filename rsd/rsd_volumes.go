package rsd

import (
	"fmt"
    "strings"
//    "net/http"
//    "encoding/json"
//	"time"
//    "log"
)

type CreateOpts struct {
    Name string
    Size int64
}

func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
    v := make(map[string]interface{})

    if opts.Size == 0 {
        return nil, fmt.Errorf("Required CreateOpts field 'Size' not set.")
    }
    v["CapacityBytes"] = opts.Size
    return v, nil
}

func (client *ServiceClient) CreateVolume(name string, size int64) (string, error) {
    opts := &CreateOpts{
        Name : name,
        Size:  size,
    }
    reqBody, err := opts.ToVolumeCreateMap()
    if err != nil {
        return "", err
    }
    url := createURL(client)
    resp, err := client.ProviderClient.Post(url, reqBody, nil, &RequestOpts{
        OkCodes: []int{201},
    })
    fmt.Printf("%q", resp)
    if err == nil {
        volume_url := strings.Join(resp.Header["Location"]," ")
        return volume_url, nil
    }
    return "", err
}

