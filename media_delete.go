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

// Expected response from the delete media API.
type deleteMediaV0Response struct {
	File File `json:"file"`
}

// Delete a media file by its uuid. Cannot be undone!
// If the delete is successful the UUID of the deleted file is returned.
// Otherwise it returns an error and an empty UUID "0000-....".
func (s *MediaService) Delete(
	ctx context.Context,
	id uuid.UUID,
) (uuid.UUID, error) {
	u := fmt.Sprintf("/api/v0/mission/media/%s/delete", id)

	r, err := s.client.newRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := doHTTP[deleteMediaV0Response](ctx, s.client, r)
	if err != nil {
		return uuid.Nil, err
	}

	uid, err := uuid.Parse(resp.File.UUID)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as UUID: %w", resp.File.UUID, err)
	}

	return uid, err
}
