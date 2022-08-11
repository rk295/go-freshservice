package freshservice

import (
	"strings"
	"time"
)

type AssetTypeDetails struct {
	ID                int64       `json:"id"`
	Name              string      `json:"name"`
	ParentAssetTypeID interface{} `json:"parent_asset_type_id"`
	Description       string      `json:"description"`
	Visible           bool        `json:"visible"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

// AssetTypeListOptions holds the available options that can be
// passed when requesting a list of Freshservice asset types
type AssetTypeListOptions struct {
	PageQuery string
}

// QueryString allows us to pass AssetTypeListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *AssetTypeListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}
	return strings.Join(qs, "&")
}
