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

// Expected response from the get flight by id API.
type getFlightV0Response struct {
	Flight Flight `json:"flight"`
}

// Fetch metadata about a single flight by its flight id.
func (s *FlightsService) Get(ctx context.Context, id string) (*Flight, error) {
	u := fmt.Sprintf("/api/v0/flight/%s", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getFlightV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.Flight, err
}
