package rsd

import (
	"fmt"
    "strings"
    "io/ioutil"
//    "net/http"
    "encoding/json"
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

func (client *ServiceClient) GetVolume(volume_id string) (string, error) {
    url := getURL(client, volume_id)
    resp, err := client.ProviderClient.Get(url, nil, &RequestOpts{
        OkCodes: []int{200},
    })
    body, err := ioutil.ReadAll(resp.Body)
    var dat map[string]interface{}

    json.Unmarshal(body, &dat)
    
    for key, value := range dat {
        fmt.Printf("%s=\"%s\"\n", key, value)
    }

    // Do update
    return "dad", err
}

func (client *ServiceClient) ListVolume() {
}

func (client *ServiceClient) AttachVolume(node_uri string, volume_id string) (error) {
    reqBody := make(map[string]interface{})
    resource := make(map[string]interface{})
    resource["@odata.id"] = "/redfish/v1/StorageServices/5-sv-1/Volumes/" + volume_id
    reqBody["Resource"] = resource
    resp, err := client.ProviderClient.Post(node_uri, reqBody, nil, &RequestOpts{
        OkCodes: []int{201},
    })

    var dat map[string]interface{}

    body, err := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &dat)

    for key, value := range dat {
        fmt.Printf("%s=\"%s\"\n", key, value)
    }

    return err
}

func (client *ServiceClient) DeleteVolume(volume_id string) (error) {

    url := getURL(client, volume_id)

    resp, err := client.ProviderClient.Delete(url, &RequestOpts{
        OkCodes: []int{204, 202},
    })

    fmt.Printf("%q", resp)
    // nil if delete OK
    return err
}
