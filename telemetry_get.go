// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

	resp, err := doHTTP[GetTelemetryV1Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
