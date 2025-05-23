// Copyright 2025 The go-skydio AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skydio

import (
	"encoding/json"
	"fmt"
)

// Returns the VehicleClass as a raw string.
func (i VehicleClass) String() string {
	return string(i)
}

// IsAVehicleClass returns "true" if the value is listed in the enum definition. "false" otherwise
func (i VehicleClass) IsAVehicleClass() bool {
	switch i {
	case
		VehicleClassSkydioR1,
		VehicleClassSkydio2,
		VehicleClassSkydioX2,
		VehicleClassSkydioX10:
		return true
	}

	return false
}

// Throws an error if the param is not part of the enum.
func VehicleClassString(s string) (VehicleClass, error) {
	vehicleClass := VehicleClass(s)

	if vehicleClass.IsAVehicleClass() {
		return vehicleClass, nil
	}

	return VehicleClass(""), fmt.Errorf(
		"%s does not belong to VehicleClass values",
		s,
	)
}

// MarshalJSON implements the json.Marshaler interface for VehicleClass
func (i VehicleClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for VehicleClass
func (i *VehicleClass) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("VehicleClass should be a string, got %s", data)
	}

	var err error
	*i, err = VehicleClassString(s)
	return err
}
