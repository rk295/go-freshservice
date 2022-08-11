package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"net/http"
	"net/url"
)

const requesterURL = "/api/v2/requesters"

// Requesters holds a list of Freshservice Requester details
type Requesters struct {
	List []RequesterDetails `json:"requesters"`
}

// Requester holds the details of a specific Freshservice Requester
type Requester struct {
	Details RequesterDetails `json:"requester"`
}

// Requesters is the interface between the HTTP client and the Freshservice requester related endpoints
func (fs *Client) Requesters() RequestersService {
	return &RequestersServiceClient{client: fs}
}

// RequestersServiceClient facilitates requests with the RequestersService methods
type RequestersServiceClient struct {
	client *Client
}

// List all requesters
func (d *RequestersServiceClient) List(ctx context.Context, filter QueryFilter) ([]RequesterDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   requesterURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Requesters{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
