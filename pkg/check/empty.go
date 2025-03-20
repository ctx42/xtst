package check

import (
	"reflect"

	"github.com/ctx42/xtst/pkg/notice"
)

// Empty checks if "have" is empty. Returns nil if it's, otherwise it returns
// an error with a message indicating the expected and actual values.
//
// Empty values are:
//   - nil
//   - int(0)
//   - float64(0)
//   - float32(0)
//   - false
//   - len(array) == 0
//   - len(slice) == 0
//   - len(map) == 0
//   - len(chan) == 0
//   - time.Time{}
func Empty(have any, opts ...Option) error {
	if isEmpty(have) {
		return nil
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected argument to be empty").
		Trail(ops.Trail).
		Want("<empty>").
		Have("%#v", have)
}

// isEmpty returns true if "have" is empty.
func isEmpty(have any) bool {
	if isNil(have) {
		return true
	}

	val := reflect.ValueOf(have)
	switch val.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice:
		if val.Len() == 0 {
			return true
		}

	case reflect.Ptr:
		return isEmpty(val.Elem().Interface())

	default:
		zero := reflect.Zero(val.Type())
		if reflect.DeepEqual(have, zero.Interface()) {
			return true
		}
	}

	return false
}

// NotEmpty checks "have" is not empty. Returns nil if it's otherwise, it
// returns an error with a message indicating the expected and actual values.
//
// See [check.Empty] for list of values which are considered empty.
func NotEmpty(have any, opts ...Option) error {
	if !isEmpty(have) {
		return nil
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected non-empty value").Trail(ops.Trail)
}
