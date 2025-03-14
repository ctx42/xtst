package types

import (
	"errors"
)

// TInt is type used in tests.
type TInt struct{ V int }

// NewTInt returns error when v is not 42.
func NewTInt(v int) (*TInt, error) {
	if v != 42 {
		return nil, errors.New("not cool")
	}
	return &TInt{V: v}, nil
}
