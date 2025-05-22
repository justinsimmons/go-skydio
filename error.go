package skydio

// The Skydio Cloud API is an HTTP-based REST API with request and response
// bodies encoded in JSON format.
type ApiResponse[T any] struct {
	Data       T      `json:"data"`              // Requested data, schema varies per endpoint.
	Meta       string `json:"meta"`              // Contains metadata about the request, including a timestamp.
	ErrorCode  int    `json:"skydio_error_code"` // Skydio specific error code, to share if contacting Skydio Support. Will be 0 for all successful requests.
	StatusCode int    `json:"status_code"`       // The HTTP status code of the response included in the response body.
}
