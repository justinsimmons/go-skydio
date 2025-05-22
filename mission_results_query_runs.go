package skydio

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Expected response from the query mission results API.
type QueryMissionRunsV0Response struct {
	MissionRuns []MissionRuns `json:"mission_runs"`
	Pagination  Pagination    `json:"pagination"`
}

type QueryMissionResultsOptions struct {
	FlightID            string     `url:"flight_id,omitempty"`             // Limit results to a specific flight ID. ID can be formatted as a full UUID, or can omit - characters, and is case insensitive.
	MissionTemplateUUID *uuid.UUID `url:"mission_template_uuid,omitempty"` // Limit results to a specific mission template ID (UUID).
	Result              string     `url:"result,omitempty"`                // Limit results to Mission Runs with this result.
	VehicleSerial       string     `url:"vehicle_serial,omitempty"`        // Limit results to a specific vehicle, by serial number.
	DockSerial          string     `url:"dock_serial,omitempty"`           // Limit results to a specific dock, by serial number.
	StartedBefore       time.Time  `url:"started_before,omitempty"`        // Search for Mission Results started before this timestamp. The timestamp should follow the ISO 8601 standard format.
	EndedSince          time.Time  `url:"ended_since,omitempty"`           // Search for Mission Results ended since this timestamp. The timestamp should follow the ISO 8601 standard format.
	PerPage             int        `url:"per_page,omitempty"`              // Number of results to return per page.
	PageNumber          int        `url:"page_number,omitempty"`           // Return a specific page number from results.
}

// Search mission runs (single executions of a mission) by various parameters.
func (s *MissionResultsService) QueryRuns(
	ctx context.Context,
	opts *QueryMissionResultsOptions,
) (*QueryMissionRunsV0Response, error) {

	u, err := addOptions("/api/v0/mission_runs", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp QueryMissionRunsV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
