// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
)

// newRequest creates an API request.
func (c *Client) newRequest(
	ctx context.Context,
	method string,
	urlStr string,
	body any,
) (*http.Request, error) {

	// Append the API URL to the base URL.
	urlStr = c.baseURL + urlStr

	// If the body is not null we need to marshal to JSON.
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		encoder := json.NewEncoder(buf)
		encoder.SetEscapeHTML(false)

		err := encoder.Encode(body)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to encode HTTP payload '%v' to JSON: %w",
				body,
				err,
			)
		}
	}

	// Construct an HTTP request.
	r, err := http.NewRequestWithContext(ctx, method, urlStr, buf)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to construct HTTP request: %w",
			err,
		)
	}

	if r.Body != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	if c.apiKey != "" {
		r.Header.Set("Authorization", c.apiKey)
	}

	return r, nil
}

func (c *Client) do(
	ctx context.Context,
	r *http.Request,
) (*http.Response, error) {

	if ctx == nil {
		return nil, errors.New("context must not be null")
	}

	resp, err := c.httpClient.Do(r)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return resp, ctx.Err()
		default:
		}

		return nil, fmt.Errorf("failed HTTP request: %w", err)
	}

	return resp, nil
}

// doHTTP sends an API request and returns the API response. The API response
// is JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func doHTTP[T any](ctx context.Context, c *Client, r *http.Request) (*T, error) {
	resp, err := c.do(ctx, r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var apiResp ApiResponse[T]
	err = json.NewDecoder(resp.Body).Decode(&apiResp)

	// ignore EOF errors caused by empty response body
	if err != nil && err != io.EOF {
		return nil, err
	}

	switch apiResp.ErrorCode {
	case ErrorCodeSuccess:
		// API call was a success :)
		return &apiResp.Data, err
	default:
		// API call encountered an error, wrap the response.
		return nil, apiResp.ApiError()
	}
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts any) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// addIdempotencyKey appends an idempotency key to the Skydio API request.
//
// Idempotency Key: unique value that allows certain requests to be idempotent,
// i.e., allows a request to be retried for the same vehicle without
// duplicating the operation. To make a request idempotent, the client must
// specify the value in the Idempotency-Key header of the request.
func addIdempotencyKey(r *http.Request, key uuid.UUID) {
	r.Header.Set("Idempotency-Key", key.String())
}
