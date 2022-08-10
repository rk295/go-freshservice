package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const locationURL = "/api/v2/locations"

// LocationService is an interface for interacting with the location
// endpoints of the Freshservice API
type LocationsService interface {
	List(context.Context, QueryFilter) ([]LocationDetails, string, error)
	Get(context.Context, int, QueryFilter) (*LocationDetails, error)
}

// locationServiceClient facilitates requests with the locationService methods
type locationServiceClient struct {
	client *Client
}

// List all locations
func (d *locationServiceClient) List(ctx context.Context, filter QueryFilter) ([]LocationDetails, string, error) {

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

// Get a specific location
func (d *locationServiceClient) Get(ctx context.Context, locID int, filter QueryFilter) (*LocationDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   fmt.Sprintf("%s/%d", locationURL, locID),
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Location{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}