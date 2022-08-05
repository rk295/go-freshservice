package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const assetTypesURL = "/api/v2/asset_types"

// AssetTypeService is an interface for interacting with
// the Asset Type endpoints of the Freshservice API
type AssetTypeService interface {
	List(context.Context, QueryFilter) ([]AssetTypeDetails, string, error)
	Get(context.Context, int, QueryFilter) (*AssetTypeDetails, error)
}

// AssetTypeServiceClient facilitates requests with the AssetTypeService methods
type AssetTypeServiceClient struct {
	client *Client
}

// List all AssetsTypes
func (a *AssetTypeServiceClient) List(ctx context.Context, filter QueryFilter) ([]AssetTypeDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   assetTypesURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &AssetTypes{}
	resp, err := a.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific AssetsType
func (a *AssetTypeServiceClient) Get(ctx context.Context, assetTypeID int, filter QueryFilter) (*AssetTypeDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", assetTypesURL, assetTypeID),
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &AssetType{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
