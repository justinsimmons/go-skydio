// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//nolint:lll
package skydio

import (
	"context"

	"github.com/google/uuid"
)

// Skydio Cloud Alerts APIs.
type AlertsServicer interface {
	GetHistory(ctx context.Context, opts *GetAlertHistoryOptions) (*GetAlertHistoryV0Response, error)
}

// Skydio Cloud Attachments APIs.
type AttachmentsServicer interface {
	Get(ctx context.Context, serial string) (*Attachment, error)
	Query(ctx context.Context, opts *QueryVehiclesOptions) ([]Attachment, *Pagination, error)
}

// Skydio Cloud Batteries APIs.
type BatteriesServicer interface {
	Get(ctx context.Context, serial string) (*Battery, error)
	Query(ctx context.Context, opts *QueryBatteriesOptions) ([]Battery, *Pagination, error)
}

// Skydio Cloud Docks APIs.
type DocksServicer interface {
	Get(ctx context.Context, serial string) (*Dock, error)
	Query(ctx context.Context, opts *QueryDocksOptions) ([]Dock, *Pagination, error)
}

// Skydio Cloud Flights APIs.
type FlightsServicer interface {
	Get(ctx context.Context, id string) (*Flight, error)
	Query(ctx context.Context, opts *QueryFlightsOptions) ([]Flight, *Pagination, error)
}

// Skydio Cloud Markers APIs.
type MarkersServicer interface {
	Delete(ctx context.Context, id string) (uuid.UUID, error)
	Get(ctx context.Context, id string) (*Marker, error)
	Query(ctx context.Context, opts *QueryMarkersOptions) ([]Marker, *Pagination, error)
	Upsert(ctx context.Context, marker *Marker) (*Marker, error)
}

// Skydio Cloud Media APIs.
type MediaServicer interface {
	Delete(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	Download(ctx context.Context, id string) (string, error)
	DownloadThumbnail(ctx context.Context, id string) (string, error)
	Get(ctx context.Context, id string) (*File, error)
	Query(ctx context.Context, opts *QueryMediaOptions) ([]File, *Pagination, error)
}

// Skydio Cloud Mission Results APIs.
type MissionResultsServicer interface {
	QueryRuns(ctx context.Context, opts *QueryMissionResultsOptions) ([]MissionRuns, *Pagination, error)
}

// Skydio Cloud Missions APIs.
type MissionsServicer interface {
	CreateScheduled(ctx context.Context, opts *CreateScheduledMissionOptions) (*ScheduledMission, error)
	CreateTemplate(ctx context.Context, opts *CreateMissionTemplateOptions) (*MissionTemplate, error)
	DeleteScheduled(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	DeleteTemplate(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	GetTemplate(ctx context.Context, id string) (*MissionTemplate, error)
	QueryScheduled(ctx context.Context, opts *QueryScheduledMissionsOptions) ([]ScheduledMission, *Pagination, error)
	QueryTemplates(ctx context.Context, opts *QueryMissionTemplatesOptions) ([]MissionTemplate, *Pagination, error)
	ScheduleNow(ctx context.Context, opts *ScheduleMissionNowOptions) (uuid.UUID, error)
	UpdateScheduled(ctx context.Context, opts *UpdateScheduledMissionOptions) (*ScheduledMission, error)
	UpdateTemplate(ctx context.Context, opts *UpdateMissionTemplateOptions) (*MissionTemplate, error)
}

// Skydio Cloud Scans APIs.
type ScansServicer interface {
	Get(ctx context.Context, id string) (*Scan, error)
	Query(ctx context.Context, opts *QueryScansOptions) ([]Scan, *Pagination, error)
}

// Skydio Cloud Telemetry APIs.
type TelemetryServicer interface {
	GetV1(ctx context.Context, flightID string) (*GetTelemetryV1Response, error)
}

// Skydio Cloud Users APIs.
type UsersServicer interface {
	Get(ctx context.Context, id uuid.UUID) (*User, error)
	Query(ctx context.Context, opts *QueryUsersOptions) ([]User, *Pagination, error)
}

// Skydio Cloud Vehicles APIs.
type VehiclesServicer interface {
	Get(ctx context.Context, serial string) (*Vehicle, error)
	Query(ctx context.Context, opts *QueryVehiclesOptions) ([]Vehicle, *Pagination, error)
}

// Skydio Cloud JSON Web Token APIs.
type JwtServicer interface {
	GetWebhookValidation(ctx context.Context, opts *GetWebhookValidationOptions) (*JWK, error)
	GetJwkValidation(ctx context.Context, opts *GetJwtValidationOptions) (*JWK, error)
}

// Skydio Cloud Who am I APIs.
type WhoAmIServicer interface {
	Get(ctx context.Context) (*ApiToken, error)
}

// Ensure each of the respective services implements its corresponding
// interface.

var _ AlertsServicer = (*AlertsService)(nil)
var _ AttachmentsServicer = (*AttachmentsService)(nil)
var _ BatteriesServicer = (*BatteriesService)(nil)
var _ DocksServicer = (*DocksService)(nil)
var _ FlightsServicer = (*FlightsService)(nil)
var _ MarkersServicer = (*MarkersService)(nil)
var _ MediaServicer = (*MediaService)(nil)
var _ MissionResultsServicer = (*MissionResultsService)(nil)
var _ MissionsServicer = (*MissionsService)(nil)
var _ ScansServicer = (*ScansService)(nil)
var _ TelemetryServicer = (*TelemetryService)(nil)
var _ UsersServicer = (*UsersService)(nil)
var _ VehiclesServicer = (*VehiclesService)(nil)
var _ JwtServicer = (*JwtService)(nil)
var _ WhoAmIServicer = (*WhoAmIService)(nil)
