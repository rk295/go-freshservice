package freshservice

import "time"

// AssetTypes holds a list of Freshservice asset types
type AssetTypes struct {
	List []AssetTypeDetails `json:"asset_types"`
}

// AssetType holds the details of a specific Freshservice Asset Type
type AssetType struct {
	Details AssetTypeDetails `json:"asset_type"`
}

type AssetTypeDetails struct {
	ID                int64       `json:"id"`
	Name              string      `json:"name"`
	ParentAssetTypeID interface{} `json:"parent_asset_type_id"`
	Description       string      `json:"description"`
	Visible           bool        `json:"visible"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
