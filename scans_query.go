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

// Expected response from the query scans API.
type queryScansV0Response struct {
	Scans      []Scan     `json:"scans"`
	Pagination Pagination `json:"pagination"`
}

type QueryScansOptions struct {
	FlightID              string    `url:"flight_id,omitempty"`                // Limit results to a specific flight ID. ID can be formatted as a full UUID, or can omit - characters, and is case insensitive.
	Latitude              string    `url:"latitude,omitempty"`                 // Latitude of query point.
	Longitude             string    `url:"longitude,omitempty"`                // Longitude of query point.
	Radius                string    `url:"radius,omitempty"`                   // Radius of query (m).
	UserID                string    `url:"user_id,omitempty"`                  // Limit results to a specific vehicle, by serial number.
	VehicleSerial         string    `url:"vehicle_serial,omitempty"`           // Search for Scans scanned before this timestamp. The timestamp should follow the ISO 8601 standard format.
	ScannedBefore         time.Time `url:"scanned_before,omitempty"`           // Search for Scans scanned before this timestamp. The timestamp should follow the ISO 8601 standard format.
	ScannedSince          time.Time `url:"scanned_since,omitempty"`            // Search for Scans scanned since this timestamp. The timestamp should follow the ISO 8601 standard format.
	PreSignedDownloadURLs bool      `url:"pre_signed_download_urls,omitempty"` // If True, the media download urls will be pre-signed, will not require authentication, and will expire after 1 hour. If False, the urls will redirect to the file and will not expire, but will still require authentication.
	PerPage               int       `url:"per_page,omitempty"`                 // Number of results to return per page.
	PageNumber            int       `url:"page_number,omitempty"`              // Return a specific page number from results.
}

// Search scans by various parameters.
func (s *ScansService) Query(
	ctx context.Context,
	opts *QueryScansOptions,
) ([]Scan, *Pagination, error) {

	u, err := addOptions("/api/v0/scans", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryScansV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.Scans, &resp.Pagination, err
}
