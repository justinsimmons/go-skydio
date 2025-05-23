package skydio

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the query scheduled missions API.
type QueryScheduledMissionsV0Response struct {
	ScheduledMissions []ScheduledMission `json:"scheduled_missions"`
	Pagination        Pagination         `json:"pagination"`
}

type QueryScheduledMissionsOptions struct {
	Active        bool       `url:"attachment_serial,omitempty"` // The active status cannot be toggled while editing the other attributes of a scheduled mission.
	TemplateUUID  *uuid.UUID `url:"template_uuid,omitempty"`
	DockID        *string    `url:"dock_id,omitempty"` // Deprecated, use dock_serial instead.
	DockSerial    *string    `url:"dock_serial,omitempty"`
	VehicleID     *string    `url:"vehicle_id,omitempty"` // Deprecated, use vehicle_serial instead.
	VehicleSerial *string    `url:"vehicle_serial,omitempty"`
	PerPage       int        `url:"per_page,omitempty"`    // Number of results to return per page
	PageNumber    int        `url:"page_number,omitempty"` // Return a specific page number from results.
}

// Search scheduled missions by various parameters.
func (s *MissionsService) QueryScheduled(
	ctx context.Context,
	opts *QueryScheduledMissionsOptions,
) (*QueryScheduledMissionsV0Response, error) {

	u, err := addOptions("/api/v0/mission/schedules", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[QueryScheduledMissionsV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
