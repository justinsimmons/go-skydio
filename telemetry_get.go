package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get telemetry by flight id API.
type GetTelemetryV1Response struct {
	Flight          Flight          `json:"flight"`
	FlightTelemetry FlightTelemetry `json:"flight_telemetry"`
}

// Fetch telemetry from a single flight by its flight id.
func (s *TelemetryService) GetV1(
	ctx context.Context,
	flightID string,
) (*GetTelemetryV1Response, error) {
	u := fmt.Sprintf(
		"/api/v1/flight/%s/telemetry",
		flightID,
	)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp GetTelemetryV1Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
