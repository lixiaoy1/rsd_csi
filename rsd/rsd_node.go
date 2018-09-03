package rsd

type Node struct {
	OdataContext string      `json:"@odata.context"`
	OdataID      string      `json:"@odata.id"`
	OdataType    string      `json:"@odata.type"`
	ID           string      `json:"Id"`
	Name         string      `json:"Name"`
	Description  interface{} `json:"Description"`
	UUID         string      `json:"UUID"`
	PowerState   string      `json:"PowerState"`
	Status       struct {
		State        string `json:"State"`
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
	} `json:"Status"`
	ComposedNodeState string `json:"ComposedNodeState"`
	Boot              struct {
		BootSourceOverrideEnabled                      string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideTarget                       string   `json:"BootSourceOverrideTarget"`
		BootSourceOverrideTargetRedfishAllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`
		BootSourceOverrideMode                         string   `json:"BootSourceOverrideMode"`
		BootSourceOverrideModeRedfishAllowableValues   []string `json:"BootSourceOverrideMode@Redfish.AllowableValues"`
	} `json:"Boot"`
	ClearTPMOnDelete bool `json:"ClearTPMOnDelete"`
	Links            struct {
		ComputerSystem struct {
			OdataID string `json:"@odata.id"`
		} `json:"ComputerSystem"`
		Processors []struct {
			OdataID string `json:"@odata.id"`
		} `json:"Processors"`
		Memory []struct {
			OdataID string `json:"@odata.id"`
		} `json:"Memory"`
		EthernetInterfaces []struct {
			OdataID string `json:"@odata.id"`
		} `json:"EthernetInterfaces"`
		Storage []struct {
			OdataID string `json:"@odata.id"`
		} `json:"Storage"`
		ManagedBy []struct {
			OdataID string `json:"@odata.id"`
		} `json:"ManagedBy"`
		Oem struct {
		} `json:"Oem"`
	} `json:"Links"`
	Actions struct {
		ComposedNodeReset struct {
			Target                          string   `json:"target"`
			ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#ComposedNode.Reset"`
		ComposedNodeAssemble struct {
			Target string `json:"target"`
		} `json:"#ComposedNode.Assemble"`
		ComposedNodeAttachResource struct {
			Target            string `json:"target"`
			RedfishActionInfo struct {
				OdataID string `json:"@odata.id"`
			} `json:"@Redfish.ActionInfo"`
		} `json:"#ComposedNode.AttachResource"`
		ComposedNodeDetachResource struct {
			Target            string `json:"target"`
			RedfishActionInfo struct {
				OdataID string `json:"@odata.id"`
			} `json:"@Redfish.ActionInfo"`
		} `json:"#ComposedNode.DetachResource"`
	} `json:"Actions"`
	Oem struct {
	} `json:"Oem"`
}

