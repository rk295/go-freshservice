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

// List all Assets
// Append the parameter "page=[:page_no]" in the url to traverse through pages.
// TODO: this needs to have filtering added: https://api.freshservice.com/#filter_assets
func (a *AssetsServiceClient) List(ctx context.Context, filter QueryFilter) ([]AssetDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   assetURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Assets{}
	resp, err := a.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
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
