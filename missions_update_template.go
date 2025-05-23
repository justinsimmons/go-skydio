// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the update mission template API.
type updateMissionTemplateV0Response struct {
	MissionTemplate MissionTemplate `json:"mission_template"`
}

// Properties that can be used to update a mission template.
type UpdateMissionTemplateOptions struct {
	MissionTemplateID     uuid.UUID              `json:"-"` // ID of the scheduled mission to update.
	CustomVideoBitrate    *float32               `json:"custom_video_bitrate"`
	DockMission           bool                   `json:"dock_mission"`
	FailedFlightDirectRtd bool                   `json:"failed_flight_direct_rtd"`
	Name                  string                 `json:"name"`
	PhotoIntervalSettings *PhotoIntervalSettings `json:"photo_interval_settings"`
	RecordingMode         *RecordingMode         `json:"recording_mode"`
	ReturnSettings        *ReturnSettings        `json:"return_settings"`
	Waypoints             []Waypoint             `json:"waypoints"`
}

// Update an existing mission template by its uuid.
func (s *MissionsService) UpdateTemplate(
	ctx context.Context,
	opts *UpdateMissionTemplateOptions,
) (*MissionTemplate, error) {
	u := fmt.Sprintf("/api/v0/mission/template/%s", opts.MissionTemplateID)

	r, err := s.client.newRequest(ctx, http.MethodPatch, u, opts)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[updateMissionTemplateV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.MissionTemplate, err
}
