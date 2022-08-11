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
	Get(context.Context, int) (*ApplicationDetails, error)
	ListLicenses(context.Context, int) ([]LicensesDetails, error)
	ListUsers(context.Context, int) ([]ApplicationUserDetails, error)
	ListInstallations(context.Context, int) ([]ApplicationInstallationDetails, error)
}

// ListLicenses lists all the licenses for an application
func (a *ApplicationsServiceClient) ListLicenses(ctx context.Context, appID int) ([]LicensesDetails, error) {

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
func (a *ApplicationsServiceClient) ListUsers(ctx context.Context, appID int) ([]ApplicationUserDetails, error) {

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
func (a *ApplicationsServiceClient) ListInstallations(ctx context.Context, appID int) ([]ApplicationInstallationDetails, error) {

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
