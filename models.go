// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"time"

	"github.com/google/uuid"
)

//go:generate enumer -type=AttachmentType -transform=snake-upper -trimprefix=AttachmentType -json
type AttachmentType byte

const (
	AttachmentTypeSpotlight AttachmentType = iota
	AttachmentTypeNightsense
	AttachmentTypeSpeakerAndMic
)

// Location of the attachment on the vehicle.
//
//go:generate enumer -type=MountPoint -transform=snake-upper -trimprefix=MountPoint -json
type MountPoint byte

const (
	MountPointTop MountPoint = iota
	MountPointBottom
	MountPointLeft
	MountPointRight
)

type Attachment struct {
	Serial         string `json:"attachment_serial"`
	AttachmentType *AttachmentType
	MountPoint     string  `json:"mount_point"`    // Location of the attachment on the vehicle.
	Name           *string `json:"name"`           //The display name for the attachment. Defaults to attachment serial.
	VehicleSerial  *string `json:"vehicle_serial"` // Serial of the vehicle linked to this attachment.
}

// Skydio flight.
type Flight struct {
	Attachments      []Attachment `json:"attachments"`
	BatterySerial    *string      `json:"battery_serial"`
	FlightID         string       `json:"flight_id"`
	HasTelemetry     bool         `json:"has_telemetry"` // True if telemetry has been uploaded for this flight.
	Landing          time.Time    `json:"landing"`
	Takeoff          time.Time    `json:"takeoff"`
	TakeoffLatitude  *float64     `json:"takeoff_latitude"`  // Populated if telemetry has been uploaded for this flight.
	TakeoffLongitude *float64     `json:"takeoff_longitude"` // Populated if telemetry has been uploaded for this flight.
	UserEmail        *string      `json:"user_email"`
	VehicleSerial    string       `json:"vehicle_serial"`
}

type Telemetry struct {
	BatteryPercentage             float32   `json:"battery_percentage"`                // The percentage of battery remaining expressed as a float from 0-1.
	GpsAltitude                   float32   `json:"gps_altitude"`                      // GPS-based absolute altitude above sea level [meters].
	GpsHorizontalAccuracy         float32   `json:"gps_horizontal_accuracy"`           // GPS position accuracy estimate in XY [meters].
	GpsLatitude                   float32   `json:"gps_latitude"`                      // GPS-based global latitude [degrees].
	GpsLongitude                  float32   `json:"gps_longitude"`                     // GPS-based global longitude [degrees].
	GpsNumberSatellitesUsed       int       `json:"gps_num_satellites_used"`           // Number of satellites used by the GPS receiver.
	GpsSpeedAccuracy              float32   `json:"gps_speed_accuracy"`                // GPS velocity accuracy estimate [meters / second].
	GpsVelocity                   []float32 `json:"gps_velocity"`                      // GPS-based velocity in the world frame (north-east-down) [meters / second].
	GpsVerticalAccuracy           float32   `json:"gps_vertical_accuracy"`             // GPS position accuracy estimate in Z [meters].
	HeightAboveTakeoff            float32   `json:"height_above_takeoff"`              // Height in meters above the takeoff point. Estimated from visual and baro systems.
	HybridAltitude                float32   `json:"hybrid_altitude"`                   // Vision + Baro estimated absolute altitude above sea level [meters].
	HybridCameraOrientationRpyNed []float32 `json:"hybrid_camera_orientation_rpy_ned"` // Vision + GPS based camera orientation in the world frame (roll, pitch, yaw) [degrees].
	SiteCameraOrientationQuatEnu  []float32 `json:"site_camera_orientation_quat_enu"`  // Orientation of the drone camera in the Site frame (quaternion [x, y, z, w]).
	SiteCameraOrientationRpyEnu   []float32 `json:"site_camera_orientation_rpy_enu"`   // Orientation of the drone camera in the Site frame (roll, pitch, yaw) [degrees].
	SiteCameraPositionEnu         []float32 `json:"site_camera_position_enu"`          // Position of the drone camera in the Site frame (east-north-up) [meters].
	SiteOrientationQuatEnu        []float32 `json:"site_orientation_quat_enu"`         // Orientation of the drone in the Site frame (quaternion [x, y, z, w]).
	SiteOrientationRpyEnu         []float32 `json:"site_orientation_rpy_enu"`          // Orientation of the drone in the Site frame (roll, pitch, yaw) [degrees].
	SitePositionEnu               []float32 `json:"site_position_enu"`                 // Position of the drone in the Site frame (east-north-up) [meters].
	Timestamp                     time.Time `json:"timestamp"`                         // Vehicle epoch time [microseconds].
}

