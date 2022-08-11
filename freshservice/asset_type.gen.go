package freshservice

// Generated Code DO NOT EDIT

const assetTypeURL = "/api/v2/asset_types"

// AssetTypes holds a list of Freshservice AssetType details
type AssetTypes struct {
	List []AssetTypeDetails `json:"asset_types"`
}

// AssetType holds the details of a specific Freshservice AssetType
type AssetType struct {
	Details AssetTypeDetails `json:"asset_types"`
}

// AssetTypes is the interface between the HTTP client and the Freshservice assetType related endpoints
func (fs *Client) AssetTypes() AssetTypesService {
	return &AssetTypesServiceClient{client: fs}
}

// AssetTypesServiceClient facilitates requests with the AssetTypesService methods
type AssetTypesServiceClient struct {
	client *Client
}
