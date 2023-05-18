// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package policy

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseCREATE_TYPEFromString Parse CREATE_TYPE from string
func ParseCREATE_TYPEFromString(str string) (CREATE_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := CREATE_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown CREATE_TYPE: %s", str)
	}

	return CREATE_TYPE(v), nil
}

// Equal type compare
func (t CREATE_TYPE) Equal(target CREATE_TYPE) bool {
	return t == target
}

// IsIn todo
func (t CREATE_TYPE) IsIn(targets ...CREATE_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t CREATE_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *CREATE_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseCREATE_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParsePOLICY_TYPEFromString Parse POLICY_TYPE from string
func ParsePOLICY_TYPEFromString(str string) (POLICY_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := POLICY_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown POLICY_TYPE: %s", str)
	}

	return POLICY_TYPE(v), nil
}

// Equal type compare
func (t POLICY_TYPE) Equal(target POLICY_TYPE) bool {
	return t == target
}

// IsIn todo
func (t POLICY_TYPE) IsIn(targets ...POLICY_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t POLICY_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *POLICY_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParsePOLICY_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseACTIVEFromString Parse ACTIVE from string
func ParseACTIVEFromString(str string) (ACTIVE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ACTIVE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ACTIVE: %s", str)
	}

	return ACTIVE(v), nil
}

// Equal type compare
func (t ACTIVE) Equal(target ACTIVE) bool {
	return t == target
}

// IsIn todo
func (t ACTIVE) IsIn(targets ...ACTIVE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ACTIVE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ACTIVE) UnmarshalJSON(b []byte) error {
	ins, err := ParseACTIVEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseLABEL_TYPEFromString Parse LABEL_TYPE from string
func ParseLABEL_TYPEFromString(str string) (LABEL_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := LABEL_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown LABEL_TYPE: %s", str)
	}

	return LABEL_TYPE(v), nil
}

// Equal type compare
func (t LABEL_TYPE) Equal(target LABEL_TYPE) bool {
	return t == target
}

// IsIn todo
func (t LABEL_TYPE) IsIn(targets ...LABEL_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t LABEL_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *LABEL_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseLABEL_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
