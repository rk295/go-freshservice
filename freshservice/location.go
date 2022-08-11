package freshservice

import (
	"context"
)

// LocationsService is an interface for interacting with the location
// endpoints of the Freshservice API
type LocationsService interface {
	List(context.Context, QueryFilter) ([]LocationDetails, string, error)
	Get(context.Context, int) (*LocationDetails, error)
}
