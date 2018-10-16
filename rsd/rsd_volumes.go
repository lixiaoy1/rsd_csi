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

type Volume struct {
	OdataContext       string      `json:"@odata.context"`
	OdataID            string      `json:"@odata.id"`
	OdataType          string      `json:"@odata.type"`
	Description        string      `json:"Description"`
	ID                 string      `json:"Id"`
	Model              interface{} `json:"Model"`
	Manufacturer       interface{} `json:"Manufacturer"`
	Name               string      `json:"Name"`
	AccessCapabilities []string    `json:"AccessCapabilities"`
	CapacityBytes      int64       `json:"CapacityBytes"`
	Actions            struct {
		VolumeInitialize struct {
			Target string `json:"target"`
		} `json:"#Volume.Initialize"`
		Oem struct {
		} `json:"Oem"`
	} `json:"Actions"`
	Capacity struct {
		Data struct {
			AllocatedBytes int64 `json:"AllocatedBytes"`
		} `json:"Data"`
	} `json:"Capacity"`
	CapacitySources []struct {
		ProvidingPools []struct {
			OdataID string `json:"@odata.id"`
		} `json:"ProvidingPools"`
		ProvidedCapacity struct {
			Data struct {
				AllocatedBytes int64 `json:"AllocatedBytes"`
			} `json:"Data"`
		} `json:"ProvidedCapacity"`
	} `json:"CapacitySources"`
	Identifiers []struct {
		DurableName       string `json:"DurableName"`
		DurableNameFormat string `json:"DurableNameFormat"`
	} `json:"Identifiers"`
	Links struct {
		Oem struct {
			IntelRackScale struct {
				OdataType string `json:"@odata.type"`
				Endpoints []struct {
					OdataID string `json:"@odata.id"`
				} `json:"Endpoints"`
				Metrics struct {
					OdataID string `json:"@odata.id"`
				} `json:"Metrics"`
			} `json:"Intel_RackScale"`
		} `json:"Oem"`
		Drives []interface{} `json:"Drives"`
	} `json:"Links"`
	ReplicaInfos []struct {
		ReplicaReadOnlyAccess string `json:"ReplicaReadOnlyAccess"`
		ReplicaType           string `json:"ReplicaType"`
		ReplicaRole           string `json:"ReplicaRole"`
		Replica               struct {
			OdataID string `json:"@odata.id"`
		} `json:"Replica"`
	} `json:"ReplicaInfos"`
	Status struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
	Oem struct {
		IntelRackScale struct {
			OdataType     string      `json:"@odata.type"`
			Bootable      bool        `json:"Bootable"`
			Erased        interface{} `json:"Erased"`
			EraseOnDetach bool        `json:"EraseOnDetach"`
		} `json:"Intel_RackScale"`
	} `json:"Oem"`
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
        vols := strings.Split(volume_url, "/")
        return vols[len(vols)-1], nil
    }
    return "", err
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

func (client *ServiceClient) GetVolume(volume_id string) (*Volume, error) {
    url := getURL(client, volume_id)
    resp, err := client.ProviderClient.Get(url, nil, &RequestOpts{
        OkCodes: []int{200},
    })
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    var volume Volume
    json.Unmarshal(body, &volume)

    // Do update
    return &volume, err
}

func (client *ServiceClient) ListVolume() {
}

func (client *ServiceClient) GetEndpoint(uri string) (*Endpoint, error) {
    url := getEndpointURL(client, uri)
    resp, err := client.ProviderClient.Get(url, nil, &RequestOpts{
        OkCodes: []int{200},
    })
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    var end Endpoint
    json.Unmarshal(body, &end)

    // Do update
    return &end, err
}

func (client *ServiceClient) GetSystem(uri string) (*System, error) {
    url := getSystemURL(client, uri)
    resp, err := client.ProviderClient.Get(url, nil, &RequestOpts{
        OkCodes: []int{200},
    })
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    var system System
    json.Unmarshal(body, &system)

    // Do update
    return &system, err
}


