package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const assetTypeURL = "/api/v2/asset_types"

// AssetTypes holds a list of Freshservice AssetType details
type AssetTypes struct {
	List []AssetTypeDetails `json:"asset_types"`
}

// AssetType holds the details of a specific Freshservice AssetType
type AssetType struct {
	Details AssetTypeDetails `json:"asset_types"`
}

// AssetTypes is the interface between the HTTP client and the Freshservice assetType related endpoints
func (fs *Client) AssetTypes() AssetTypesService {
	return &AssetTypesServiceClient{client: fs}
}

// AssetTypesServiceClient facilitates requests with the AssetTypesService methods
type AssetTypesServiceClient struct {
	client *Client
}

// List all assetTypes
func (d *AssetTypesServiceClient) List(ctx context.Context, filter QueryFilter) ([]AssetTypeDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   assetTypeURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &AssetTypes{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific assetType
func (d *AssetTypesServiceClient) Get(ctx context.Context, id int) (*AssetTypeDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   path.Join(assetTypeURL, fmt.Sprintf("%d", id)),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &AssetType{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
