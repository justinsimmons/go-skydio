// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
)

// Expected response from the create mission template API.
type createMissionTemplateV0Response struct {
	MissionTemplate MissionTemplate `json:"mission_template"`
}

// Properties that can be used to create a mission template.
type CreateMissionTemplateOptions struct {
	AutoStart             bool
	CustomVideoBitrate    *int
	DockMission           bool
	FailedFlightDirectRTD bool
	Name                  string
	PhotoIntervalSettings *PhotoIntervalSettings
	RecordingMode         *RecordingMode
	ReturnSettings        *ReturnSettings
	Waypoints             []Waypoint
}

// Create a new mission template.
func (s *MissionsService) CreateTemplate(
	ctx context.Context,
	opts *CreateMissionTemplateOptions,
) (*MissionTemplate, error) {
	u := "/api/v0/mission/template"

	r, err := s.client.newRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[createMissionTemplateV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.MissionTemplate, err
}
