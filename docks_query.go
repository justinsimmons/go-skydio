package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query docks API.
type QueryDocksV0Response struct {
	Docks      []Dock     `json:"docks"`
	Pagination Pagination `json:"pagination"`
}

type QueryDocksOptions struct {
	DockSerial string    `url:"dock_serial,omitempty"` // Limit results to a specific dock, by serial number.
	DockType   *DockType `url:"dock_type,omitempty"`
	PerPage    int       `url:"per_page,omitempty"`    // Number of results to return per page.
	PageNumber int       `url:"page_number,omitempty"` // Return a specific page number from results.
}

// Search docks by various parameters.
func (s *DocksService) Query(
	ctx context.Context,
	opts *QueryDocksOptions,
) (*QueryDocksV0Response, error) {

	u, err := addOptions("/api/v0/docks", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[QueryDocksV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
