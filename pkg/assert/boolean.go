// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// True asserts "have" is true. Returns true if it's, otherwise marks the test
// as failed, writes error message to test log and returns false.
func True(t tester.T, have bool, opts ...check.Option) bool {
	t.Helper()
	if e := check.True(have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}

// False asserts "have" is false. Returns true if it's, otherwise marks the
// test as failed, writes error message to test log and returns false.
func False(t tester.T, have bool, opts ...check.Option) bool {
	t.Helper()
	if err := check.False(have, opts...); err != nil {
		t.Error(err)
		return false
	}
	return true
}
