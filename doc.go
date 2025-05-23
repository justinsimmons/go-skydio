// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package skydio implements a client for the Skydio Cloud API.
//
// Skydio Cloud allows you to manage your Skydio vehicles using Fleet Manager
// and Media Sync (additional license required) to provide automatic media and
// flight data uploads.
//
// Documentation: https://apidocs.skydio.com
//
// # Usage
//
// A client can be instantiated with or without an API token. It is
// recommended that you use an API token otherwise none of the requests will
// succeed. Instructions on how to obtain an API token can be found on Skydio's
// API documentation page here:
// https://apidocs.skydio.com/reference/authentication#creating-an-api-token
//
//	import (
//		"github.com/justinsimmons/go-skydio"
//	)
//	// Initialize a new client using an API token.
//	client := skydio.NewAuthenticatedClient(
//		context.TODO(),
//		"super secret api token",
//	)
//
// With a client we can then interact with the Skydio Cloud API.
//
//	client := skydio.NewAuthenticatedClient(context.TODO(), token)
//	flights, err := client.Flights.Get(context.TODO(), id)
//	if err != nil {
//		// handle err
//	}
//
// # Rate Limiting.
//
// TODO :/
//
// # Configuration
//
// There are several options that can be specified during the creation of a
// new client. For a complete list see client.go.
//
//	m, err := skydio.NewClient(
//		skydio.WithApiToken(token),
//		skydio.HttpClient(httpClient),
//		skydio.WithURL(sandboxURL),
//	)
//
// # Request Options
//
// Some API methods have optional parameters that can be passed:
//
//	opts := &skydio.QueryVehiclesOptions {UserEmail: "test@gmail.comm"}
//	vehicles, page, err := client.Vehicles.Query(ctx, opts)
//	if err != nil {
//		// handle error
//	}
//
// # Page Based Pagination
//
// To use page based pagination use the [NextPage] and [PerPage] options on
// the request options. The resultant Page struct will have an IsLast method
// that can be used to check if there is remaining data.
//
//	var page int
//	for {
//		flights, page, err := client.Flights.Query(
//			context.TODO(),
//			&skydio.QueryFlightsOptions{
//				PerPage: 50,
//				NextPage: page,
//			},
//		)
//		if err != nil {
//			return err
//		}
//		// Accumulate here the results or check for a specific flight.
//
//		// The `HasNext` helper func checks whether
//		// the API has informed us that there is
//		// more data to retrieve or not.
//		if !clients.HasNext() {
//			break
//		}
//		page++
//	}
package skydio
