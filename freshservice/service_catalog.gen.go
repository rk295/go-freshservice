package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"net/http"
	"net/url"
)

const serviceCatalogURL = "/api/v2/service_catalog"

// ServiceCatalogs holds a list of Freshservice ServiceCatalog details
type ServiceCatalogs struct {
	List []ServiceCatalogDetails `json:"service_items"`
}

// ServiceCatalog holds the details of a specific Freshservice ServiceCatalog
type ServiceCatalog struct {
	Details ServiceCatalogDetails `json:"service_items"`
}

// ServiceCatalogs is the interface between the HTTP client and the Freshservice serviceCatalog related endpoints
func (fs *Client) ServiceCatalogs() ServiceCatalogsService {
	return &ServiceCatalogsServiceClient{client: fs}
}

// ServiceCatalogsServiceClient facilitates requests with the ServiceCatalogsService methods
type ServiceCatalogsServiceClient struct {
	client *Client
}

// List all serviceCatalogs
func (d *ServiceCatalogsServiceClient) List(ctx context.Context) ([]ServiceCatalogDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   serviceCatalogURL,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &ServiceCatalogs{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
