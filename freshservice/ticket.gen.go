package freshservice

// Generated Code DO NOT EDIT

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
