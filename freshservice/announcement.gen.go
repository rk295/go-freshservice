package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const announcementURL = "/api/v2/announcements"

// Announcements holds a list of Freshservice Announcement details
type Announcements struct {
	List []AnnouncementDetails `json:"announcements"`
}

// Announcement holds the details of a specific Freshservice Announcement
type Announcement struct {
	Details AnnouncementDetails `json:"announcement"`
}

// Announcements is the interface between the HTTP client and the Freshservice announcement related endpoints
func (fs *Client) Announcements() AnnouncementsService {
	return &AnnouncementsServiceClient{client: fs}
}

// AnnouncementsServiceClient facilitates requests with the AnnouncementsService methods
type AnnouncementsServiceClient struct {
	client *Client
}

// List all announcements
func (d *AnnouncementsServiceClient) List(ctx context.Context, filter QueryFilter) ([]AnnouncementDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   announcementURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Announcements{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific announcement
func (d *AnnouncementsServiceClient) Get(ctx context.Context, id int) (*AnnouncementDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   path.Join(announcementURL, fmt.Sprintf("%d", id)),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Announcement{}
	if _, err = d.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
