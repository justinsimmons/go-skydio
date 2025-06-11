package skydio

import "fmt"

// Sentinel errors for the Skydio API.
var (
	ErrSuccess                 = &ApiError{Code: ErrorCodeSuccess}
	ErrRateLimitExceeded       = &ApiError{Code: ErrorCodeRateLimitExceededError}
	ErrPermissions             = &ApiError{Code: ErrorCodePermissionsError}
	ErrFeatureNotEnabled       = &ApiError{Code: ErrorCodeFeatureNotEnabledError}
	ErrInvalidToken            = &ApiError{Code: ErrorCodeInvalidToken}
	ErrNotFound                = &ApiError{Code: ErrorCodeNotFound}
	ErrVehicleNotFound         = &ApiError{Code: ErrorCodeVehicleNotFound}
	ErrDockNotFound            = &ApiError{Code: ErrorCodeDockNotFound}
	ErrAlertConfigNotFound     = &ApiError{Code: ErrorCodeAlertConfigNotFound}
	ErrWebhookNotFound         = &ApiError{Code: ErrorCodeWebhookNotFound}
	ErrFlightNotFound          = &ApiError{Code: ErrorCodeFlightNotFound}
	ErrBatteryNotFound         = &ApiError{Code: ErrorCodeBatteryNotFound}
	ErrAttachmentNotFound      = &ApiError{Code: ErrorCodeAttachmentNotFound}
	ErrMarkerNotFound          = &ApiError{Code: ErrorCodeMarkerNotFound}
	ErrUserNotFound            = &ApiError{Code: ErrorCodeUserNotFound}
	ErrScanNotFound            = &ApiError{Code: ErrorCodeScanNotFound}
	ErrApiTokenNotFound        = &ApiError{Code: ErrorCodeApiTokenNotFound}
	ErrFlightDataFileNotFound  = &ApiError{Code: ErrorCodeFlightDataFileNotFound}
	ErrMissionTemplateNotFound = &ApiError{Code: ErrorCodeMissionTemplateNotFound}
	ErrBadRequest              = &ApiError{Code: ErrorCodeBadRequest}
	ErrBadArgs                 = &ApiError{Code: ErrorCodeBadArgs}
	ErrSchemaValidation        = &ApiError{Code: ErrorCodeSchemaValidationError}
	ErrIntegrity               = &ApiError{Code: ErrorCodeIntegrityError}
	ErrOverwriteProtected      = &ApiError{Code: ErrorCodeOverwriteProtected}
	ErrDeviceOffline           = &ApiError{Code: ErrorCodeDeviceOffline}
	ErrExternalTelemetry       = &ApiError{Code: ErrorCodeExternalTelemetryError}
	ErrStructureNotFound       = &ApiError{Code: ErrorCodeStructureNotFound}
	ErrCloudInMaintenance      = &ApiError{Code: ErrorCodeCloudInMaintenance}
)

type ApiError struct {
	// Skydio error code.
	Code ErrorCode
	// Error message returned from the Skydio API.
	Message string
	// The HTTP status code of the response included in the response body.
	HttpStatusCode int
	// This will be of a type *skydio.ApiResponse[any], due to generics in Go
	// being invariant we cannot use it as a field.
	Response any
}

// Error implements the error interface.
func (e *ApiError) Error() string {
	return fmt.Sprintf("skydio API error %d: %s", e.Code, e.Message)
}

// Is enables matching by sentinel error codes.
func (e *ApiError) Is(target error) bool {
	apiErr, ok := target.(*ApiError)
	if !ok {
		return false
	}

	return e.Code == apiErr.Code
}
