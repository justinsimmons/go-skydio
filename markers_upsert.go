// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
)

// Expected response from the upsert marker by id API.
type upsertMarkerV0Response struct {
	Marker Marker `json:"marker"`
}

// Create or update a marker. This will show up as an incident in the map view
// of Remote Flight Deck for DFR Command customers.
func (s *MarkersService) Upsert(
	ctx context.Context,
	marker *Marker,
) (*Marker, error) {

	u := "/api/v0/marker"

	r, err := s.client.newRequest(ctx, http.MethodPost, u, marker)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[upsertMarkerV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.Marker, err
}
