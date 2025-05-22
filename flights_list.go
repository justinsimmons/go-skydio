package skydio

import (
	"context"
	"net/http"
	"time"
)

// Expected response from the query flights API.
type QueryFlightsV0Response struct {
	Flights    []Flight   `json:"flights"`
	Pagination Pagination `json:"pagination"`
}

type QueryFlightsOptions struct {
	VehicleSerial string    `url:"vehicle_serial,omitempty"`
	TakeoffBefore time.Time `url:"takeoff_before,omitempty"`
	TakeoffSince  time.Time `url:"takeoff_since,omitempty"`
	PerPage       int       `url:"per_page,omitempty"`
	PageNumber    int       `url:"page_number,omitempty"`
}

// Search flights by various parameters.
func (s *FlightsService) Query(
	ctx context.Context,
	opts *QueryFlightsOptions,
) (*QueryFlightsV0Response, error) {

	u, err := addOptions("/api/v0/flights", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp QueryFlightsV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
