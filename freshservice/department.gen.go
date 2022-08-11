package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"net/http"
	"net/url"
)

const departmentURL = "/api/v2/departments"

// Departments holds a list of Freshservice Department details
type Departments struct {
	List []DepartmentDetails `json:"departments"`
}

// Department holds the details of a specific Freshservice Department
type Department struct {
	Details DepartmentDetails `json:"department"`
}

// Departments is the interface between the HTTP client and the Freshservice department related endpoints
func (fs *Client) Departments() DepartmentsService {
	return &DepartmentsServiceClient{client: fs}
}

// DepartmentsServiceClient facilitates requests with the DepartmentsService methods
type DepartmentsServiceClient struct {
	client *Client
}

// List all departments
func (d *DepartmentsServiceClient) List(ctx context.Context, filter QueryFilter) ([]DepartmentDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   departmentURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Departments{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
