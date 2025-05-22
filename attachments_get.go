package skydio

import (
	"context"
	"fmt"
	"net/http"
)

// Expected response from the get attachment by id API.
type getAttachmentV0Response struct {
	Attachment Attachment `json:"attachment"`
}

// Fetch metadata about a single attachment by its serial.
func (s *AttachmentsService) Get(
	ctx context.Context,
	serial string,
) (*Attachment, error) {
	u := fmt.Sprintf("/api/v0/attachment/%s", serial)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var resp getAttachmentV0Response
	err = s.client.doHTTP(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Attachment, err
}
