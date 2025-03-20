// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package assert provides assertion functions.
package assert

import (
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// Error asserts "err" is not nil. Returns true if it's, otherwise marks the
// test as failed, writes error message to test log and returns false.
func Error(t tester.T, err error, opts ...check.Option) bool {
	t.Helper()
	if e := check.Error(err, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// NoError asserts "err" is nil. Returns true if it is not, otherwise marks the
// test as failed, writes error message to test log and returns false.
func NoError(t tester.T, err error, opts ...check.Option) bool {
	t.Helper()
	if e := check.NoError(err, opts...); e != nil {
		t.Fatal(e)
		return false
	}
	return true
}

// ErrorIs asserts whether any error in "err" tree matches target. Returns true
// if it does, otherwise marks the test as failed, writes error message to test
// log and returns false.
func ErrorIs(t tester.T, err, target error, opts ...check.Option) bool {
	t.Helper()
	if e := check.ErrorIs(err, target, opts...); e != nil {
		t.Fatal(e)
		return false
	}
	return true
}

// ErrorAs finds the first error in "err" tree that matches target, and if one
// is found, sets target to that error. Returns true if it does, otherwise
// marks the test as failed, writes error message to test log and returns false.
func ErrorAs(t tester.T, err error, target any, opts ...check.Option) bool {
	t.Helper()
	if e := check.ErrorAs(err, target, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// ErrorEqual asserts "err" is not nil and its message equals to "want".
// Returns true if it's, otherwise marks the test as failed, writes error
// message to test log and returns false.
func ErrorEqual(t tester.T, want string, err error, opts ...check.Option) bool {
	t.Helper()
	if e := check.ErrorEqual(want, err, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// Nil asserts "have" is nil. Returns true if it is, otherwise marks the test
// as failed, writes error message to test log and returns false.
func Nil(t tester.T, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Nil(have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// NotNil asserts "have" is not nil. Returns true if it is not, otherwise marks
// the test as failed, writes error message to test log and returns false.
func NotNil(t tester.T, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.NotNil(have, opts...); e != nil {
		t.Fatal(e)
		return false
	}
	return true
}
