package freshservice

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Assets holds a list of Freshservice asset details
type Assets struct {
	List []AssetDetails `json:"assets"`
}

// Asset holds the details of a specific Freshservice asset
type Asset struct {
	Details AssetDetails `json:"asset"`
}

// AssetDetails are the details related to a specific asset in Freshservice
type AssetDetails struct {
	AgentID      int64     `json:"agent_id"`
	AssetTag     string    `json:"asset_tag"`
	AssetTypeID  int       `json:"asset_type_id"`
	AssignedOn   time.Time `json:"assigned_on"`
	AuthorType   string    `json:"author_type"`
	CreatedAt    time.Time `json:"created_at"`
	DepartmentID int64     `json:"department_id"`
	Description  string    `json:"description"`
	DisplayID    int       `json:"display_id"`
	GroupID      int64     `json:"group_id"`
	ID           int64     `json:"id"`
	Impact       string    `json:"impact"`
	LocationID   int64     `json:"location_id"`
	Name         string    `json:"name"`
	UpdatedAt    time.Time `json:"updated_at"`
	UsageType    string    `json:"usage_type"`
	UserID       int64     `json:"user_id"`
	// TODO: Support type fields
	// TypeFields   TypeFields `json:"type_fields"`
}

// AssetListOptions holds the available options that can be
// passed when requesting a list of Freshservice assets
type AssetListOptions struct {
	PageQuery string
	SortBy    *SortOptions
	Embed     *AssetEmbedOptions
	FilterBy  *AssetFilter
}

// AssetEmbedOptions will optonally embed desired metadata in an asset list response
// Each include will consume an additional 2 credits. For example if you embed the stats
// information you will be charged a total of 3 API credits (1 credit for the API call,
// and 2 credits for the additional stats embedding).
type AssetEmbedOptions struct {
	TypeFields bool
	Trashed    bool
}

type AssetFilter struct {
	AssetTypeID  *int64
	DepartmentID *int64
	LocationID   *int64
	AssetState   *string
	UserID       *int64
	AgentID      *int64
	Name         *string
	AssetTag     *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

// QueryString allows us to pass AssetListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *AssetListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	if opts.FilterBy != nil {

		filterStr := []string{}

		if opts.FilterBy.AssetTypeID != nil {
			filterStr = append(filterStr, fmt.Sprintf("asset_type_id:%d", *opts.FilterBy.AssetTypeID))
		}
		if opts.FilterBy.DepartmentID != nil {
			filterStr = append(filterStr, fmt.Sprintf("department_id:%d", *opts.FilterBy.DepartmentID))
		}
		if opts.FilterBy.LocationID != nil {
			filterStr = append(filterStr, fmt.Sprintf("location_id:%d", *opts.FilterBy.LocationID))
		}
		if opts.FilterBy.AssetState != nil {
			filterStr = append(filterStr, fmt.Sprintf("asset_state:'%s'", *opts.FilterBy.AssetState))
		}
		if opts.FilterBy.UserID != nil {
			filterStr = append(filterStr, fmt.Sprintf("user_id:%d", *opts.FilterBy.UserID))
		}
		if opts.FilterBy.AgentID != nil {
			filterStr = append(filterStr, fmt.Sprintf("agent_id:%d", *opts.FilterBy.AgentID))
		}
		if opts.FilterBy.Name != nil {
			filterStr = append(filterStr, fmt.Sprintf("name:'%s'", *opts.FilterBy.Name))
		}
		if opts.FilterBy.AssetTag != nil {
			filterStr = append(filterStr, fmt.Sprintf("asset_tag:'%s'", *opts.FilterBy.AssetTag))
		}
		// TODO: Add CreatedAt and UpdatedAt filters

		filter := fmt.Sprintf("filter=%s", url.PathEscape("\""+strings.Join(filterStr, " AND ")+"\""))
		qs = append(qs, filter)
	}

	// Old below here, leave
	if opts.Embed != nil {
		if opts.Embed.TypeFields {
			qs = append(qs, "include=type_fields")
		}
		if opts.Embed.Trashed {
			qs = append(qs, fmt.Sprintf("trashed=%v", opts.Embed.Trashed))
		}
	}
	return strings.Join(qs, "&")
}
