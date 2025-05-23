// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query users API.
type queryUsersV0Response struct {
	User       []User     `json:"users"`
	Pagination Pagination `json:"pagination"`
}

type QueryUsersOptions struct {
	FirstName        string            `url:"first_name,omitempty"`
	LastName         string            `url:"last_name,omitempty"`
	Email            string            `url:"email,omitempty"`
	OrganizationRole *OrganizationRole `url:"organization_role,omitempty"`
	PerPage          int               `url:"per_page,omitempty"`    // Number of results to return per page.
	PageNumber       int               `url:"page_number,omitempty"` // Return a specific page number from results.
}

// Search users by various parameters.
func (s *UsersService) Query(
	ctx context.Context,
	opts *QueryUsersOptions,
) ([]User, *Pagination, error) {

	u, err := addOptions("/api/v0/users", opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := doHTTP[queryUsersV0Response](ctx, s.client, r)
	if err != nil {
		return nil, nil, err
	}

	return resp.User, &resp.Pagination, err
}
