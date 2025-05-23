// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the delete marker API.
type deleteMarkerV0Response struct {
	UUID uuid.UUID `json:"uuid"`
}

// Delete a marker. This will remove the incident from the map view of Remote
// Flight Deck for DFR Command customers.
// If the delete is successful the UUID of the deleted marker is returned.
// Otherwise it returns an error and an empty UUID "0000-....".
//
// API is in BETA, if there is a breaking change please submit a PR :).
func (s *MarkersService) Delete(
	ctx context.Context,
	id string,
) (uuid.UUID, error) {
	u := fmt.Sprintf("/api/v0/mission/marker/%s/delete", id)

	r, err := s.client.newRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := doHTTP[deleteMarkerV0Response](ctx, s.client, r)
	if err != nil {
		return uuid.Nil, err
	}

	return resp.UUID, err
}
