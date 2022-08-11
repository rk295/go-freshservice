package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

var (
	serviceCatalogItemURL     = path.Join(serviceCatalogURL, "items")
	serviceCatalogCategoryURL = path.Join(serviceCatalogURL, "categories")
)

// ServiceCatalogsService is an interface for interacting with
// the service catalog endpoints of the Freshservice API
type ServiceCatalogsService interface {
	List(context.Context) ([]ServiceCatalogDetails, string, error)
	Categories(context.Context) ([]ServiceCategory, error)
	CreateRequest(context.Context, int, *ServiceRequestOptions) (*ServiceRequestResponse, error)
	Get(context.Context, int) (*ServiceCatalogDetails, error)
}

// Categories will list all service catalog item categories in freshservice
func (sc *ServiceCatalogsServiceClient) Categories(ctx context.Context) ([]ServiceCategory, error) {
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
func (sc *ServiceCatalogsServiceClient) Get(ctx context.Context, id int) (*ServiceCatalogDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   sc.client.Domain,
		Path:   fmt.Sprintf("%s/%d", serviceCatalogItemURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ServiceCatalog{}
	if _, err := sc.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// CreateRequest will create a new Freshservice Service Request
func (sc *ServiceCatalogsServiceClient) CreateRequest(ctx context.Context, id int, sr *ServiceRequestOptions) (*ServiceRequestResponse, error) {
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
