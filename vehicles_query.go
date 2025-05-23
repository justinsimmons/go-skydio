package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query vehicles API.
type QueryVehiclesV0Response struct {
	Vehicles   []Vehicle  `json:"vehicles"`
	Pagination Pagination `json:"pagination"`
}

type QueryVehiclesOptions struct {
	VehicleSerial string        `url:"vehicle_serial,omitempty"`
	VehicleType   *VehicleClass `url:"vehicle_type,omitempty"`
	VehicleClass  *VehicleClass `url:"vehicle_class,omitempty"`
	UserEmail     string        `url:"user_email,omitempty"`
	PerPage       int           `url:"per_page,omitempty"`
	PageNumber    int           `url:"page_number,omitempty"`
}

// Search vehicles by various parameters.
func (s *VehiclesService) Query(
	ctx context.Context,
	opts *QueryVehiclesOptions,
) (*QueryVehiclesV0Response, error) {

	u, err := addOptions("/api/v0/vehicles", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[QueryVehiclesV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
