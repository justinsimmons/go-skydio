package skydio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the delete scheduled mission API.
type deleteScheduledMissionV0Response struct {
	UUID uuid.UUID `json:"uuid"`
}

// Delete an existing schedule for a mission template.
// If the delete is successful the UUID of the deleted mission is returned.
// Otherwise it returns an error and an empty UUID "0000-....".
func (s *MissionsService) DeleteScheduled(
	ctx context.Context,
	id uuid.UUID,
) (uuid.UUID, error) {
	u := fmt.Sprintf("/api/v0/mission/schedule/%s/delete", id)

	r, err := s.client.newRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := doHTTP[deleteScheduledMissionV0Response](ctx, s.client, r)
	if err != nil {
		return uuid.Nil, err
	}

	return resp.UUID, err
}
