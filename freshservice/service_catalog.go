package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	serviceCatalogItemURL     = "/api/v2/service_catalog/items"
	serviceCatalogCategoryURL = "/api/v2/service_catalog/categories"
)

// ServiceCatalogService is an interface for interacting with
// the service catalog endpoints of the Freshservice API
type ServiceCatalogService interface {
	List(context.Context, QueryFilter) ([]ServiceCatalogItemDetails, error)
	Categories(context.Context) ([]ServiceCategory, error)
	CreateRequest(context.Context, int, *ServiceRequestOptions) (*ServiceRequestResponse, error)
	Get(context.Context, int) (*ServiceCatalogItemDetails, error)
}

// ServiceCatalogServiceClient facilitates requests with the ServiceCatalogService methods
type ServiceCatalogServiceClient struct {
	client *Client
}

// List all service category items in Freshservice
// Optional filter: category_id=[category_id]
func (sc *ServiceCatalogServiceClient) List(ctx context.Context, filter QueryFilter) ([]ServiceCatalogItemDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   sc.client.Domain,
		Path:   serviceCatalogItemURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ServiceCatalog{}
	if _, err := sc.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.Items, nil
}

// Categories will list all service catalog item categories in freshservice
func (sc *ServiceCatalogServiceClient) Categories(ctx context.Context) ([]ServiceCategory, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   sc.client.Domain,
		Path:   serviceCatalogCategoryURL,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ServiceCategories{}
	if _, err := sc.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// Get a specific service category item from Freshservice via the item's ID
func (sc *ServiceCatalogServiceClient) Get(ctx context.Context, id int) (*ServiceCatalogItemDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   sc.client.Domain,
		Path:   fmt.Sprintf("%s/%d", serviceCatalogItemURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ServiceCatalogItem{}
	if _, err := sc.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// CreateRequest will create a new Freshservice Service Request
func (sc *ServiceCatalogServiceClient) CreateRequest(ctx context.Context, id int, sr *ServiceRequestOptions) (*ServiceRequestResponse, error) {
	res := &ServiceRequestResponse{}

	url := &url.URL{
		Scheme: "https",
		Host:   sc.client.Domain,
		Path:   fmt.Sprintf("%s/%d/place_request", serviceCatalogItemURL, id),
	}

	reqContent, err := json.Marshal(sr)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(reqContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return res, err
	}

	if _, err := sc.client.makeRequest(req, res); err != nil {
		return res, err
	}

	return res, nil
}
