package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const departmentURL = "/api/v2/departments"

// DepartmentsService is an interface for interacting with the department
// endpoints of the Freshservice API
type DepartmentsService interface {
	List(context.Context, QueryFilter) ([]DepartmentDetails, string, error)
	Get(context.Context, int, QueryFilter) (*DepartmentDetails, error)
}

// DepartmentServiceClient facilitates requests with the DepartmentService methods
type DepartmentServiceClient struct {
	client *Client
}

// List all Departments
func (d *DepartmentServiceClient) List(ctx context.Context, filter QueryFilter) ([]DepartmentDetails, string, error) {

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

// Get a specific department
func (d *DepartmentServiceClient) Get(ctx context.Context, deptID int, filter QueryFilter) (*DepartmentDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   fmt.Sprintf("%s/%d", departmentURL, deptID),
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Department{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
