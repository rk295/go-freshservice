package freshservice

// Generated Code DO NOT EDIT

const serviceCatalogURL = "/api/v2/service_catalogs"

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
