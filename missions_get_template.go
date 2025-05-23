package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get mission template by id API.
type getMissionTemplateV0Response struct {
	MissionTemplate MissionTemplate `json:"mission_template"`
}

// Get an existing Mission Template by its uuid.
func (s *MissionsService) GetTemplate(
	ctx context.Context,
	id string,
) (*MissionTemplate, error) {
	u := fmt.Sprintf("/api/v0/mission/template/%s", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getMissionTemplateV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.MissionTemplate, err
}
