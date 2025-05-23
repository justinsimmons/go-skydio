package skydio

import (
	"context"
	"net/http"
)

// Expected response from the query vehicles API.
type QueryAttachmentsV0Response struct {
	Attachments []Attachment `json:"attachments"`
	Pagination  Pagination   `json:"pagination"`
}

type QueryAttachmentsOptions struct {
	AttachmentSerial string          `url:"attachment_serial,omitempty"`
	AttachmentType   *AttachmentType `url:"attachment_type,omitempty"`
	PerPage          int             `url:"per_page,omitempty"`
	PageNumber       int             `url:"page_number,omitempty"`
}

// Search attachments by various parameters.
func (s *AttachmentsService) Query(
	ctx context.Context,
	opts *QueryVehiclesOptions,
) (*QueryAttachmentsV0Response, error) {

	u, err := addOptions("/api/v0/attachments", opts)
	if err != nil {
		return nil, err
	}

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doHTTP[QueryAttachmentsV0Response](ctx, s.client, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
