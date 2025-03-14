// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package affirm is an internal package that provides simple affirmation
// functions designed to improve readability and minimize boilerplate code in
// test cases by offering concise, semantically meaningful functions.
package affirm

import (
	"testing"
)

// True affirms "have" is true. Returns true if it is, otherwise marks the test
// as failed, writes error message to the test log and returns false.
func True(t *testing.T, have bool) bool {
	t.Helper()
	if !have {
		const format = "expected bool to be equal:\n" +
			"\twant: %v\n" +
			"\thave: %v"
		t.Errorf(format, true, have)
		return false
	}
	return true
}

// False affirms "have" is false. Returns true if it is, otherwise marks the
// test as failed, writes error message to the test log and returns false.
func False(t *testing.T, have bool) bool {
	t.Helper()
	if have {
		const format = "expected bool to be equal:\n" +
			"\twant: %v\n" +
			"\thave: %v"
		t.Errorf(format, false, have)
		return false
	}
	return true
}

// Equal affirms two comparable types are equal. Returns true if it is,
// otherwise marks the test as failed, writes error message to the test log and
// returns false.
func Equal[T comparable](t *testing.T, want, have T) bool {
	t.Helper()
	if want != have {
		const format = "expected %T to be equal:\n" +
			"\twant: %#v\n" +
			"\thave: %#v"
		t.Errorf(format, want, want, have)
		return false
	}
	return true
}

// Nil affirms "have" is nil. Returns true if it is, otherwise marks the
// test as failed, writes error message to the test log and returns false.
func Nil(t *testing.T, have any) bool {
	t.Helper()
	if have == nil {
		return true
	}
	const format = "expected argument to be nil:\n" +
		"\twant: <nil>\n" +
		"\thave: %+v"
	t.Errorf(format, have)
	return false
}

// NotNil affirms "have" is not nil. Returns true if it is not, otherwise
// marks the test as failed, writes error message to the test log and returns
// false.
func NotNil(t *testing.T, have any) bool {
	t.Helper()
	if have != nil {
		return true
	}
	const format = "expected argument not to be nil:\n" +
		"\twant: <not-nil>\n" +
		"\thave: <nil>"
	t.Error(format)
	return false
}

// Panic affirms fn panics with a string message equal to "want". Returns true
// if fn panicked with the message, otherwise marks the test as failed, writes
// error message to the test log and returns false.
func Panic(t *testing.T, want string, fn func()) (success bool) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r != nil {
			var have string
			if have, success = r.(string); success {
				if want != have {
					format := "expected panic message:\n" +
						"\twant: %s\n" +
						"\thave: %s"
					t.Errorf(format, want, have)
					success = false
				}
				return
			}
			t.Error("expected panic() with string argument")
		}
	}()
	fn()
	t.Error("expected panic()")
	return false
}
