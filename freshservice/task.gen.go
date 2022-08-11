package freshservice

// Generated Code DO NOT EDIT

const taskURL = "/api/v2/tasks"

// Tasks holds a list of Freshservice Task details
type Tasks struct {
	List []TaskDetails `json:"tasks"`
}

// Task holds the details of a specific Freshservice Task
type Task struct {
	Details TaskDetails `json:"task"`
}

// Tasks is the interface between the HTTP client and the Freshservice task related endpoints
func (fs *Client) Tasks() TasksService {
	return &TasksServiceClient{client: fs}
}

// TasksServiceClient facilitates requests with the TasksService methods
type TasksServiceClient struct {
	client *Client
}