type FlightTelemetry struct {
	AlignedTelemetry []Telemetry `json:"aligned_telemetry"`
}

// Pagination information for paginated APIs.
type Pagination struct {
	CurrentPage int `json:"current_page"` // Current page number of results.
	MaxPerPage  int `json:"max_per_page"` // Maximum number of entries per page, could be less for final page.
	TotalPages  int `json:"total_page"`   // Total pages of results.
}

// HasNext returns true if there are more pages to fetch.
func (p *Pagination) HasNext() bool {
	// Current page starts at 1, if it is 0 then it has not been initialized,
	// so we return true assuming this is the first run.
	if p.CurrentPage <= 0 {
		return true
	}

	return p.TotalPages != p.CurrentPage
}

// Skydio API token.
type ApiToken struct {
	Enabled        bool       `json:"enabled"`
	ID             *uuid.UUID `json:"id"`
	Name           string     `json:"name"`
	OrganizationID string     `json:"organization_id"`
}

// Information about the vehicle's battery.
type BatteryStatus struct {
	Charging   bool // Defaults to false.
	Percentage *float32
}

// Type of dock used.
//
//go:generate enumer -type=DockType -transform=snake-upper -trimprefix=DockType -json
type DockType byte

const (
	DockTypeDock    DockType = iota // Dock for S2/2+, X2.
	DockTypeLite                    // Dock Lite for S2/2+.
	DockTypeX10Dock                 // Dock for X10.
)

// Information about the vehicle's linked dock.
type Dock struct {
	Serial *string  `json:"dock_serial"`
	Type   DockType `json:"dock_type"`
}

// Type of flight.
//
//go:generate enumer -type=FlightStatus -transform=snake-upper -trimprefix=FlightStatus -json
type FlightStatus byte

const (
	FlightStatusUnknown FlightStatus = iota
	FlightStatusFlying
	FlightStatusPostFlight
	FlightStatusPrep
	FlightStatusRest
)

//go:generate enumer -type=MissionState -transform=snake-upper -trimprefix=MissionState -json
type MissionState byte

const (
	MissionStateInProgress MissionState = iota
	MissionStatePaused
	MissionStatePostMissionAction
)

type Mission struct {
	MissionName         *string      `json:"mission_name"`
	MissionTemplateUUID uuid.UUID    `json:"mission_template_id"`
	SecondsUntilTakeoff int          `json:"seconds_until_takeoff"`
	State               MissionState `json:"state"`
}

type MissionStatus struct {
	CurrentMission       *Mission `json:"current_mission"`
	NextScheduledMission *Mission `json:"next_scheduled_mission"`
}

//go:generate enumer -type=RemoteStreamState -transform=snake-lower -trimprefix=RemoteStreamState -json
type RemoteStreamState byte

const (
	RemoteStreamStateDisabled RemoteStreamState = iota
	RemoteStreamStatePending
	RemoteStreamStateActive
	RemoteStreamStateBlocked
	RemoteStreamStateDeviceOffline
	RemoteStreamStateDeviceUnsupported
)

// Information about the file uploads on a vehicle.
type UploadStatus struct {
	FilesToUpload int  `json:"files_to_upload"`
	Uploading     bool `json:"uploading"`
}

type VehicleClass string

const (
	VehicleClassSkydioR1  VehicleClass = "Skydio R1"
	VehicleClassSkydio2   VehicleClass = "Skydio R2"
	VehicleClassSkydioX2  VehicleClass = "Skydio X2"
	VehicleClassSkydioX10 VehicleClass = "Skydio X10"
)

// Skydio vehicle type.
//
// deprecated.
//
//go:generate enumer -type=VehicleType -transform=title -trimprefix=VehicleType -json
type VehicleType byte

const (
	VehicleTypeR1 VehicleType = iota
	VehicleTypeR3
	VehicleTypeE1
	VehicleTypeX10
)

