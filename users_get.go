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

	var resp getUserV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.User, err
}
