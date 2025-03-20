// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package check provides equality toolkit used by assert package.
package check

import (
	"reflect"

	"github.com/ctx42/xtst/pkg/dump"
	"github.com/ctx42/xtst/pkg/notice"
)

// Error checks "err" is not nil. Returns an error if it's nil.
func Error(err error, opts ...Option) error {
	if err != nil {
		return nil // nolint: nilerr
	}
	ops := DefaultOptions().set(opts)
	const mHeader = "expected non-nil error"
	return notice.New(mHeader).Path(ops.Path)
}

// NoError checks "err" is nil. Returns error it's not nil.
func NoError(err error, opts ...Option) error {
	if err == nil {
		return nil
	}
	ops := DefaultOptions().set(opts)
	const mHeader = "expected error to be nil"
	if isNil(err) {
		return notice.New(mHeader).
			Path(ops.Path).
			Want("<nil>").Have("%T", err)
	}
	return notice.New(mHeader).
		Path(ops.Path).
		Want("<nil>").
		Have("%q", err.Error())
}

// Nil checks "have" is nil. Returns nil if it's, otherwise returns an error
// with a message indicating the expected and actual values.
func Nil(have any, opts ...Option) error {
	if isNil(have) {
		return nil
	}
	ops := DefaultOptions().set(opts)
	const mHeader = "expected value to be nil"
	return notice.New(mHeader).Want("<nil>").
		Path(ops.Path).
		Have("%s", dump.New(ops.DumpConfig).DumpAny(have))
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
	ops := DefaultOptions().set(opts)
	const mHeader = "expected non-nil value"
	return notice.New(mHeader).Path(ops.Path)
}
