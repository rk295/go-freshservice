package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationsService is an interface for interacting with
// the application endpoints of the Freshservice API
type ApplicationsService interface {
	List(context.Context, QueryFilter) ([]ApplicationDetails, string, error)
	Get(context.Context, int64) (*ApplicationDetails, error)
	ListLicenses(context.Context, int64) ([]LicensesDetails, error)
	ListUsers(context.Context, int64) ([]ApplicationUserDetails, error)
	ListInstallations(context.Context, int64) ([]ApplicationInstallationDetails, error)
}

// Get a specific all application
func (a *ApplicationsServiceClient) Get(ctx context.Context, appID int64) (*ApplicationDetails, error) {

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

// ListLicenses lists all the licenses for an application
func (a *ApplicationsServiceClient) ListLicenses(ctx context.Context, appID int64) ([]LicensesDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d/licenses", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Licenses{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// ListUsers lists all the users of an application
func (a *ApplicationsServiceClient) ListUsers(ctx context.Context, appID int64) ([]ApplicationUserDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d/users", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ApplicationUsers{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// ListInstallations lists all the installations of an application
func (a *ApplicationsServiceClient) ListInstallations(ctx context.Context, appID int64) ([]ApplicationInstallationDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d/installations", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ApplicationInstallations{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
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
