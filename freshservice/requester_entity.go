package freshservice

import (
	"fmt"
	"net/url"
	"strings"
)

// Requesters holds a list of Freshservice task details
type Requesters struct {
	List []RequesterDetails `json:"requesters"`
}

// Requester holds the details of a specific Freshservice requster
type Requester struct {
	Details RequesterDetails `json:"requester"`
}

// RequesterDetails are the details related to a specific task in Freshservice
type RequesterDetails struct {
	ID                                        int                   `json:"id"`
	IsAgent                                   bool                  `json:"is_agent"`
	FirstName                                 string                `json:"first_name"`
	LastName                                  string                `json:"last_name"`
	JobTitle                                  string                `json:"job_title"`
	PrimaryEmail                              string                `json:"primary_email"`
	SecondaryEmails                           []string              `json:"secondary_emails"`
	WorkPhoneNumber                           string                `json:"work_phone_number"`
	MobilePhoneNumber                         string                `json:"mobile_phone_number"`
	DepartmentIds                             []int                 `json:"department_ids"`
	CanSeeAllTicketsFromAssociatedDepartments bool                  `json:"can_see_all_tickets_from_associated_departments"`
	ReportingManagerID                        int                   `json:"reporting_manager_id"`
	Address                                   string                `json:"address"`
	TimeZone                                  string                `json:"time_zone"`
	TimeFormat                                string                `json:"time_format"`
	Language                                  string                `json:"language"`
	LocationID                                int                   `json:"location_id"`
	BackgroundInformation                     string                `json:"background_information"`
	CustomFields                              RequesterCustomFields `json:"custom_fields"`
	Active                                    bool                  `json:"active"`
	HasLoggedIn                               bool                  `json:"has_logged_in"`
}

type RequesterCustomFields struct {
	QuidditchRole       string `json:"quidditch_role"`
	HogsmeadePermission bool   `json:"hogsmeade_permission"`
}

// RequesterListOptions holds the available options that can be
// passed when requesting a list of Freshservice requesters
type RequesterListOptions struct {
	PageQuery string
	FilterBy  *RequesterFilter
}

// RequesterFilter are optional filters that can be enabled when querying a requester list
type RequesterFilter struct {
	Email             *string
	FirstName         *string
	LastName          *string
	MobilePhoneNumber *string
	WorkPhoneNumber   *string
}

// QueryString allows us to pass RequesterListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *RequesterListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	if opts.FilterBy != nil {
		switch {
		case opts.FilterBy.Email != nil:
			qs = append(qs, fmt.Sprintf("email=%s", *opts.FilterBy.Email))
		case opts.FilterBy.MobilePhoneNumber != nil:
			qs = append(qs, fmt.Sprintf("mobile_phone_number=%s", *opts.FilterBy.MobilePhoneNumber))
		case opts.FilterBy.WorkPhoneNumber != nil:
			qs = append(qs, fmt.Sprintf("work_phone_number=%s", *opts.FilterBy.WorkPhoneNumber))
		}

		filterStr := []string{}

		if opts.FilterBy.LastName != nil {
			filterStr = append(filterStr, fmt.Sprintf("last_name:'%s'", *opts.FilterBy.LastName))
		}
		if opts.FilterBy.FirstName != nil {
			filterStr = append(filterStr, fmt.Sprintf("first_name:'%s'", *opts.FilterBy.FirstName))
		}

		filter := fmt.Sprintf("query=%s", url.PathEscape("\""+strings.Join(filterStr, " AND ")+"\""))
		qs = append(qs, filter)
	}
	return strings.Join(qs, "&")
}
