package rsd


func createURL(c *ServiceClient) string {
	return c.ServiceURL("redfish/v1/StorageServices/5-sv-1/Volumes")
}

func listURL(c *ServiceClient) string {
	return c.ServiceURL("redfish/v1/StorageServices/5-sv-1/Volumes")
}

func deleteURL(c *ServiceClient, id string) string {
	return c.ServiceURL("redfish/v1/StorageServices/5-sv-1/Volumes", id)
}

func getURL(c *ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *ServiceClient, id string) string {
	return deleteURL(c, id)
}
