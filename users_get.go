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

// Expected response from the get user by id API.
type getUserV0Response struct {
	User User `json:"User"`
}

// Fetch metadata about a single user by its uuid.
func (s *UsersService) Get(
	ctx context.Context,
	id uuid.UUID,
) (*User, error) {
	u := fmt.Sprintf("/api/v0/user/%s", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getUserV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.User, err
}