type Vehicle struct {
	BatteryStatus     *BatteryStatus     `json:"battery_status"` // Information about the vehicle's battery. Only populated if vehicle is online
	Dock              *Dock              `json:"dock"`           // Information about the vehicle's linked dock. Only populated if vehicle is linked to a dock.
	FlightStatus      *FlightStatus      `json:"flight_status"`
	IsOnline          *bool              `json:"is_online"`            // Whether or not the vehicle has a direct connection to the internet.
	IsOnlineViaMobile *bool              `json:"is_online_via_mobile"` // Whether or not the vehicle is connected to the internet via a connected mobile device.
	MissionStatus     *MissionStatus     `json:"mission_status"`       // Information about the current mission and next scheduled mission on the vehicle.
	Name              *string            `json:"name"`                 // The display name / nickname for the vehicle. Defaults to vehicle serial.
	RemoteStreamState *RemoteStreamState `json:"remote_stream_state"`
	UploadStatus      *UploadStatus      `json:"upload_status"` // Information about the file uploads on the vehicle. Only populated if vehicle is online.
	UserEmails        []string           `json:"user_emails"`   // Users that have flown this vehicle.
	VehicleClass      VehicleClass       `json:"vehicle_class"`
	VehicleSerial     string             `json:"vehicle_serial"` // Serial number of the vehicle.
	VehicleType       *VehicleType       `json:"vehicle_type"`   // Deprecated.
}

// Type of Alert that was triggered.
//
//go:generate enumer -type=AlertType -transform=snake-upper -trimprefix=AlertType -json
type AlertType byte

const (
	AlertTimeHumanDetected AlertType = iota
	AlertTimeMissionIncomplete
	AlertTimeDockError
	AlertTimeScheduledMissionFailedTakeoff
	AlertTimeFlightStatus
	AlertTimeOnlineStatus
	AlertTimeMediaFileAvailable
	AlertTimeMediaAvailableForScan
	AlertTimeTelemetryAvailable
	AlertTimeWaypointProgress
	AlertTimeLiveStreamStatusChanged
)

type Alert struct {
	AlertID            uuid.UUID  `json:"alert_id"`             // Identifier of the alert.
	AlertTime          time.Time  `json:"alert_time"`           // DateTime the Alert was triggered.
	AlertType          AlertType  `json:"alert_type"`           // Type of Alert that was triggered.
	FlightID           *uuid.UUID `json:"flight_id"`            // Flight ID of Flight that generated Alert.
	MissionExecutionID *uuid.UUID `json:"mission_execution_id"` // Execution ID of Mission that generated Alert.
	MissionResult      *string    `json:"mission_result"`       // Outcome of Mission that generated Alert, if alert_type is MISSION_INCOMPLETE.
	MissionTemplateID  *uuid.UUID `json:"mission_template_id"`  // Template ID of Mission that generated Alert.
	VehicleSerial      *string    `json:"vehicle_serial"`       // Serial of Vehicle that generated Alert.
}

// Individual cell within a Battery.
type BatteryCell struct {
	ID         int     `json:"cell_id"`     // Identifies which cell, 0, 1, 2, etc. this is within a Battery.
	MaxVoltage float32 `json:"max_voltage"` // Maximimum voltage reported by a Battery Cell over its lifetime.
	MinVoltage float32 `json:"min_voltage"` // Minimum voltage reported by a Battery Cell over its lifetime.
}

type Battery struct {
	Cells           []BatteryCell `json:"battery_cells"`     // Array of individual cells within a Battery.
	Name            *string       `json:"battery_name"`      // Name of the battery.
	Serial          string        `json:"battery_serial"`    // Serial Number and primary Battery idenfitier.
	Cycles          *int          `json:"cycles"`            // Count of charge/discharge cycles reported by a Battery.
	FlightCount     *int          `json:"flight_count"`      // How many Flights a Battery has been used with.
	MaxCellTemp     *int          `json:"max_cell_temp"`     // Maximum temperature reported by a Battery across all cells.
	MaxVoltage      *float32      `json:"max_voltage"`       // Maximum voltage reported by a Battery across all cells.
	MinCellTemp     *int          `json:"min_cell_temp"`     // Minimum temperature reported by a Battery across all cells.
	MinVoltage      *float32      `json:"min_voltage"`       // Minimum voltage reported by a Battery across all cells.
	TotalFlightTime *int          `json:"total_flight_time"` // Total Flight Time of this battery, in seconds.
}

// JSON Web Key used for JWT validation.
type JWK struct {
	Alg string `json:"alg"` // Algorithm used for the token.
	E   string `json:"e"`   // RSA public exponent.
	KID string `json:"kid"` // Short for 'key id', uniquely identifies a key.
	KTY string `json:"kty"` // Short for 'key type', represents the family of algorithm used for the token.
	N   string `json:"n"`   // RSA modulus.
}

type OrganizationRole string

