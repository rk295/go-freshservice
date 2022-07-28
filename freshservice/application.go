package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const applicationURL = "/api/v2/applications"

// ApplicationService is an interface for interacting with
// the application endpoints of the Freshservice API
type ApplicationService interface {
	List(context.Context) ([]ApplicationDetails, string, error)
}

// ApplicationServiceClient facilitates requests with the TicketService methods
type ApplicationServiceClient struct {
	client *Client
}

// List all application
func (a *ApplicationServiceClient) List(ctx context.Context) ([]ApplicationDetails, string, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   applicationURL,
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
