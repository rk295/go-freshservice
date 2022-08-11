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
	Categories(context.Context) ([]ServiceCategory, error)
	CreateRequest(context.Context, int, *ServiceRequestOptions) (*ServiceRequestResponse, error)
	Get(context.Context, int) (*ServiceCatalogDetails, error)
	List(context.Context) ([]ServiceCatalogDetails, string, error)
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
