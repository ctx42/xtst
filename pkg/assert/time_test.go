// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package assert

import (
	"testing"
	"time"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

func Test_TimeEqual(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 5, 0, types.WAW)

		// --- When ---
		got := TimeEqual(tspy, want, have)

		// --- Then ---
		affirm.True(t, got)
		affirm.True(t, want.Equal(have))
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)

		// --- When ---
		got := TimeEqual(tspy, want, have)

		// --- Then ---
		affirm.False(t, got)
		affirm.False(t, want.Equal(have))
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)
		opt := check.WithTrail("type.field")

		// --- When ---
		got := TimeEqual(tspy, want, have, opt)

		// --- Then ---
		affirm.False(t, got)
		affirm.False(t, want.Equal(have))
	})
}
