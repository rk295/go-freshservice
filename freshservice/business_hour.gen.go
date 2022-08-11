package freshservice

// Generated Code DO NOT EDIT

const businessHourURL = "/api/v2/business_hours"

// BusinessHours holds a list of Freshservice BusinessHour details
type BusinessHours struct {
	List []BusinessHourDetails `json:"business_hours"`
}

// BusinessHour holds the details of a specific Freshservice BusinessHour
type BusinessHour struct {
	Details BusinessHourDetails `json:"business_hours"`
}

// BusinessHours is the interface between the HTTP client and the Freshservice businessHour related endpoints
func (fs *Client) BusinessHours() BusinessHoursService {
	return &BusinessHoursServiceClient{client: fs}
}

// BusinessHoursServiceClient facilitates requests with the BusinessHoursService methods
type BusinessHoursServiceClient struct {
	client *Client
}
