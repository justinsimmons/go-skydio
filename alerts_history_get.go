package skydio

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Optional filters for the get alert history API.
type GetAlertHistoryOptions struct {
	FlightID          *uuid.UUID `url:"flight_id,omitempty"`
	VehicleSerial     string     `url:"vehicle_serial,omitempty"`
	AlertType         string     `url:"alert_type,omitempty"`
	MissionTemplateID string     `url:"mission_template_id,omitempty"`
	AlertsBefore      time.Time  `url:"alerts_before,omitempty"`
	AlertsSince       time.Time  `url:"alerts_since,omitempty"`
	PerPage           int        `url:"per_page,omitempty"`
	PageNumber        int        `url:"page_number,omitempty"`
}

type GetAlertHistoryV0Response struct {
	Alerts     []Alert    `json:"alerts"`
	Pagination Pagination `json:"pagination"`
}

// Search mission alerts by various parameters.
func (s *AlertsService) GetHistory(
	ctx context.Context,
	opts ...GetAlertHistoryOptions,
) (*GetAlertHistoryV0Response, error) {

	u, err := addOptions("/api/v0/alerts_history", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[GetAlertHistoryV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
