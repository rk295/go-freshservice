package freshservice

// Generated Code DO NOT EDIT

const departmentURL = "/api/v2/departments"

// Departments holds a list of Freshservice Department details
type Departments struct {
	List []DepartmentDetails `json:"departments"`
}

// Department holds the details of a specific Freshservice Department
type Department struct {
	Details DepartmentDetails `json:"department"`
}

// Departments is the interface between the HTTP client and the Freshservice department related endpoints
func (fs *Client) Departments() DepartmentsService {
	return &DepartmentsServiceClient{client: fs}
}

// DepartmentsServiceClient facilitates requests with the DepartmentsService methods
type DepartmentsServiceClient struct {
	client *Client
}
