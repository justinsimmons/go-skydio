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

// Expected response from the get marker by id API.
type getMarkerV0Response struct {
	Marker Marker `json:"marker"`
}

// Fetch metadata about a marker by its UUID.
//
// API is in BETA, if there is a breaking change please submit a PR :).
func (s *MarkersService) Get(ctx context.Context, id string) (*Marker, error) {
	u := fmt.Sprintf("/api/v0/marker/%s", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getMarkerV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.Marker, err
}
