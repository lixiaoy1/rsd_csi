package rsd


func createURL(c *ServiceClient) string {
	return c.ServiceURL("/redfish/v1/StorageServices/6-sv-1/Volumes")
}

func listURL(c *ServiceClient) string {
	return c.ServiceURL("/redfish/v1/StorageServices/6-sv-1/Volumes")
}

func deleteURL(c *ServiceClient, id string) string {
	return c.ServiceURL("/redfish/v1/StorageServices/6-sv-1/Volumes", id)
}

func getURL(c *ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *ServiceClient, id string) string {
	return deleteURL(c, id)
}

func getEndpointURL(c *ServiceClient, endpoint_path string) string {
    return c.ServiceURL(endpoint_path)
}

func getSystemURL(c *ServiceClient, system_path string) string {
    return c.ServiceURL(system_path)
}

func getNodeURL(c *ServiceClient, id string) string {
    return c.ServiceURL("/redfish/v1/Nodes", id)
}

func getNodeAttachURL(c *ServiceClient, id string) string {
    return c.ServiceURL("/redfish/v1/Nodes", id, "Actions/ComposedNode.AttachResource")
}

func getNodeDetachURL(c *ServiceClient, id string) string {
    return c.ServiceURL("/redfish/v1/Nodes", id, "Actions/ComposedNode.DetachResource")
}

