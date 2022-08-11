package freshservice

// Generated Code DO NOT EDIT

const locationURL = "/api/v2/locations"

// Locations holds a list of Freshservice Location details
type Locations struct {
	List []LocationDetails `json:"locations"`
}

// Location holds the details of a specific Freshservice Location
type Location struct {
	Details LocationDetails `json:"location"`
}

// Locations is the interface between the HTTP client and the Freshservice location related endpoints
func (fs *Client) Locations() LocationsService {
	return &LocationsServiceClient{client: fs}
}

// LocationsServiceClient facilitates requests with the LocationsService methods
type LocationsServiceClient struct {
	client *Client
}
