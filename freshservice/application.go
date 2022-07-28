package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const applicationURL = "/api/v2/applications"

// ApplicationService is an interface for interacting with
// the application endpoints of the Freshservice API
type ApplicationService interface {
	List(context.Context, QueryFilter) ([]ApplicationDetails, string, error)
	Get(context.Context, int64) (*ApplicationDetails, error)
}

// ApplicationServiceClient facilitates requests with the TicketService methods
type ApplicationServiceClient struct {
	client *Client
}

// List all application
// All the below requests are paginated to return only 30 tickets per page.
// Append the parameter "page=[:page_no]" in the url to traverse through pages.
func (a *ApplicationServiceClient) List(ctx context.Context, filter QueryFilter) ([]ApplicationDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   applicationURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Applications{}
	resp, err := a.client.makeRequest(req, res)
	if err != nil {
		fmt.Println("error making request to ", url.String())
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific all application
func (a *ApplicationServiceClient) Get(ctx context.Context, appID int64) (*ApplicationDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Application{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
// QueryString allows us to pass TicketListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *ApplicationListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	return strings.Join(qs, "&")
}
