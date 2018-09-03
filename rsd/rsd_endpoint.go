package rsd

type Endpoint struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Status       struct {
		State        string `json:"State"`
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
	} `json:"Status"`
	EndpointProtocol string      `json:"EndpointProtocol"`
	PciID            interface{} `json:"PciId"`
	Identifiers      []struct {
		DurableName       string `json:"DurableName"`
		DurableNameFormat string `json:"DurableNameFormat"`
	} `json:"Identifiers"`
	ConnectedEntities []struct {
		EntityRole        string      `json:"EntityRole"`
		PciFunctionNumber interface{} `json:"PciFunctionNumber"`
		PciClassCode      interface{} `json:"PciClassCode"`
		EntityPciID       interface{} `json:"EntityPciId"`
		EntityLink        struct {
			OdataID string `json:"@odata.id"`
		} `json:"EntityLink"`
		Identifiers []interface{} `json:"Identifiers"`
	} `json:"ConnectedEntities"`
	Redundancy         []interface{} `json:"Redundancy"`
	IPTransportDetails []struct {
		TransportProtocol string `json:"TransportProtocol"`
		IPv4Address       struct {
			Address string `json:"Address"`
			Oem     struct {
			} `json:"Oem"`
		} `json:"IPv4Address"`
		Port int `json:"Port"`
	} `json:"IPTransportDetails"`
	HostReservationMemoryBytes interface{} `json:"HostReservationMemoryBytes"`
	Links                      struct {
		OdataType string        `json:"@odata.type"`
		Ports     []interface{} `json:"Ports"`
		Oem       struct {
			IntelRackScale struct {
				Zones []struct {
					OdataID string `json:"@odata.id"`
				} `json:"Zones"`
				Interfaces []struct {
					OdataID string `json:"@odata.id"`
				} `json:"Interfaces"`
			} `json:"Intel_RackScale"`
		} `json:"Oem"`
	} `json:"Links"`
	Actions struct {
		Oem struct {
		} `json:"Oem"`
	} `json:"Actions"`
	Oem struct {
		IntelRackScale struct {
			OdataType      string `json:"@odata.type"`
			Authentication struct {
				Username interface{} `json:"Username"`
				Password interface{} `json:"Password"`
			} `json:"Authentication"`
		} `json:"Intel_RackScale"`
	} `json:"Oem"`
}

