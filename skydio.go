package skydio

import (
	"net/http"
)

const defaultBaseURL = "https://api.skydio.com"

// Client used to interface with the Skydio API.
type Client struct {
	// HTTP client used to perform API calls to the Skydio API.
	httpClient *http.Client
	// Base URL of the Skydio API.
	baseURL string
	// API key used to authenticate the Skydio API.
	apiKey string

	// Reuse a single struct instead of allocating one for each service on the
	// heap.
	common service

	// Services used to talk to different parts of the Skydio API.

	Alerts         *AlertsService
	Attachments    *AttachmentsService
	Batteries      *BatteriesService
	Docks          *DocksService
	Flights        *FlightsService
	Telemetry      *TelemetryService
	JWT            *JwtService
	Vehicles       *VehiclesService
	MissionResults *MissionResultsService
	Scans          *ScansService
	Users          *UsersService
	WhoAmI         *WhoAmIService
}

type service struct {
	client *Client
}

type option func(*Client)

// WithApiKey overrides the API key on the Client used to Authenticate the
// Skydio API.
func WithApiKey(apiKey string) option {
	return func(c *Client) {
		if apiKey != "" {
			c.apiKey = apiKey
		}
	}
}

// WithHttpClient overrides the default HTTP client used to call the Skydio
// API. This is commonly used to set a default timeout value or to add custom
// request middleware.
func WithHttpClient(httpClient *http.Client) option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

// WithURL overrides the default Skydio API URL. This is used if you have a
// private or sandbox instance of the Skydio API.
func WithURL(url string) option {
	return func(c *Client) {
		if url != "" {
			c.baseURL = url
		}
	}
}

// NewClient creates a new Skydio API client.
// The default client is unauthenticated and therefore unable to make most API
// calls. To authenticate please supply the skydio.WithApiKey("api key here")
// option.
func NewClient(opts ...option) *Client {
	client := &Client{
		httpClient: http.DefaultClient,
		baseURL:    defaultBaseURL,
	}

	client.common.client = client
	client.Alerts = (*AlertsService)(&client.common)
	client.Attachments = (*AttachmentsService)(&client.common)
	client.Batteries = (*BatteriesService)(&client.common)
	client.Docks = (*DocksService)(&client.common)
	client.Flights = (*FlightsService)(&client.common)
	client.Telemetry = (*TelemetryService)(&client.common)
	client.JWT = (*JwtService)(&client.common)
	client.Vehicles = (*VehiclesService)(&client.common)
	client.MissionResults = (*MissionResultsService)(&client.common)
	client.Scans = (*ScansService)(&client.common)
	client.Users = (*UsersService)(&client.common)
	client.WhoAmI = (*WhoAmIService)(&client.common)

	for _, opt := range opts {
		if opt != nil {
			opt(client)
		}
	}

	return client
}

func NewAuthenticatedClient(apiKey string, opts ...option) *Client {
	return NewClient(append(opts, WithApiKey(apiKey))...)
}
