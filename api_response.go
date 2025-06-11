// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

// Error codes returned from the Skydio Cloud API.
//
//go:generate enumer -type=ErrorCode -transform=snake-upper -trimprefix=ErrorCode
type ErrorCode int

const (
	ErrorCodeSuccess                 ErrorCode = 0
	ErrorCodeRateLimitExceededError  ErrorCode = 1429
	ErrorCodePermissionsError        ErrorCode = 2000
	ErrorCodeFeatureNotEnabledError  ErrorCode = 2005
	ErrorCodeInvalidToken            ErrorCode = 3300
	ErrorCodeNotFound                ErrorCode = 4100
	ErrorCodeVehicleNotFound         ErrorCode = 4200
	ErrorCodeDockNotFound            ErrorCode = 4250
	ErrorCodeAlertConfigNotFound     ErrorCode = 4345
	ErrorCodeWebhookNotFound         ErrorCode = 4360
	ErrorCodeFlightNotFound          ErrorCode = 4400
	ErrorCodeBatteryNotFound         ErrorCode = 4450
	ErrorCodeAttachmentNotFound      ErrorCode = 4480
	ErrorCodeMarkerNotFound          ErrorCode = 4490
	ErrorCodeUserNotFound            ErrorCode = 4500
	ErrorCodeScanNotFound            ErrorCode = 4551
	ErrorCodeApiTokenNotFound        ErrorCode = 4575
	ErrorCodeFlightDataFileNotFound  ErrorCode = 4700
	ErrorCodeMissionTemplateNotFound ErrorCode = 4996
	ErrorCodeBadRequest              ErrorCode = 5000
	ErrorCodeBadArgs                 ErrorCode = 5100
	ErrorCodeSchemaValidationError   ErrorCode = 5110
	ErrorCodeIntegrityError          ErrorCode = 7000
	ErrorCodeOverwriteProtected      ErrorCode = 7100
	ErrorCodeDeviceOffline           ErrorCode = 7800
	ErrorCodeExternalTelemetryError  ErrorCode = 8000
	ErrorCodeStructureNotFound       ErrorCode = 8200
	ErrorCodeCloudInMaintenance      ErrorCode = 8300
)

// Contains metadata about the request, including a timestamp.
// Not very well documented...
type ApiResponseMetadata struct {
	Time any `json:"time"`
}

// The Skydio Cloud API is an HTTP-based REST API with request and response
// bodies encoded in JSON format.
type ApiResponse[T any] struct {
	Data         T                   `json:"data"`                    // Requested data, schema varies per endpoint.
	Meta         ApiResponseMetadata `json:"meta"`                    // Contains metadata about the request, including a timestamp.
	ErrorCode    ErrorCode           `json:"skydio_error_code"`       // Skydio specific error code, to share if contacting Skydio Support. Will be 0 for all successful requests.
	ErrorMessage string              `json:"error_message,omitempty"` // Error message.
	StatusCode   int                 `json:"status_code"`             // The HTTP status code of the response included in the response body.
}

// ApiError converts the response into an API Error, useful for generating,
// sentinel errors.
func (r *ApiResponse[T]) ApiError() *ApiError {
	return &ApiError{
		Code:     r.ErrorCode,
		Message:  r.ErrorMessage,
		Response: r,
	}
}
