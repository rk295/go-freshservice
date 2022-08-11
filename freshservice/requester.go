package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// RequestersService is an interface for interacting with
// the requester endpoints of the Freshservice API
type RequestersService interface {
	List(context.Context, QueryFilter) ([]RequesterDetails, string, error)
	Get(context.Context, int, QueryFilter) (*RequesterDetails, error)
}

// Get a specific requester
func (r *RequestersServiceClient) Get(ctx context.Context, reqID int, filter QueryFilter) (*RequesterDetails, error) {
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