const (
	OrganizationRoleMember      OrganizationRole = "MEMBER"
	OrganizationRoleTester      OrganizationRole = "TESTER"
	OrganizationRoleModerator   OrganizationRole = "MODERATOR"
	OrganizationRoleAdmin       OrganizationRole = "ADMIN"
	OrganizationRoleRemotePilot OrganizationRole = "REMOTE PILOT"
)

// User in the Skydio system.
type User struct {
	Email            string           `json:"email"`
	FirstName        *string          `json:"first_name"`
	LastName         *string          `json:"last_name"`
	OrganizationRole OrganizationRole `json:"organization_id"`
	ID               string           `json:"user_id"`
}

//go:generate enumer -type=FileType -transform=title-lower -trimprefix=FileType -json
type FileType byte

const (
	FileTypeVehicleImage FileType = iota
	FileTypeVehicleImageDng
	FileTypeVehicleIrImage
	FileTypeVehicleRadiometrucImage
	FileTypeVehicleVideoPreview
	FileTypeVechcleVideoRaw
	FileTypeVehicleTelemetryCsv
)

type File struct {
	CapturedTime        *time.Time `json:"captured_time"`
	DownloadURL         *string    `json:"download_url"` // download_url is included only if the file is uploaded. If the file is not yet uploaded the media endpoints will not return it, however it is returned as part of this response to indicate the file was captured as part of the scan.
	Name                string     `json:"filename"`
	FlightID            string     `json:"flight_id"`
	Kind                FileType   `json:"kind"`
	MissionRunUUID      *string    `json:"mission_run_uuid"`
	MissionTemplateUUID *string    `json:"mission_template_uuid"`
	MissionWaypointName *string    `json:"mission_waypoint_name"` // The name of the waypoint where the media was captured.
	SHA256              *string    `json:"sha256"`
	Size                *int       `json:"size"`
	UploadedTime        *time.Time `json:"uploaded_time"`
	UserID              *string    `json:"user_id"`
	UUID                string     `json:"uuid"`
}

type Scan struct {
	Altitude              *float32   `json:"altitude"` // Altitude of the scan-volume center (m).
	Description           *string    `json:"description"`
	FilePagination        Pagination `json:"file_pagination"`
	Files                 []File     `json:"files"`
	Flights               []Flight   `json:"flights"`
	Latitude              *float64   `json:"latitude"`
	Longitude             *float32   `json:"longitude"`
	Name                  *string    `json:"name"`
	PhotoCount            *float32   `json:"photo_count"`
	ScanTime              *time.Time `json:"scan_time"`
	ID                    uuid.UUID  `json:"uuid"`
	ViewpointMediaID      *string    `json:"viewpoint_media_id"`
	ViewpointThumbnailURL *string    `json:"viewpoint_thumbnail_url"`
}

type MissionRuns struct {
	DockSerial          *string    `json:"dock_serial"`
	EndTime             *time.Time `json:"end_time"`
	FlightID            *string    `json:"flight_id"`
	MissionTemplateUUID *string    `json:"mission_template_id"`
	Result              string     `json:"result"`
	StartTime           *time.Time `json:"start_time"`
	UUID                string     `json:"uuid"`
	VehicleSerial       *string    `json:"vehicle_serial"`
}

type DayOfTheWeek uint8

