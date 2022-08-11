package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DepartmentsService is an interface for interacting with the department
// endpoints of the Freshservice API
type DepartmentsService interface {
	List(context.Context, QueryFilter) ([]DepartmentDetails, string, error)
	Get(context.Context, int, QueryFilter) (*DepartmentDetails, error)
}

// Get a specific department
func (d *DepartmentsServiceClient) Get(ctx context.Context, deptID int, filter QueryFilter) (*DepartmentDetails, error) {

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
