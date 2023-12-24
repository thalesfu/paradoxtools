package pserialize

import "fmt"

type PBool bool

func (b *PBool) UnmarshalP(data []byte) error {
	s, ok := unquoteBytes(data)
	if !ok {
		return fmt.Errorf("pserialize: invalid bool %q", data)
	}
	if string(s) == "yes" {
		*b = true
	} else {
		*b = false
	}
	return nil
}
