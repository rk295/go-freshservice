package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// RequestersService is an interface for interacting with
// the requester endpoints of the Freshservice API
type RequestersService interface {
	List(context.Context, QueryFilter) ([]RequesterDetails, string, error)
	Get(context.Context, int) (*RequesterDetails, error)
	Update(context.Context, int, *RequesterUpdateDetails) (*RequesterDetails, error)
}

// Update a Freshservice requster
func (as *RequestersServiceClient) Update(ctx context.Context, id int, ad *RequesterUpdateDetails) (*RequesterDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", requesterURL, id),
	}

	requesterContent, err := json.Marshal(ad)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(requesterContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &RequesterDetails{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil

}
