package freshservice

import (
	"context"
)

// BusinessHoursService is an interface for interacting with
// the business hours endpoints of the Freshservice API
type BusinessHoursService interface {
	List(context.Context) ([]BusinessHourDetails, string, error)
	Get(context.Context, int) (*BusinessHourDetails, error)
}
