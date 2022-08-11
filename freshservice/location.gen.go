package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"net/http"
	"net/url"
)

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

// List all locations
func (d *LocationsServiceClient) List(ctx context.Context, filter QueryFilter) ([]LocationDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   locationURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Locations{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
