package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get vehicle by serial number API.
type getVehicleV0Response struct {
	Vehicle Vehicle `json:"vehicle"`
}

// Fetch metadata about a single vehicle by its serial.
func (s *VehiclesService) Get(
	ctx context.Context,
	serial string,
) (*Vehicle, error) {
	u := fmt.Sprintf("/api/v0/vehicle/%s", serial)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp getVehicleV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Vehicle, err
}
