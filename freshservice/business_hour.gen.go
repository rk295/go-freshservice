package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const businessHourURL = "/api/v2/business_hours"

// BusinessHours holds a list of Freshservice BusinessHour details
type BusinessHours struct {
	List []BusinessHourDetails `json:"business_hours"`
}

// BusinessHour holds the details of a specific Freshservice BusinessHour
type BusinessHour struct {
	Details BusinessHourDetails `json:"business_hours"`
}

// BusinessHours is the interface between the HTTP client and the Freshservice businessHour related endpoints
func (fs *Client) BusinessHours() BusinessHoursService {
	return &BusinessHoursServiceClient{client: fs}
}

// BusinessHoursServiceClient facilitates requests with the BusinessHoursService methods
type BusinessHoursServiceClient struct {
	client *Client
}

// List all businessHours
func (d *BusinessHoursServiceClient) List(ctx context.Context) ([]BusinessHourDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   businessHourURL,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &BusinessHours{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific businessHour
func (d *BusinessHoursServiceClient) Get(ctx context.Context, id int) (*BusinessHourDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   path.Join(businessHourURL, fmt.Sprintf("%d", id)),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &BusinessHour{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
