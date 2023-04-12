// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package target

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseDESCRIBE_BYFromString Parse DESCRIBE_BY from string
func ParseDESCRIBE_BYFromString(str string) (DESCRIBE_BY, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := DESCRIBE_BY_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown DESCRIBE_BY: %s", str)
	}

	return DESCRIBE_BY(v), nil
}

// Equal type compare
func (t DESCRIBE_BY) Equal(target DESCRIBE_BY) bool {
	return t == target
}

// IsIn todo
func (t DESCRIBE_BY) IsIn(targets ...DESCRIBE_BY) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t DESCRIBE_BY) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *DESCRIBE_BY) UnmarshalJSON(b []byte) error {
	ins, err := ParseDESCRIBE_BYFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
