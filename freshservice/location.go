package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// LocationsService is an interface for interacting with the location
// endpoints of the Freshservice API
type LocationsService interface {
	List(context.Context, QueryFilter) ([]LocationDetails, string, error)
	Get(context.Context, int, QueryFilter) (*LocationDetails, error)
}

// Get a specific location
func (d *LocationsServiceClient) Get(ctx context.Context, locID int, filter QueryFilter) (*LocationDetails, error) {

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
