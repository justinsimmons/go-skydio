package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query batteries API.
type QueryBatteriesV0Response struct {
	Batteries  []Battery  `json:"batteries"`
	Pagination Pagination `json:"pagination"`
}

type QueryBatteriesOptions struct {
	BatterySerial string `url:"battery_serial,omitempty"` // Limit results by battery serial, with % as wildcard.
	VehicleSerial string `url:"vehicle_serial,omitempty"` // Limit results to a specific vehicle, by serial number.
	PerPage       int    `url:"per_page,omitempty"`       // Number of results to return per page.
	PageNumber    int    `url:"page_number,omitempty"`    // Return a specific page number from results.
}

// Search batteries by various parameters.
func (s *BatteriesService) Query(
	ctx context.Context,
	opts *QueryBatteriesOptions,
) (*QueryBatteriesV0Response, error) {

	u, err := addOptions("/api/v0/batteries", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp QueryBatteriesV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
