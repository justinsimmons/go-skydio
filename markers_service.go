// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

// MarkersService handles communication with the markers related
// methods of the Skydio API.
//
// Markers are typically used to represent incidents like reported suspicious
// activity or other key events during remote flight operations, and will show
// up as an incident in the map view and left panel of Remote Flight Deck for
// DFR Command customers.
type MarkersService service