const (
	Monday DayOfTheWeek = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type ScheduleEvent struct {
	DayOfTheWeek *DayOfTheWeek `json:"day_of_week"` // Numerical representation of the day of the week, where 1 is Monday and 7 is Sunday.
	TimeOfDay    *string       `json:"time_of_day"` // Clock time conforming to the full-time spec in RFC3339.
}

type ScheduledMission struct {
	Active           bool            `json:"active"`            // The active status cannot be toggled while editing the other attributes of a scheduled mission.
	BatteryThreshold int             `json:"battery_threshold"` // The miminum battery percentage required for the vehicle to run the mission. If the battery percentage reaches this mid-flight, the vehicle will return to the dock.
	DockID           *string         `json:"dock_id"`           // Deprecated, use dock_serial instead
	DockSerial       *string         `json:"dock_serial"`
	ScheduleEvents   []ScheduleEvent `json:"schedule_events"`
	TemplateUUID     uuid.UUID       `json:"template_uuid"`
	UUID             uuid.UUID       `json:"uuid"`
	VehicleID        *string         `json:"vehicle_id"` // Deprecated //  Use vehicle_serial instead
	VehicleSerial    *string         `json:"vehicle_serial"`
}

type MissionTemplate struct {
	// TODO: Fill in.
}

// Shutter speed expressed as a fraction of a second. A faster shutter speed
// will reduce blurring of images taken while flying quickly but may also
// decrease image quality.
type ShutterSpeed string

const (
	ShutterSpeedAuto ShutterSpeed = "AUTO"
	ShutterSpeed24   ShutterSpeed = "1/24"
	ShutterSpeed30   ShutterSpeed = "1/30"
	ShutterSpeed48   ShutterSpeed = "1/48"
	ShutterSpeed60   ShutterSpeed = "1/60"
	ShutterSpeed96   ShutterSpeed = "1/96"
	ShutterSpeed120  ShutterSpeed = "1/120"
	ShutterSpeed240  ShutterSpeed = "1/240"
	ShutterSpeed480  ShutterSpeed = "1/480"
	ShutterSpeed960  ShutterSpeed = "1/960"
	ShutterSpeed1920 ShutterSpeed = "1/1920"
	ShutterSpeed3840 ShutterSpeed = "1/3840"
	ShutterSpeed7680 ShutterSpeed = "1/7680"
)

// Photo settings to apply if recording_mode is set to a photo mode.
type PhotoIntervalSettings struct {
	CustomPhotoQuality *int          `json:"custom_photo_quality"` // Determines quality/size of photos. 100 is no compression, 0 is maximum compression.
	CustomShutterSpeed *ShutterSpeed `json:"custom_shutter_speed"` // Shutter speed expressed as a fraction of a second. A faster shutter speed will reduce blurring of images taken while flying quickly but may also decrease image quality. Defaults to AUTO.
	TimeInterval       int           `json:"time_interval"`        // Time between photos, in seconds. 0 to 10 Defaults to 1
}

type RecordingMode string

const (
	RecordingModeVideoDefault        RecordingMode = "VIDEO_DEFAULT"
	RecordingModeVideo1080p30Fps     RecordingMode = "VIDEO_1080P_30FPS"
	RecordingModeVideo1080pFull30Fps RecordingMode = "VIDEO_1080P_FULL_30FPS"
	RecordingModeVideo4k30Fps        RecordingMode = "VIDEO_4K_30FPS"
	RecordingModeVideo4k60FpsHdr     RecordingMode = "VIDEO_4K_60FPS_HDR"
	RecordingModePhotoHdr            RecordingMode = "PHOTO_HDR"
	RecordingModePhotoDefault        RecordingMode = "PHOTO_Default"
)

//go:generate enumer -type=HeightBehavior -transform=snake-upper -trimprefix=HeightBehavior -json
type HeightBehavior byte

const (
	HeightBehaviorAbsolute HeightBehavior = iota // Return height is measured in feet above the Launch Point.
	HeightBehaviorRelative                       // Return height is measured in feet relative to the last waypoint.
)

//go:generate enumer -type=MissionCompletedReturnType -transform=snake-upper -trimprefix=MissionCompletedReturnType -json
type MissionCompletedReturnType byte

const (
	MissionCompletedReturnTypeDefault MissionCompletedReturnType = iota
)

//go:generate enumer -type=MissionInterruptedReturnType -transform=snake-upper -trimprefix=MissionInterruptedReturnType -json
type MissionInterruptedReturnType byte

const (
	MissionInterruptedReturnTypeBacktrack MissionInterruptedReturnType = iota
	MissionInterruptedReturnTypeUpAndOver
)

type ReturnSettings struct {
	HeightBehavior                *HeightBehavior               `json:"height_behavior"`                   // If mission_interrupted_return_type is 'UP_AND_OVER' setting this to 'ABSOLUTE' means the return_height is measured in feet above the Launch Point. Setting this to 'RELATIVE' means the return_height is measured in feet relative to the last waypoint.
	LostConnectionWaitTimeSeconds *int                          `json:"lost_connection_wait_time_seconds"` // Amount of time in seconds that the vehicle should wait, allowing time to reconnect, before it initiates a return flight.
	MissionCompletedReturnType    *MissionCompletedReturnType   `json:"mission_completed_return_type"`     // If included will execute a direct linear return at the end of a successful mission. Requires direct line of sight between last waypoint and dock. If left blank will default to the mission_interrupted_return_type.
	MissionInterruptedReturnType  *MissionInterruptedReturnType `json:"mission_interrupted_return_type"`   // UP_AND_OVER will fly from the last waypoint to the dock at the specified return_height. BACKTRACK returns to dock using the shortest known path
	ReturnHeight                  *float32                      `json:"return_height"`                     // If mission_interrupted_return_type is "UP_AND_OVER", the vehicle will return to the dock at this height.
	Speed                         *int                          `json:"speed"`                             // Speed in m/s. 1 to 14 Defaults to 2.
}

//go:generate enumer -type=RecordAction -transform=snake-upper -trimprefix=RecordAction -json
type RecordAction byte

const (
	RecordActionContinue RecordAction = iota
	RecordActionStop
	RecordActionStart
)

type Orientation struct {
	GimbalPitchDegrees float32 `json:"gimbal_pitch_degrees"`
	HeadingDegrees     float32 `json:"heading_degrees"`
}

//go:generate enumer -type=ReferenceFrame -transform=snake-upper -trimprefix=ReferenceFrame -json
type ReferenceFrame byte

const (
	ReferenceFrameNav ReferenceFrame = iota
	ReferenceFrameGps
)

type Position struct {
	Frame     ReferenceFrame `json:"frame"`     // The reference frame for this waypoint's coordinates. The origin of the NAV frame is the vehicle's location at startup.
	Latitude  *float32       `json:"latitude"`  // Latitude in degrees. This value should only be provided if the frame is 'GPS'.
	Longitude *float32       `json:"longitude"` // Longitude in degrees. This value should only be provided if the frame is 'GPS'.
	X         float32        `json:"x"`         // X-coordinate of the vehicle position in meters. This value should only be provided if the frame is 'NAV'.
	Y         float32        `json:"y"`         // Y-coordinate of the vehicle position in meters. This value should only be provided if the frame is 'NAV'.
	Z         float32        `json:"z"`         // Z-coordinate of the vehicle position in feet if frame or z_frame is 'NAV', or altitude if frame or z_frame is 'GPS'. Since GPS altitude tends to be unreliable, z should almost always be provided in the NAV frame. See z_frame below for specifying this when X/Y coordinates are GPS-based.
	ZFrame    *string        `json:"z_frame"`   // If provided, this will override the frame type for just the z coordinate. z_frame should usually be specified as 'NAV' if frame is 'GPS'.
}

type Waypoint struct {
	Action             *string       `json:"action"`
	Name               *string       `json:"name"`
	Orientation        Orientation   `json:"orientation"`
	Position           Position      `json:"position"`
	RecordAction       *RecordAction `json:"record_action"`
	TransitOrientation *Orientation  `json:"transit_orientation"`
	TransitSpeed       *float32      `json:"transit_speed"`
	WaitTimeSeconds    *int          `json:"wait_time_seconds"`
}

// Additional details about the marker.
type IncidentDetails struct {
	Code       *string `json:"code"`        // The incident type code.
	IncidentID *string `json:"incident_id"` // The identifier of the incident. Ex. case number, incident number, etc.
	Piority    *string `json:"priority"`    // The priority of the incident. Ex. P0, P1, P2.
}

//go:generate enumer -type=MarkerType -transform=snake-upper -trimprefix=MarkerType -json
type MarkerType byte

const (
	MarkerTypeIncident MarkerType = iota
)

// Markers are typically used to represent incidents like reported suspicious
// activity or other key events during remote flight operations, and will show
// up as an incident in the map view and left panel of Remote Flight Deck for
// DFR Command customers.
type Marker struct {
	Area          *string          `json:"area"`           // The area that the marker is associated with. Ex. "5" (precinct number), "Downtown" (sector name), etc.
	Description   string           `json:"description"`    // Detailed description of the marker that will appear in the marker details.
	EventTime     time.Time        `json:"event_time"`     // The time of the event that the marker is associated with.
	Latitude      float32          `json:"latitude"`       // Latitude in degrees.
	Longitude     float32          `json:"longitude"`      // Longitude in degrees.
	MarkerDetails *IncidentDetails `json:"marker_details"` // Additional details about the marker. Fields contained within depend on the type of marker.
	SourceName    *string          `json:"source_name"`    // Partner application name. Ex. Axon, Fusus, etc.
	Title         *string          `json:"title"`          // A short title for the marker that will appear on the list view for the marker.
	Type          MarkerType       `json:"type"`           // The type of the marker. This will affect how the marker is displayed on the map view.
	UUID          *uuid.UUID       `json:"uuid"`           // ID of the marker.
	Version       float32          `json:"version"`        // An increasing number that represents the version of the marker. Once this field is set, future updates must use a higher version number for changes to take effect. Updates with the same or a lower version will be ignored.
}
