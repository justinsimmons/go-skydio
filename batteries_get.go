package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get attachment by id API.
type getBatteryV0Response struct {
	Battery Battery `json:"battery"`
}

// Fetch metadata about a single battery by its serial.
func (s *BatteriesService) Get(
	ctx context.Context,
	serial string,
) (*Battery, error) {
	u := fmt.Sprintf("/api/v0/battery/%s", serial)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getBatteryV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.Battery, err
}
