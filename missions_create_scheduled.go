package skydio

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the create scheduled mission API.
type createScheduledMissionV0Response struct {
	ScheduledMission ScheduledMission `json:"scheduled_mission"`
}

// Properties that can be used to generate a scheduled mission.
type CreateScheduledMissionOptions struct {
	BatteryThreshold int             `json:"battery_threshold,omitempty"` // The miminum battery percentage required for the vehicle to run the mission. If the battery percentage reaches this mid-flight, the vehicle will return to the dock.
	DockID           string          `json:"dock_id"`                     // Deprecated, use dock_serial instead.
	DockSerial       *string         `json:"dock_serial,omitempty"`
	ScheduleEvents   []ScheduleEvent `json:"schedule_events"`
	TemplateUUID     uuid.UUID       `json:"template_uuid"`
	VehicleID        *string         `json:"vehicle_id,omitempty"` // Deprecated, use vehicle_serial instead.
	VehicleSerial    *string         `json:"vehicle_serial,omitempty"`
}

// Create a new schedule for a specific mission template.
func (s *MissionsService) CreateScheduled(
	ctx context.Context,
	opts *CreateScheduledMissionOptions,
) (*ScheduledMission, error) {
	u := "/api/v0/mission/schedule"

	r, err := s.client.newRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[createScheduledMissionV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.ScheduledMission, err
}
