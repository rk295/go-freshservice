package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const assetURL = "/api/v2/assets"

// Assets holds a list of Freshservice Asset details
type Assets struct {
	List []AssetDetails `json:"assets"`
}

// Asset holds the details of a specific Freshservice Asset
type Asset struct {
	Details AssetDetails `json:"asset"`
}

// Assets is the interface between the HTTP client and the Freshservice asset related endpoints
func (fs *Client) Assets() AssetsService {
	return &AssetsServiceClient{client: fs}
}

// AssetsServiceClient facilitates requests with the AssetsService methods
type AssetsServiceClient struct {
	client *Client
}

// List all assets
func (d *AssetsServiceClient) List(ctx context.Context, filter QueryFilter) ([]AssetDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
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
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific asset
func (d *AssetsServiceClient) Get(ctx context.Context, id int) (*AssetDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   path.Join(assetURL, fmt.Sprintf("%d", id)),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Asset{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
