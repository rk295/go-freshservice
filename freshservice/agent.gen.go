package freshservice

// Generated Code DO NOT EDIT

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
