// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Expected response from the query media API.
type queryMediaV0Response struct {
	Files      []File     `json:"files"`
	Pagination Pagination `json:"pagination"`
}

type QueryMediaOptions struct {
	VehicleSerial         *string    `url:"vehicle_serial,omitempty"`           //  Limit results to a specific dock, by serial number.
	FlightID              *uuid.UUID `url:"flight_id,omitempty"`                // Limit results to a specific flight ID. ID can be formatted as a full UUID, or can omit - characters, and is case insensitive.
	Kind                  *FileType  `url:"kind,omitempty"`                     // Filter by media type.
	CapturedBefore        *time.Time `url:"captured_before,omitempty"`          // Search for Media Files captured before this timestamp. The timestamp should follow the ISO 8601 standard format.
	CapturedSince         *time.Time `url:"captured_since,omitempty"`           // Search for Media Files captured since this timestamp. The timestamp should follow the ISO 8601 standard format.
	UploadedBefore        *time.Time `url:"uploaded_before,omitempty"`          // Search for Media Files uploaded before this timestamp. The timestamp should follow the ISO 8601 standard format.
	UploadedSince         *time.Time `url:"uploaded_since,omitempty"`           // Search for Media Files uploaded since this timestamp. The timestamp should follow the ISO 8601 standard format.
	PreSignedDownloadURLs bool       `url:"pre_signed_download_urls,omitempty"` // If True, the media download urls will be pre-signed, will not require authentication, and will expire after 1 hour. If False, the urls will redirect to the file and will not expire, but will still require authentication.
	MissionRunUUID        *uuid.UUID `url:"mission_run_uuid,omitempty"`         // Limit results to a specific mission run ID. Input must be a full UUID.
	MissionTemplateUUID   *uuid.UUID `url:"mission_template_uuid,omitempty"`    // Limit results to a specific mission template ID (UUID).
	MissionWaypointName   *string    `url:"mission_waypoint_name,omitempty"`    // Limit results to a specific mission waypoint name.
	PerPage               int        `url:"per_page,omitempty"`                 // Number of results to return per page.
	PageNumber            int        `url:"page_number,omitempty"`              // Return a specific page number from results.
}

// Search media files by various parameters.
func (s *MediaService) Query(
	ctx context.Context,
	opts *QueryMediaOptions,
) ([]File, *Pagination, error) {

	u, err := addOptions("/api/v0/media_files", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryMediaV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.Files, &resp.Pagination, err
}