func (client *ServiceClient) GetNode(id string) (*Node, error) {
    url := getNodeURL(client, id)
    resp, err := client.ProviderClient.Get(url, nil, &RequestOpts{
        OkCodes: []int{200},
    })
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    var node Node
    json.Unmarshal(body, &node)

    // Do update
    return &node, err
}

func GetNQN(endpoint *Endpoint) (string) {
    identifiers := endpoint.Identifiers
    for i:=0; i < len(identifiers); i++ {
        iden := identifiers[i]
        if iden.DurableNameFormat == "NQN"  {
            return strings.Split(iden.DurableName, ":")[2]
        }
    }
    return ""
}

func GetTargetIP(endpoint *Endpoint) (string, int) {
    transports := endpoint.IPTransportDetails
    for i := 0; i < len(transports); i++ {
        tr := transports[i]
        // Add all protocols supporting NVMeoF
        if tr.TransportProtocol == "RoCEv2" {
            address := tr.IPv4Address.Address
            port := tr.Port
            return address, port
        }
    }
    return "", 0
}

func (client *ServiceClient) GetTargetNQN(volume_id string) (string, string, int, error) {
    volume, err := client.GetVolume(volume_id)
    if err != nil {
        return "", "", 0, err
    }
    volume_end := volume.Links.Oem.IntelRackScale.Endpoints[0].OdataID
    volume_endpoint, err := client.GetEndpoint(volume_end)
    if err != nil {
        return "", "", 0, err
    }
    target_nqn:= GetNQN(volume_endpoint)

    target_ip, port := GetTargetIP(volume_endpoint)
    return target_nqn, target_ip, port, nil
}

func (client *ServiceClient) GetNodeNQN(node_id string) (string, error) {
    node, err := client.GetNode(node_id)
    if err != nil {
        return "", err
    }
    node_sys,err := client.GetSystem(node.Links.ComputerSystem.OdataID)
    if err != nil {
        return "", err
    }

    for i := 0; i < len(node_sys.Links.Endpoints); i++ {
        node_endpoint, err := client.GetEndpoint(node_sys.Links.Endpoints[i].OdataID)
        if err != nil {
            continue
        }
        initiator_nqn := GetNQN(node_endpoint)
        if initiator_nqn != "" {
            return initiator_nqn, nil
        }
    }

    // TODO: handle no NVMeoF connection
    return "", nil
}


func (client *ServiceClient) AttachVolume(node_id string, volume_id string) (string, string, string, error) {
    reqBody := make(map[string]interface{})
    resource := make(map[string]interface{})
    resource["@odata.id"] = "/redfish/v1/StorageServices/6-sv-1/Volumes/" + volume_id
    reqBody["Resource"] = resource
    node_attach_url := getNodeAttachURL(client, node_id)
    _, err := client.ProviderClient.Post(node_attach_url, reqBody, nil, &RequestOpts{
        OkCodes: []int{204},
    })
    if err != nil {
        return "", "", "", err
    }

    // Get target nqn
    target_nqn, target_ip, port, err := client.GetTargetNQN(volume_id)
    if err != nil {
        return "", "", "", err
    }
    fmt.Printf("%v", port)

    // Get initiator nqn
    node_nqn, err := client.GetNodeNQN(node_id)
    if err != nil {
        return "", "", "", err
    }

    return node_nqn, target_nqn, target_ip, nil
}

func (client *ServiceClient) DetachVolume(node_id string, volume_id string) (error) {
    reqBody := make(map[string]interface{})
    resource := make(map[string]interface{})
    resource["@odata.id"] = "/redfish/v1/StorageServices/6-sv-1/Volumes/" + volume_id
    reqBody["Resource"] = resource
    node_attach_url := getNodeDetachURL(client, node_id)
    _, err := client.ProviderClient.Post(node_attach_url, reqBody, nil, &RequestOpts{
        OkCodes: []int{204},
    })
    if err != nil {
        return err
    }

    return nil
}

