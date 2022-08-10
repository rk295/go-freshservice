package freshservice

import (
	"fmt"
	"net/url"
	"strings"
)

// Departments holds a list of Freshservice department details
type Departments struct {
	List []DepartmentDetails `json:"departments"`
}

// Department holds the details of a specific Freshservice department
type Department struct {
	Details DepartmentDetails `json:"department"`
}

// DepartmentDetails are the details related to a specific department in Freshservice
type DepartmentDetails struct {
	ID           int                           `json:"id"`
	Name         string                        `json:"name"`
	Description  string                        `json:"description"`
	HeadUserID   int                           `json:"head_user_id"`
	PrimeUserID  int                           `json:"prime_user_id"`
	Domains      []string                      `json:"domains"`
	CustomFields DepartmentDetailsCustomFields `json:"custom_fields"`
}

type DepartmentDetailsCustomFields struct {
	Location string `json:"location"`
}

// DepartmentListOptions holds the available options that can be
// passed when requesting a list of Freshservice departments
type DepartmentListOptions struct {
	PageQuery string
	FilterBy  *DepartmentFilter
}

type DepartmentFilter struct {
	Name *string
}

// QueryString allows us to pass DepartmentListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *DepartmentListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	filterStr := []string{}
	if opts.FilterBy.Name != nil {
		filterStr = append(filterStr, fmt.Sprintf("name:'%s'", *opts.FilterBy.Name))
	}

	filter := fmt.Sprintf("query=%s", url.PathEscape("\""+strings.Join(filterStr, " AND ")+"\""))
	qs = append(qs, filter)

	return strings.Join(qs, "&")
}
