// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

// ScansService handles communication with the scan related
// methods of the Skydio API.
//
// A scan is a set of photos taken autonomously of a single volume, possibly
// over multiple flights. Each scan has a name, a description, a scan time,
// and is associated with flight(s) and corresponding media files. The
// latitude, longitude, and altitude of a scan correspond to the center of
// the scanned volume.
type ScansService service
