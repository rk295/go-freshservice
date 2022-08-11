package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const ticketURL = "/api/v2/tickets"

// Tickets holds a list of Freshservice Ticket details
type Tickets struct {
	List []TicketDetails `json:"tickets"`
}

// Ticket holds the details of a specific Freshservice Ticket
type Ticket struct {
	Details TicketDetails `json:"ticket"`
}

// Tickets is the interface between the HTTP client and the Freshservice ticket related endpoints
func (fs *Client) Tickets() TicketsService {
	return &TicketsServiceClient{client: fs}
}

// TicketsServiceClient facilitates requests with the TicketsService methods
type TicketsServiceClient struct {
	client *Client
}

// List all tickets
func (d *TicketsServiceClient) List(ctx context.Context, filter QueryFilter) ([]TicketDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   ticketURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Tickets{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific ticket
func (d *TicketsServiceClient) Get(ctx context.Context, id int) (*TicketDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   path.Join(ticketURL, fmt.Sprintf("%d", id)),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Ticket{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
