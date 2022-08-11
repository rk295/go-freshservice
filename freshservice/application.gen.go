package freshservice

// Generated Code DO NOT EDIT

const applicationURL = "/api/v2/applications"

// Applications holds a list of Freshservice Application details
type Applications struct {
	List []ApplicationDetails `json:"applications"`
}

// Application holds the details of a specific Freshservice Application
type Application struct {
	Details ApplicationDetails `json:"application"`
}

// Applications is the interface between the HTTP client and the Freshservice application related endpoints
func (fs *Client) Applications() ApplicationsService {
	return &ApplicationsServiceClient{client: fs}
}

// ApplicationsServiceClient facilitates requests with the ApplicationsService methods
type ApplicationsServiceClient struct {
	client *Client
}
