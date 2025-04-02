// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package affirm is an internal package that provides simple affirmation
// functions designed to improve readability and minimize boilerplate code in
// test cases by offering concise, semantically meaningful functions.
package affirm

import (
	"fmt"
	"reflect"
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

// DeepEqual affirms "want" and "have" are equal using [reflect.DeepEqual].
// Returns true if it is, otherwise marks the test as failed, writes error
// message to the test log and returns false.
func DeepEqual(t *testing.T, want, have any) bool {
	t.Helper()
	if !reflect.DeepEqual(want, have) {
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
		"\twant: nil\n" +
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
		"\thave: nil"
	t.Error(format) // TODO(rz): Consider t.Fatal.
	return false
}

// Panic affirms fn panics with a string message equal to "want". Returns true
// if fn panicked with the message, otherwise marks the test as failed, writes
// error message to the test log and returns false.
func Panic(t *testing.T, fn func()) (msg *string, success bool) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r != nil {
			success = true
			var val string
			switch v := r.(type) {
			case string:
				val = v
			case error:
				val = v.Error()
			default:
				val = fmt.Sprint(v)
			}
			msg = &val
		}
	}()

	fn()
	t.Error("expected fn to panic")
	return nil, false
}
