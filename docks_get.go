package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get dock by id API.
type getDockV0Response struct {
	Dock Dock `json:"dock"`
}

// Fetch metadata about a single dock by its serial.
func (s *DocksService) Get(
	ctx context.Context,
	serial string,
) (*Dock, error) {
	u := fmt.Sprintf("/api/v0/dock/%s", serial)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp getDockV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Dock, err
}
