package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// BusinessHoursService is an interface for interacting with
// the business hours endpoints of the Freshservice API
type BusinessHoursService interface {
	List(context.Context) ([]BusinessHourDetails, string, error)
	Get(context.Context, int) (*BusinessHourDetails, error)
}

// Get a details for a specific business hour configuration in Freshservice
func (c *BusinessHoursServiceClient) Get(ctx context.Context, id int) (*BusinessHourDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d", businessHourURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &BusinessHour{}
	if _, err := c.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
