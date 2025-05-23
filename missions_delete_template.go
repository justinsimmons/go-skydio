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

// Expected response from the delete mission template API.
type deleteMissionTemplateV0Response struct {
	UUID uuid.UUID `json:"uuid"`
}

// Delete an existing Mission Template by its uuid.
// If the delete is successful the UUID of the deleted mission template is
// returned. Otherwise it returns an error and an empty UUID "0000-....".
func (s *MissionsService) DeleteTemplate(
	ctx context.Context,
	id uuid.UUID,
) (uuid.UUID, error) {
	u := fmt.Sprintf("/api/v0/mission/template/%s/delete", id)

	r, err := s.client.newRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := doHTTP[deleteMissionTemplateV0Response](ctx, s.client, r)
	if err != nil {
		return uuid.Nil, err
	}

	return resp.UUID, err
}
