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

// Download a thumbnail for a media file by its uuid.
// This returns a JSON string of the thumbnail you query. You will have to
// unmarshal and handle the values yourself.
func (s *MediaService) DownloadThumbnail(
	ctx context.Context,
	id string,
) (string, error) {

	u := fmt.Sprintf("/api/v0/media/download/%s/thumbnail", id)

	r, err := s.client.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", err
	}

	resp, err := doHTTP[string](ctx, s.client, r)
	if err != nil {
		return "", err
	}

	// This should never happen, this check is for added redundancy.
	if resp == nil {
		return "", fmt.Errorf(
			"received an empty JSON string from the Skydio API",
		)
	}

	return *resp, err
}
