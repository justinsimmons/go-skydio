package skydio

import (
	"context"
	"net/http"
	"time"
)

// Expected response from the query mission templates API.
type QueryMissionTemplatesV0Response struct {
	MissionTemplates []MissionTemplate `json:"mission_templates"`
	Pagination       Pagination        `json:"pagination"`
}

type QueryMissionTemplatesOptions struct {
	UpdatedAfter  *time.Time `url:"updated_after,omitempty"`  // Search for Mission Templates updated after this timestamp. The timestamp should follow the ISO 8601 standard format.
	UpdatedBefore *time.Time `url:"updated_before,omitempty"` // Search for Mission Templates updated before this timestamp. The timestamp should follow the ISO 8601 standard format.
	Name          *string    `url:"name,omitempty"`           // Search for Missions Template by name.
	PerPage       int        `url:"per_page,omitempty"`       // Number of results to return per page
	PageNumber    int        `url:"page_number,omitempty"`    // Return a specific page number from results.
}

// Get mission templates in your organization.
func (s *MissionsService) QueryTemplates(
	ctx context.Context,
	opts *QueryMissionTemplatesOptions,
) (*QueryMissionTemplatesV0Response, error) {

	u, err := addOptions("/api/v0/mission/templates", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[QueryMissionTemplatesV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
