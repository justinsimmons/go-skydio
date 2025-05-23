package skydio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the update scheduled mission API.
type updateScheduledMissionV0Response struct {
	ScheduledMission ScheduledMission `json:"scheduled_mission"`
}

// Properties that can be used to update a scheduled mission.
type UpdateScheduledMissionOptions struct {
	MissionID        uuid.UUID       `json:"-"`                           // ID of the scheduled mission to update.
	Active           bool            `json:"active,omitempty"`            // The active status cannot be toggled while editing the other attributes of a scheduled mission.
	BatteryThreshold int             `json:"battery_threshold,omitempty"` // The miminum battery percentage required for the vehicle to run the mission. If the battery percentage reaches this mid-flight, the vehicle will return to the dock.
	DockID           string          `json:"dock_id"`                     // Deprecated, use dock_serial instead.
	DockSerial       *string         `json:"dock_serial,omitempty"`
	ScheduleEvents   []ScheduleEvent `json:"schedule_events"`
	VehicleID        *string         `json:"vehicle_id,omitempty"` // Deprecated, use vehicle_serial instead.
	VehicleSerial    *string         `json:"vehicle_serial,omitempty"`
}

// Update an existing schedule for a mission template.
func (s *MissionsService) UpdateScheduled(
	ctx context.Context,
	opts *UpdateScheduledMissionOptions,
) (*ScheduledMission, error) {
	u := fmt.Sprintf("/api/v0/mission/schedule/%s", opts.MissionID)

	r, err := s.client.newRequest(ctx, http.MethodPatch, u, opts)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[updateScheduledMissionV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.ScheduledMission, err
}
