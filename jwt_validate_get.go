package skydio

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Expected response from the get JWT validation API.
type getJwtValidationV0Response struct {
	JWK JWK `json:"jwk"`
}

type GetJwtValidationOptions struct {
	KeyID *uuid.UUID `url:"key_id,omitempty"` // Uniquely identifies a key; corresponds to the 'kid' claim in the Skydio-issued JWT.
}

// Fetch JSON Web Key for a Skydio-issued JWT.
func (s *JwtService) GetJwkValidation(
	ctx context.Context,
	opts *GetJwtValidationOptions,
) (*JWK, error) {
	u, err := addOptions("/api/v0/jwt_validation", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[getJwtValidationV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return &resp.JWK, err
}
