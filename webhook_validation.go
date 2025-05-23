// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the get JWK for webhook validation API.
type getWebhookValidationResponse struct {
	JWK JWK `json:"jwk"`
}

type GetWebhookValidationOptions struct {
	KeyID *uuid.UUID `url:"key_id,omitempty"` // Uniquely identifies a key; corresponds to the 'kid' claim in the Skydio-issued JWT.
}

// Fetch JSON Web Key for JWT corresponding to a webhook request.
func (s *JwtService) GetWebhookValidation(
	ctx context.Context,
	opts *GetWebhookValidationOptions,
) (*JWK, error) {
	u, err := addOptions("/api/v0/webhook_validation", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getWebhookValidationResponse](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.JWK, err
}
