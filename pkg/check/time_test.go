// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_FormatDates(t *testing.T) {
	t.Run("default format", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1)

		// --- Then ---
		affirm.Equal(t, "2000-01-02T03:04:05Z (2000-01-02T03:04:05Z)", have0)
		affirm.Equal(t, "2001-01-02T03:04:05Z (2001-01-02T03:04:05Z)", have1)
	})

	t.Run("custom format", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)
		opt := WithTimeFormat(time.ANSIC)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1, opt)

		// --- Then ---
		affirm.Equal(t, "Sun Jan  2 03:04:05 2000 (Sun Jan  2 03:04:05 2000)", have0)
		affirm.Equal(t, "Tue Jan  2 03:04:05 2001 (Tue Jan  2 03:04:05 2001)", have1)
	})

	t.Run("dates in different timezones", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, types.WAW)

		// --- When ---
		have0, have1 := FormatDates(tim0, tim1)

		// --- Then ---
		affirm.Equal(t, "2000-01-02T03:04:05Z      (2000-01-02T03:04:05Z)", have0)
		affirm.Equal(t, "2001-01-02T03:04:05+01:00 (2001-01-02T02:04:05Z)", have1)
	})
}
