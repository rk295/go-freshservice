package freshservice

import (
	"context"
)

// AssetsService is an interface for interacting with
// the asset endpoints of the Freshservice API
type AssetsService interface {
	List(context.Context, QueryFilter) ([]AssetDetails, string, error)
	Get(context.Context, int) (*AssetDetails, error)
}
