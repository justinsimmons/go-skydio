package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get scan by id API.
type getScanV0Response struct {
	Scan Scan `json:"scan"`
}

// Fetch metadata about a single scan by its uuid.
func (s *ScansService) Get(
	ctx context.Context,
	id string,
) (*Scan, error) {
	u := fmt.Sprintf("/api/v0/scan/%s", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getScanV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.Scan, err
}
