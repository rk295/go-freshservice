package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TicketsService is an interface for interacting with
// the ticket endpoints of the Freshservice API
type TicketsService interface {
	List(context.Context, QueryFilter) ([]TicketDetails, string, error)
	Create(context.Context, *TicketDetails) (*TicketDetails, error)
	CreateWithAttachment() (*Ticket, error)
	Get(context.Context, int) (*TicketDetails, error)
	Update(context.Context, int, *TicketDetails) (*TicketDetails, error)
	Delete(context.Context, int) error
}

// Create a new Freshservice ticket
func (t *TicketsServiceClient) Create(ctx context.Context, td *TicketDetails) (*TicketDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   ticketURL,
	}

	ticketContent, err := json.Marshal(td)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(ticketContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Ticket{}
	if _, err := t.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// CreateWithAttachment creates new Freshservice ticket with attachment
func (t *TicketsServiceClient) CreateWithAttachment() (*Ticket, error) {
	return nil, nil
}

// Update a Freshservice ticket
func (t *TicketsServiceClient) Update(ctx context.Context, id int, details *TicketDetails) (*TicketDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   fmt.Sprintf("%s/%d", ticketURL, id),
	}

	ticketContent, err := json.Marshal(details)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(ticketContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Ticket{}
	if _, err := t.client.makeRequest(req, res); err != nil {
		return nil, err
	}
	return &res.Details, nil
}

// Delete Freshservice ticket
func (t *TicketsServiceClient) Delete(ctx context.Context, id int) error {
	url := &url.URL{
		Scheme: "https",
		Host:   t.client.Domain,
		Path:   fmt.Sprintf("%s/%d", ticketURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	if _, err := t.client.makeRequest(req, nil); err != nil {
		return err
	}

	return nil
}
