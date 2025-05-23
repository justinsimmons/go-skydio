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

// Expected response from the get dock by id API.
type getDockV0Response struct {
	Dock Dock `json:"dock"`
}

// Fetch metadata about a single dock by its serial.
func (s *DocksService) Get(
	ctx context.Context,
	serial string,
) (*Dock, error) {
	u := fmt.Sprintf("/api/v0/dock/%s", serial)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getDockV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.Dock, err
}
