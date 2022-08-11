package freshservice

// Generated Code DO NOT EDIT

const requesterURL = "/api/v2/requesters"

// Requesters holds a list of Freshservice Requester details
type Requesters struct {
	List []RequesterDetails `json:"requesters"`
}

// Requester holds the details of a specific Freshservice Requester
type Requester struct {
	Details RequesterDetails `json:"requester"`
}

// Requesters is the interface between the HTTP client and the Freshservice requester related endpoints
func (fs *Client) Requesters() RequestersService {
	return &RequestersServiceClient{client: fs}
}

// RequestersServiceClient facilitates requests with the RequestersService methods
type RequestersServiceClient struct {
	client *Client
}