type System struct {
	OdataContext string      `json:"@odata.context"`
	OdataID      string      `json:"@odata.id"`
	OdataType    string      `json:"@odata.type"`
	ID           string      `json:"Id"`
	Name         string      `json:"Name"`
	Description  string      `json:"Description"`
	SystemType   string      `json:"SystemType"`
	AssetTag     interface{} `json:"AssetTag"`
	Manufacturer string      `json:"Manufacturer"`
	Model        string      `json:"Model"`
	SKU          string      `json:"SKU"`
	SerialNumber string      `json:"SerialNumber"`
	PartNumber   string      `json:"PartNumber"`
	UUID         string      `json:"UUID"`
	HostName     interface{} `json:"HostName"`
	Status       struct {
		State        string `json:"State"`
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
	} `json:"Status"`
	IndicatorLED interface{} `json:"IndicatorLED"`
	PowerState   string      `json:"PowerState"`
	Boot         struct {
		OdataType                                      string   `json:"@odata.type"`
		BootSourceOverrideEnabled                      string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideTarget                       string   `json:"BootSourceOverrideTarget"`
		BootSourceOverrideTargetRedfishAllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`
		BootSourceOverrideMode                         string   `json:"BootSourceOverrideMode"`
		BootSourceOverrideModeRedfishAllowableValues   []string `json:"BootSourceOverrideMode@Redfish.AllowableValues"`
	} `json:"Boot"`
	BiosVersion      string `json:"BiosVersion"`
	ProcessorSummary struct {
		Count  int    `json:"Count"`
		Model  string `json:"Model"`
		Status struct {
			State        string `json:"State"`
			Health       string `json:"Health"`
			HealthRollup string `json:"HealthRollup"`
		} `json:"Status"`
	} `json:"ProcessorSummary"`
	MemorySummary struct {
		TotalSystemMemoryGiB float64 `json:"TotalSystemMemoryGiB"`
		Status               struct {
			State        string `json:"State"`
			Health       string `json:"Health"`
			HealthRollup string `json:"HealthRollup"`
		} `json:"Status"`
	} `json:"MemorySummary"`
	Processors struct {
		OdataID string `json:"@odata.id"`
	} `json:"Processors"`
	EthernetInterfaces struct {
		OdataID string `json:"@odata.id"`
	} `json:"EthernetInterfaces"`
	SimpleStorage struct {
		OdataID string `json:"@odata.id"`
	} `json:"SimpleStorage"`
	Storage struct {
		OdataID string `json:"@odata.id"`
	} `json:"Storage"`
	Memory struct {
		OdataID string `json:"@odata.id"`
	} `json:"Memory"`
	PCIeDevices []struct {
		OdataID string `json:"@odata.id"`
	} `json:"PCIeDevices"`
	PCIeFunctions     []interface{} `json:"PCIeFunctions"`
	NetworkInterfaces struct {
		OdataID string `json:"@odata.id"`
	} `json:"NetworkInterfaces"`
	TrustedModules []struct {
		OdataType              string      `json:"@odata.type"`
		FirmwareVersion        interface{} `json:"FirmwareVersion"`
		FirmwareVersion2       interface{} `json:"FirmwareVersion2"`
		InterfaceType          string      `json:"InterfaceType"`
		InterfaceTypeSelection string      `json:"InterfaceTypeSelection"`
		Status                 struct {
			State        string      `json:"State"`
			Health       interface{} `json:"Health"`
			HealthRollup interface{} `json:"HealthRollup"`
		} `json:"Status"`
		Oem struct {
		} `json:"Oem"`
	} `json:"TrustedModules"`
	HostingRoles   []interface{} `json:"HostingRoles"`
	HostedServices struct {
		StorageServices []interface{} `json:"StorageServices"`
	} `json:"HostedServices"`
	Links struct {
		OdataType string `json:"@odata.type"`
		Chassis   []struct {
			OdataID string `json:"@odata.id"`
		} `json:"Chassis"`
		ManagedBy []struct {
			OdataID string `json:"@odata.id"`
		} `json:"ManagedBy"`
		Endpoints []struct {
			OdataID string `json:"@odata.id"`
		} `json:"Endpoints"`
		Oem struct {
		} `json:"Oem"`
	} `json:"Links"`
	Actions struct {
		ComputerSystemReset struct {
			Target                          string   `json:"target"`
			ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#ComputerSystem.Reset"`
		Oem struct {
			IntelOemStartDeepDiscovery struct {
				Target string `json:"target"`
			} `json:"#Intel.Oem.StartDeepDiscovery"`
			IntelOemChangeTPMState struct {
				Target                              string   `json:"target"`
				InterfaceTypeRedfishAllowableValues []string `json:"InterfaceType@Redfish.AllowableValues"`
			} `json:"#Intel.Oem.ChangeTPMState"`
		} `json:"Oem"`
	} `json:"Actions"`
	Oem struct {
		IntelRackScale struct {
			OdataType  string `json:"@odata.type"`
			PciDevices []struct {
				VendorID string `json:"VendorId"`
				DeviceID string `json:"DeviceId"`
			} `json:"PciDevices"`
			PCIeConnectionID                  []interface{} `json:"PCIeConnectionId"`
			DiscoveryState                    string        `json:"DiscoveryState"`
			ProcessorSockets                  int           `json:"ProcessorSockets"`
			MemorySockets                     int           `json:"MemorySockets"`
			UserModeEnabled                   bool          `json:"UserModeEnabled"`
			TrustedExecutionTechnologyEnabled bool          `json:"TrustedExecutionTechnologyEnabled"`
			Metrics                           struct {
				OdataID string `json:"@odata.id"`
			} `json:"Metrics"`
		} `json:"Intel_RackScale"`
	} `json:"Oem"`
}
