package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// AnnouncementsService is an interface for interacting with
// the announcement endpoints of the Freshservice API
type AnnouncementsService interface {
	List(context.Context, QueryFilter) ([]AnnouncementDetails, string, error)
	Get(context.Context, int) (*AnnouncementDetails, error)
	Create(context.Context, *AnnouncementDetails) (*AnnouncementDetails, error)
	Update(context.Context, int, *AnnouncementDetails) (*AnnouncementDetails, error)
	Delete(context.Context, int) error
}

// Get a specific Freshservice announcement
func (a *AnnouncementsServiceClient) Get(ctx context.Context, id int) (*AnnouncementDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", announcementURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Announcement{}
	if _, err := a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// Create a new announcement in Freshservice
func (a *AnnouncementsServiceClient) Create(ctx context.Context, details *AnnouncementDetails) (*AnnouncementDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   announcementURL,
	}

	announcementContent, err := json.Marshal(details)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(announcementContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Announcement{}
	if _, err := a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// Update an announcement in Freshservice
func (a *AnnouncementsServiceClient) Update(ctx context.Context, id int, details *AnnouncementDetails) (*AnnouncementDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", announcementURL, id),
	}

	announcementContent, err := json.Marshal(details)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(announcementContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Announcement{}
	if _, err := a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// Delete an announcement in Freshservice
func (a *AnnouncementsServiceClient) Delete(ctx context.Context, id int) error {
	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", announcementURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	if _, err := a.client.makeRequest(req, nil); err != nil {
		return err
	}
	return nil
}
