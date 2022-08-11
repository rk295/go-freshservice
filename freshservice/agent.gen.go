package freshservice

// Generated Code DO NOT EDIT

import (
	"context"
	"net/http"
	"net/url"
)

const agentURL = "/api/v2/agents"

// Agents holds a list of Freshservice Agent details
type Agents struct {
	List []AgentDetails `json:"agents"`
}

// Agent holds the details of a specific Freshservice Agent
type Agent struct {
	Details AgentDetails `json:"agent"`
}

// Agents is the interface between the HTTP client and the Freshservice agent related endpoints
func (fs *Client) Agents() AgentsService {
	return &AgentsServiceClient{client: fs}
}

// AgentsServiceClient facilitates requests with the AgentsService methods
type AgentsServiceClient struct {
	client *Client
}

// List all agents
func (d *AgentsServiceClient) List(ctx context.Context, filter QueryFilter) ([]AgentDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   agentURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Agents{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
