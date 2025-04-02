// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"reflect"

	"github.com/ctx42/testing/pkg/notice"
)

// Nil checks "have" is nil. Returns nil if it's, otherwise returns an error
// with a message indicating the expected and actual values.
func Nil(have any, opts ...Option) error {
	if isNil(have) {
		return nil
	}
	ops := DefaultOptions(opts...)
	const mHeader = "expected value to be nil"
	return notice.New(mHeader).Want("<nil>").
		Trail(ops.Trail).
		Have("%s", ops.Dumper.Any(have))
}

// isNil returns true if "have" is nil.
func isNil(have any) bool {
	if have == nil {
		return true
	}
	val := reflect.ValueOf(have)
	kind := val.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && val.IsNil() {
		return true
	}
	return false
}

// NotNil checks if "have" is not nil. Returns nil if it is not nil,
// otherwise returns an error with a message indicating the expected and
// actual values.
//
// The returned error might be one or more errors joined with [errors.Join].
func NotNil(have any, opts ...Option) error {
	if !isNil(have) {
		return nil
	}
	ops := DefaultOptions(opts...)
	const mHeader = "expected non-nil value"
	return notice.New(mHeader).Trail(ops.Trail)
}
