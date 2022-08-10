package freshservice

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Locations holds a list of Freshservice Location details
type Locations struct {
	List []LocationDetails `json:"locations"`
}

// Location holds the details of a specific Freshservice Location
type Location struct {
	Details LocationDetails `json:"location"`
}

// LocationDetails are the details related to a specific Location in Freshservice
type LocationDetails struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ParentLocationID int       `json:"parent_location_id"`
	PrimaryContactID int       `json:"primary_contact_id"`
	Address          Address   `json:"address"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Address holds the address details of a specific Location
type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
}

// LocationListOptions holds the available options that can be
// passed when requesting a list of Freshservice Locations
type LocationListOptions struct {
	PageQuery string
	FilterBy  *LocationFilter
}

type LocationFilter struct {
	Name *string
}

// QueryString allows us to pass LocationListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *LocationListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	filterStr := []string{}
	if opts.FilterBy != nil {
		if opts.FilterBy.Name != nil {
			filterStr = append(filterStr, fmt.Sprintf("name:'%s'", *opts.FilterBy.Name))
		}
	}

	filter := fmt.Sprintf("query=%s", url.PathEscape("\""+strings.Join(filterStr, " AND ")+"\""))
	qs = append(qs, filter)

	return strings.Join(qs, "&")
}
