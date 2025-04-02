// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package assert provides assertion functions.
package assert

import (
	"github.com/ctx42/testing/pkg/check"
	"github.com/ctx42/testing/pkg/tester"
)

// Count asserts there is "count" occurrences of "what" in "where". Returns
// true if the count matches, otherwise marks the test as failed, writes error
// message to test log and returns false.
func Count(t tester.T, count int, what, where any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Count(count, what, where, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// SameType asserts that both arguments are of the same type. Returns true if
// they are, otherwise marks the test as failed, writes error message to test
// log and returns false.
func SameType(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.SameType(want, have, opts...); e != nil {
		t.Fatal(e)
		return false
	}
	return true
}
