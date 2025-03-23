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

func Test_Time(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 4, 4, 5, 0, types.WAW)

		// --- When ---
		got := Time(tspy, want, have)

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
		got := Time(tspy, want, have)

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
		got := Time(tspy, want, have, opt)

		// --- Then ---
		affirm.False(t, got)
		affirm.False(t, want.Equal(have))
	})
}

func Test_TimeExact(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		got := TimeExact(tspy, want, have)

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
		have := time.Date(2000, 1, 2, 3, 4, 5, 0, types.WAW)

		// --- When ---
		got := TimeExact(tspy, want, have)

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
		have := time.Date(2000, 1, 2, 3, 4, 6, 0, time.UTC)
		opt := check.WithTrail("type.field")

		// --- When ---
		got := TimeExact(tspy, want, have, opt)

		// --- Then ---
		affirm.False(t, got)
		affirm.False(t, want.Equal(have))
	})
}

func Test_Before(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have := Before(tspy, date, mark)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("equal", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have := Before(tspy, date, mark)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		opt := check.WithTrail("type.field")

		// --- When ---
		have := Before(tspy, date, mark, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_After(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		date := time.Date(2000, 1, 2, 3, 4, 6, 0, time.UTC)
		mark := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		got := After(tspy, date, mark)

		// --- Then ---
		affirm.True(t, got)
	})

	t.Run("equal", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		got := After(tspy, date, mark)

		// --- Then ---
		affirm.False(t, got)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		opt := check.WithTrail("type.field")

		// --- When ---
		got := After(tspy, date, mark, opt)

		// --- Then ---
		affirm.False(t, got)
	})
}

func Test_EqualOrAfter(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		date := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		got := EqualOrAfter(tspy, date, mark)

		// --- Then ---
		affirm.True(t, got)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		got := EqualOrAfter(tspy, date, mark)

		// --- Then ---
		affirm.False(t, got)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		date := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		mark := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)
		opt := check.WithTrail("type.field")

		// --- When ---
		got := EqualOrAfter(tspy, date, mark, opt)

		// --- Then ---
		affirm.False(t, got)
	})
}

func Test_Within(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 3, 4, 6, 0, time.UTC)

		// --- When ---
		got := Within(tspy, want, "1s", have)

		// --- Then ---
		affirm.True(t, got)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 3, 4, 6, int(500*time.Millisecond), time.UTC)

		// --- When ---
		got := Within(tspy, want, "1s", have)

		// --- Then ---
		affirm.False(t, got)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\t    trail: type.field\n")
		tspy.Close()

		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		have := time.Date(2000, 1, 2, 3, 4, 6, int(500*time.Millisecond), time.UTC)
		opt := check.WithTrail("type.field")

		// --- When ---
		got := Within(tspy, want, "1s", have, opt)

		// --- Then ---
		affirm.False(t, got)
	})

	t.Run("want is not time.Time", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		wMsg := "[want] failed to parse time:\n\tcause: not supported time type"
		tspy.ExpectLogEqual(wMsg)
		tspy.Close()

		have := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)

		// --- When ---
		got := Within(tspy, true, "1s", have)

		// --- Then ---
		affirm.False(t, got)
	})

	t.Run("have is not time.Time", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		wMsg := "[have] failed to parse time:\n\tcause: not supported time type"
		tspy.ExpectLogEqual(wMsg)
		tspy.Close()

		want := time.Date(2000, 1, 2, 4, 4, 6, 0, types.WAW)

		// --- When ---
		got := Within(tspy, want, "1s", true)

		// --- Then ---
		affirm.False(t, got)
	})
}

func Test_Zone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := Zone(tspy, time.UTC, time.UTC)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := Zone(tspy, nil, time.UTC)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		opt := check.WithTrail("type.field")

		// --- When ---
		have := Zone(tspy, nil, time.UTC, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}

func Test_Duration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t).Close()

		// --- When ---
		have := Duration(tspy, time.Second, time.Second)

		// --- Then ---
		affirm.True(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.IgnoreLogs()
		tspy.Close()

		// --- When ---
		have := Duration(tspy, time.Second, 2*time.Second)

		// --- Then ---
		affirm.False(t, have)
	})

	t.Run("log message with trail", func(t *testing.T) {
		// --- Given ---
		tspy := tester.New(t)
		tspy.ExpectFail()
		tspy.ExpectLogContain("\ttrail: type.field\n")
		tspy.Close()

		opt := check.WithTrail("type.field")

		// --- When ---
		have := Duration(tspy, time.Second, 2*time.Second, opt)

		// --- Then ---
		affirm.False(t, have)
	})
}
