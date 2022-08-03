package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const requesterURL = "/api/v2/requesters"

// RequesterService is an interface for interacting with
// the requester endpoints of the Freshservice API
type RequesterService interface {
	List(context.Context, QueryFilter) ([]RequesterDetails, string, error)
	Get(context.Context, int, QueryFilter) (*RequesterDetails, error)
}

// RequesterServiceClient facilitates requests with the Freshservice requester methods
type RequesterServiceClient struct {
	client *Client
}

// List all Freshservice requesters
func (r *RequesterServiceClient) List(ctx context.Context, filter QueryFilter) ([]RequesterDetails, string, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   r.client.Domain,
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
	resp, err := r.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}
	return res.List, HasNextPage(resp), nil
}

// Get a specific requester
func (r *RequesterServiceClient) Get(ctx context.Context, reqID int, filter QueryFilter) (*RequesterDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   r.client.Domain,
		Path:   fmt.Sprintf("%s/%d", requesterURL, reqID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	_, err = r.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}
	return &res.Details, nil
}
