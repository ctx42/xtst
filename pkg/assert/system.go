// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// ExitCode asserts "err" is pointer to [exec.ExitError] with exit code equal
// to "want". Returns true if it is, otherwise marks the test as failed, writes
// error message to test log and returns false.
func ExitCode(t tester.T, want int, err error, opts ...check.Option) bool {
	t.Helper()
	if e := check.ExitCode(want, err, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}
