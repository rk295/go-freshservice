package freshservice

import (
	"context"
)

// AssetTypesService is an interface for interacting with
// the Asset Type endpoints of the Freshservice API
type AssetTypesService interface {
	List(context.Context, QueryFilter) ([]AssetTypeDetails, string, error)
	Get(context.Context, int) (*AssetTypeDetails, error)
}
