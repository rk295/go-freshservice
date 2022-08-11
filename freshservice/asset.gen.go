package freshservice

// Generated Code DO NOT EDIT

const assetURL = "/api/v2/assets"

// Assets holds a list of Freshservice Asset details
type Assets struct {
	List []AssetDetails `json:"assets"`
}

// Asset holds the details of a specific Freshservice Asset
type Asset struct {
	Details AssetDetails `json:"asset"`
}

// Assets is the interface between the HTTP client and the Freshservice asset related endpoints
func (fs *Client) Assets() AssetsService {
	return &AssetsServiceClient{client: fs}
}

// AssetsServiceClient facilitates requests with the AssetsService methods
type AssetsServiceClient struct {
	client *Client
}
