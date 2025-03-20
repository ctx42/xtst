// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package check provides equality toolkit used by assert package.
package check

import (
	"errors"
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

// ErrorIs checks whether any error in "err" tree matches target. Returns nil
// if it's, otherwise returns an error with a message indicating the expected
// and actual values.
func ErrorIs(err, target error, opts ...Option) error {
	if errors.Is(err, target) {
		return nil
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected err to have target in its tree").
		Path(ops.Path).
		Want("(%T) %v", target, target).
		Have("(%T) %v", err, err)
}

// ErrorAs checks there is an error in "err" tree that matches target, and if
// one is found, sets target to that error. Returns nil if target is found,
// otherwise returns an error with a message indicating the expected and actual
// values.
func ErrorAs(err error, target any, opts ...Option) error {
	if e := Error(err); e != nil {
		return e
	}
	//goland:noinspection GoErrorsAs
	if errors.As(err, target) {
		return nil
	}
	ops := DefaultOptions().set(opts)
	return notice.New("expected err to have target in its tree").
		Path(ops.Path).
		Want("(%T) %#v", err, err).
		Have("(%T) %#v", target, target)
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
