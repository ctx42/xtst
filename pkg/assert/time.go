// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// TimeEqual asserts both arguments are dates and are equal. The "want" and
// "have" might be date representations in form of string, int, int64 or
// [time.Time]. For string representations the [check.Options.TimeFormat] is
// used during parsing and the returned date is always in UTC. The int and
// int64 types are interpreted as Unix Timestamp and the date returned is also
// in UTC. Returns true if they are, otherwise marks the test as failed, writes
// error message to test log and returns false.
func TimeEqual(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.TimeEqual(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}
