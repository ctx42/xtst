// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"fmt"
	"strings"
	"time"
)

// FormatDates formats two dates for comparison in an error message.
//
// Example:
//
//	2000-01-02T03:04:05Z (2000-01-02T03:04:05Z)
//	2001-01-02T03:04:05+01:00 (2001-01-02T02:04:05Z)
func FormatDates(tim0, tim1 time.Time, opts ...Option) (string, string) {
	ops := DefaultOptions().set(opts)
	tim0date := tim0.Format(ops.TimeFormat)
	tim1date := tim1.Format(ops.TimeFormat)
	tim0inUTC := tim0.In(time.UTC).Format(ops.TimeFormat)
	tim1inUTC := tim1.In(time.UTC).Format(ops.TimeFormat)
	tim0len := len(tim0date)
	tim1len := len(tim1date)
	var tim0pad, tim1pad string
	if tim0len < tim1len {
		tim0pad = strings.Repeat(" ", tim1len-tim0len)
	}
	if tim1len < tim0len {
		tim1pad = strings.Repeat(" ", tim0len-tim1len)
	}
	ret0 := fmt.Sprintf("%s %s(%s)", tim0date, tim0pad, tim0inUTC)
	ret1 := fmt.Sprintf("%s %s(%s)", tim1date, tim1pad, tim1inUTC)
	return ret0, ret1
}
