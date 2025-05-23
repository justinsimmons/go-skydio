package skydio

import (
	"context"
	"net/http"
	"time"
)

// Expected response from the query flights API.
type queryFlightsV0Response struct {
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
) ([]Flight, *Pagination, error) {

	u, err := addOptions("/api/v0/flights", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryFlightsV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.Flights, &resp.Pagination, err
}
