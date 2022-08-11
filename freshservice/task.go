package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TasksService is an interface for interacting with
// the task endpoints of the Freshservice API
type TasksService interface {
	List(context.Context, int) ([]TaskDetails, error)
	Create(context.Context, int, *TaskDetails) (*TaskDetails, error)
	Get(context.Context, int, int) (*TaskDetails, error)
	Update(context.Context, int, int, *TaskDetails) (*TaskDetails, error)
	Delete(context.Context, int, int) error
}

// List all tasks assigned to a given ticket ID
func (c *TasksServiceClient) List(ctx context.Context, tickID int) ([]TaskDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d/tasks", ticketURL, tickID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Tasks{}
	_, err = c.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}

	return res.List, nil
}

// Get a specific task assigned to a given ticket ID
func (c *TasksServiceClient) Get(ctx context.Context, tickID int, tid int) (*TaskDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d/tasks/%d", ticketURL, tickID, tid),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Task{}
	_, err = c.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}
	return &res.Details, nil
}

// Create a task on a given ticket by ID
func (c *TasksServiceClient) Create(ctx context.Context, tickID int, td *TaskDetails) (*TaskDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d/tasks", ticketURL, tickID),
	}

	taskContent, err := json.Marshal(td)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(taskContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Task{}
	_, err = c.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}
	return &res.Details, nil
}

// Update a specific task for a given ticket ID
func (c *TasksServiceClient) Update(ctx context.Context, tickID int, tid int, td *TaskDetails) (*TaskDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d/tasks/%d", ticketURL, tickID, tid),
	}

	taskContent, err := json.Marshal(td)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(taskContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Task{}
	_, err = c.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}
	return &res.Details, nil
}

// Delete a specific task for a given ticket ID
// Note: Deleted tasks are permanently lost. You can't retrieve them once it's get deleted.
func (c *TasksServiceClient) Delete(ctx context.Context, tickID int, tid int) error {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d/tasks/%d", ticketURL, tickID, tid),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	res := &Task{}
	_, err = c.client.makeRequest(req, res)
	if err != nil {
		return err
	}

	return nil
}
