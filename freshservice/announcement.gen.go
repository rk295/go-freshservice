package freshservice

// Generated Code DO NOT EDIT

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
