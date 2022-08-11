package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"net/http"
	"net/url"
)

const applicationURL = "/api/v2/applications"

// Applications holds a list of Freshservice Application details
type Applications struct {
	List []ApplicationDetails `json:"applications"`
}

// Application holds the details of a specific Freshservice Application
type Application struct {
	Details ApplicationDetails `json:"application"`
}

// Applications is the interface between the HTTP client and the Freshservice application related endpoints
func (fs *Client) Applications() ApplicationsService {
	return &ApplicationsServiceClient{client: fs}
}

// ApplicationsServiceClient facilitates requests with the ApplicationsService methods
type ApplicationsServiceClient struct {
	client *Client
}

// List all applications
func (d *ApplicationsServiceClient) List(ctx context.Context, filter QueryFilter) ([]ApplicationDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   applicationURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Applications{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
