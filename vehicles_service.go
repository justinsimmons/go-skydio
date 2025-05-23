// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

// VehiclesService handles communication with the vehicle related
// methods of the Skydio API.
//
// Every vehicle claimed within your organization can be identified by the
// vehicle WiFi/UAV Name (eg Skydio2P-xxxx). In our API, we refer to this as
// the vehicle_serial. Users who have piloted a vehicle will be listed in the
// user_emails field. Vehicles connected to Skydio Cloud will also return
// information about their current state, including upload, battery, and
// flight status.
type VehiclesService service
