package rsd_csi

import (
    "github.com/rsd_csi/rsd"

)

var rsd_client *rsd.ServiceClient = nil

func GetRSDProvider() (*rsd.ServiceClient, error) {
    if rsd_client == nil {
        //
        provider, err := NewClient("")
        if err != nil {
            return nil, err
        }
        rsd_client, err = NewServiceClient(provider, "")
        if err != nil {
            return nil, err
        }
    }
    return rsd_client, nil
}

func NewClient(endpoint string) (*rsd.ProviderClient, error) {
    p := new(rsd.ProviderClient)
    p.IdentityBase = "https://podm-otc.jf.intel.com:8443"
    p.IdentityEndpoint = endpoint
    p.User = "admin"
    p.Password = "admin"

    return p, nil
}

func NewServiceClient(client *rsd.ProviderClient, url string) (*rsd.ServiceClient, error) {
    sc := new(rsd.ServiceClient)
    sc.ProviderClient = client
    sc.Endpoint = client.IdentityBase
    return sc, nil
}
