// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
)

// WhoAmIService handles communication with the who am I related
// methods of the Skydio API.
type WhoAmIService service

// Expected response from the who am I API.
type whoAmIV0Response struct {
	ApiToken `json:"api_token"`
}

// Fetch metadata about your API token.
func (s *WhoAmIService) Get(ctx context.Context) (*ApiToken, error) {
	u := "/api/v0/whoami"

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[whoAmIV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.ApiToken, err
}
