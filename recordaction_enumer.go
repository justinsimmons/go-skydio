// Code generated by "enumer -type=RecordAction -transform=snake-upper -trimprefix=RecordAction -json"; DO NOT EDIT.

package skydio

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _RecordActionName = "CONTINUESTOPSTART"

var _RecordActionIndex = [...]uint8{0, 8, 12, 17}

const _RecordActionLowerName = "continuestopstart"

func (i RecordAction) String() string {
	if i >= RecordAction(len(_RecordActionIndex)-1) {
		return fmt.Sprintf("RecordAction(%d)", i)
	}
	return _RecordActionName[_RecordActionIndex[i]:_RecordActionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RecordActionNoOp() {
	var x [1]struct{}
	_ = x[RecordActionContinue-(0)]
	_ = x[RecordActionStop-(1)]
	_ = x[RecordActionStart-(2)]
}

var _RecordActionValues = []RecordAction{RecordActionContinue, RecordActionStop, RecordActionStart}

var _RecordActionNameToValueMap = map[string]RecordAction{
	_RecordActionName[0:8]:        RecordActionContinue,
	_RecordActionLowerName[0:8]:   RecordActionContinue,
	_RecordActionName[8:12]:       RecordActionStop,
	_RecordActionLowerName[8:12]:  RecordActionStop,
	_RecordActionName[12:17]:      RecordActionStart,
	_RecordActionLowerName[12:17]: RecordActionStart,
}

var _RecordActionNames = []string{
	_RecordActionName[0:8],
	_RecordActionName[8:12],
	_RecordActionName[12:17],
}

// RecordActionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RecordActionString(s string) (RecordAction, error) {
	if val, ok := _RecordActionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _RecordActionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RecordAction values", s)
}

// RecordActionValues returns all values of the enum
func RecordActionValues() []RecordAction {
	return _RecordActionValues
}

// RecordActionStrings returns a slice of all String values of the enum
func RecordActionStrings() []string {
	strs := make([]string, len(_RecordActionNames))
	copy(strs, _RecordActionNames)
	return strs
}

// IsARecordAction returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RecordAction) IsARecordAction() bool {
	for _, v := range _RecordActionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for RecordAction
func (i RecordAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordAction
func (i *RecordAction) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RecordAction should be a string, got %s", data)
	}

	var err error
	*i, err = RecordActionString(s)
	return err
}
