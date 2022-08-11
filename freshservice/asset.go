package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AssetsService is an interface for interacting with
// the asset endpoints of the Freshservice API
type AssetsService interface {
	List(context.Context, QueryFilter) ([]AssetDetails, string, error)
	Get(context.Context, int, QueryFilter) (*AssetDetails, error)
}

// Get a specific asset
func (a *AssetsServiceClient) Get(ctx context.Context, assetID int, filter QueryFilter) (*AssetDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", assetURL, assetID),
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Asset{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
