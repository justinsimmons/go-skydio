// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query batteries API.
type queryBatteriesV0Response struct {
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
) ([]Battery, *Pagination, error) {

	u, err := addOptions("/api/v0/batteries", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryBatteriesV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.Batteries, &resp.Pagination, err
}
