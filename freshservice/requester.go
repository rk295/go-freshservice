package freshservice

import (
	"context"
)

// RequestersService is an interface for interacting with
// the requester endpoints of the Freshservice API
type RequestersService interface {
	List(context.Context, QueryFilter) ([]RequesterDetails, string, error)
	Get(context.Context, int) (*RequesterDetails, error)
}
