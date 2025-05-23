// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
	"time"
)

// Expected response from the query markers API.
type queryMarkersV0Response struct {
	Markers    []Marker   `json:"markers"`
	Pagination Pagination `json:"pagination"`
}

type QueryMarkersOptions struct {
	EventTimeBefore *time.Time `url:"event_time_before,omitempty"` // Search for Markers before this timestamp. The timestamp should follow the ISO 8601 standard format.
	EventTimeSince  *time.Time `url:"event_time_since,omitempty"`  // Search for Markers since this timestamp. The timestamp should follow the ISO 8601 standard format.
	PerPage         int        `url:"per_page,omitempty"`          // Number of results to return per page.
	PageNumber      int        `url:"page_number,omitempty"`       // Return a specific page number from results.
}

// Search markers by various parameters.
func (s *MarkersService) Query(
	ctx context.Context,
	opts *QueryMarkersOptions,
) ([]Marker, *Pagination, error) {

	u, err := addOptions("/api/v0/marker", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryMarkersV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.Markers, &resp.Pagination, err
}
