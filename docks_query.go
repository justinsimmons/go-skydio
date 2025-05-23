package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query docks API.
type queryDocksV0Response struct {
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
) ([]Dock, *Pagination, error) {

	u, err := addOptions("/api/v0/docks", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryDocksV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.Docks, &resp.Pagination, err
}
