// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"github.com/ctx42/xtst/internal"
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// Len asserts "have" has "want" elements. Returns true if it is, otherwise it
// marks the test as failed, writes error message to test log and returns false.
func Len(t tester.T, want int, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Len(want, have, opts...); e != nil {
		cnt, _ := internal.Len(have)
		if want > cnt {
			t.Fatal(e)
		} else {
			t.Error(e)
		}
		return false
	}
	return true
}

// Has asserts slice has "want" value. Returns true if it does, otherwise marks
// the test as failed, writes error message to test log and returns false.
func Has[T comparable](t tester.T, want T, bag []T, opts ...check.Option) bool {
	t.Helper()
	if e := check.Has(want, bag, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// HasNo asserts slice does not have "want" value. Returns true if it does not,
// otherwise marks the test as failed, writes error message to test log and
// returns false.
func HasNo[T comparable](t tester.T, want T, bag []T, opts ...check.Option) bool {
	t.Helper()
	if e := check.HasNo(want, bag, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}
