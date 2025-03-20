// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

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
