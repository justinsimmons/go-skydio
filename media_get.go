// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get media by id API.
type getMediaV0Response struct {
	File File `json:"file"`
}

// Fetch metadata about a single media file by its uuid.
func (s *MediaService) Get(ctx context.Context, id string) (*File, error) {
	u := fmt.Sprintf("/api/v0/media/%s", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getMediaV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.File, err
}
