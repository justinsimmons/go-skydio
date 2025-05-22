package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query users API.
type QueryUsersV0Response struct {
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
) (*QueryUsersV0Response, error) {

	u, err := addOptions("/api/v0/users", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp QueryUsersV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
