package freshservice

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetBusinessHours(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	ctx := context.Background()

	r, err := readTestData("business-hours.json")
	assert.Nil(t, err)

	httpmock.RegisterResponder("GET", "https://domain/api/v2/business_hours",
		httpmock.NewBytesResponder(http.StatusOK, r))

	api, err := New(ctx, "domain", "token", nil)
	assert.Nil(t, err)

	bh, _, err := api.BusinessHours().List(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "Default Business Calendar", bh[0].Description)

}

func TestGetSingleBusinessHours(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	ctx := context.Background()

	r, err := readTestData("business-hours-single.json")
	assert.Nil(t, err)

	httpmock.RegisterResponder("GET", "https://domain/api/v2/business_hours/51000010610",
		httpmock.NewBytesResponder(http.StatusOK, r))

	api, err := New(ctx, "domain", "token", nil)
	assert.Nil(t, err)

	bh, err := api.BusinessHours().Get(ctx, 51000010610)
	assert.Nil(t, err)
	assert.Equal(t, "Default Business Calendar", bh.Description)

}
